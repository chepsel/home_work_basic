package restapi

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/chepsel/home_work_basic/hw15_go_sql/internal/source"
)

func (src *Router) GetProduct(database *source.Database,
	resp http.ResponseWriter,
	req *http.Request,
) int {
	const queryType = "GetUsers"
	params := req.URL.Query()
	id := params.Get("id")

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	} else if i == 0 {
		return src.LogError(queryType, WrongParam, http.StatusBadRequest)
	}

	data, err := database.GetProduct(i)
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

func (src *Router) InsertProduct(database *source.Database,
	resp http.ResponseWriter,
	req *http.Request,
) int {
	product := database.NewProduct()
	queryType := "InsertProduct"
	bodyBytes, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		return src.LogError(queryType, err, http.StatusBadRequest)
	}

	err = product.UnmarshalJSON(bodyBytes)
	if err != nil {
		return src.LogError(queryType, err, http.StatusBadRequest)
	}

	err = database.InsertProduct(product)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}

	bytes, err := product.MarshalJSON()
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}
	resp.Write(bytes)
	return http.StatusOK
}

func (src *Router) UpdateProduct(database *source.Database,
	_ http.ResponseWriter,
	req *http.Request,
) int {
	product := database.NewProduct()
	queryType := "UpdateProduct"
	bodyBytes, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		return src.LogError(queryType, err, http.StatusBadRequest)
	}

	err = product.UnmarshalJSON(bodyBytes)
	if err != nil {
		return src.LogError(queryType, err, http.StatusBadRequest)
	}

	err = database.UpdateProduct(product)
	if errors.Is(err, source.NoRowsAffected) {
		return src.LogError(queryType, err, http.StatusNotFound)
	} else if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}

	return http.StatusOK
}

func (src *Router) DeleteProduct(database *source.Database,
	_ http.ResponseWriter,
	req *http.Request,
) int {
	queryType := "DeleteProduct"
	params := req.URL.Query()
	id := params.Get("id")
	defer req.Body.Close()
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	} else if i == 0 {
		return src.LogError(queryType, WrongParam, http.StatusBadRequest)
	}

	err = database.DeleteProduct(i)
	if errors.Is(err, source.NoRowsAffected) {
		return src.LogError(queryType, err, http.StatusNotFound)
	} else if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}

	return http.StatusOK
}

func (src *Router) GetProductsList(database *source.Database,
	resp http.ResponseWriter,
	_ *http.Request,
) int {
	queryType := "GetProductsList"
	data, err := database.GetProductsList()
	if err != nil {
		return src.LogError(queryType, err, http.StatusNotFound)
	} else if len(data) == 0 {
		return src.LogError(queryType, WrongParam, http.StatusNotFound)
	}

	bytes, err := database.MarshalJSONProducts(data)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}
	resp.Write(bytes)
	return http.StatusOK
}
