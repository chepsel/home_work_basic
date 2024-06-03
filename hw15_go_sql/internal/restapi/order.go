package restapi

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/chepsel/home_work_basic/hw15_go_sql/internal/source"
)

func (src *Router) GetOrder(database *source.Database,
	resp http.ResponseWriter,
	req *http.Request,
) int {
	queryType := "GetOrder"
	params := req.URL.Query()
	id := params.Get("id")

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	} else if i == 0 {
		return src.LogError(queryType, WrongParam, http.StatusBadRequest)
	}

	data, err := database.GetOrder(i)
	if err != nil {
		return src.LogError(queryType, err, http.StatusNotFound)
	} else if data.ID == 0 {
		return src.LogError(queryType, WrongParam, http.StatusNotFound)
	}

	bytes, err := data.MarshalJSON()
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}
	resp.Write(bytes)
	return http.StatusOK
}

func (src *Router) InsertOrder(database *source.Database,
	resp http.ResponseWriter,
	req *http.Request,
) int {
	order := database.NewOrder()
	queryType := "InsertOrder"
	bodyBytes, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		return src.LogError(queryType, err, http.StatusBadRequest)
	}

	err = order.UnmarshalJSON(bodyBytes)
	if err != nil {
		return src.LogError(queryType, err, http.StatusBadRequest)
	}

	err = database.InsertOrder(order)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}

	bytes, err := order.MarshalJSON()
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}
	resp.Write(bytes)
	return http.StatusOK
}

func (src *Router) DeleteOrder(database *source.Database,
	_ http.ResponseWriter,
	req *http.Request,
) int {
	queryType := "DeleteOrder"
	params := req.URL.Query()
	id := params.Get("id")
	defer req.Body.Close()
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	} else if i == 0 {
		return src.LogError(queryType, WrongParam, http.StatusBadRequest)
	}

	err = database.DeleteOrder(i)
	if errors.Is(err, source.NoRowsAffected) {
		return src.LogError(queryType, err, http.StatusNotFound)
	} else if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}
	return http.StatusOK
}

func (src *Router) GetUserOrders(database *source.Database,
	resp http.ResponseWriter,
	req *http.Request,
) int {
	queryType := "GetUserOrders"
	params := req.URL.Query()
	id := params.Get("id")

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	} else if i == 0 {
		return src.LogError(queryType, WrongParam, http.StatusBadRequest)
	}

	data, err := database.GetUserOrders(i)
	if err != nil {
		return src.LogError(queryType, err, http.StatusNotFound)
	} else if len(data) == 0 {
		return src.LogError(queryType, WrongParam, http.StatusNotFound)
	}

	bytes, err := database.MarshalJSONOrders(data)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}
	resp.Write(bytes)
	return http.StatusOK
}
