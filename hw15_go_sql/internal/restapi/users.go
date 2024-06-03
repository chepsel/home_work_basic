package restapi

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/chepsel/home_work_basic/hw15_go_sql/internal/source"
)

func (src *Router) emptyOrErrorWrongID(queryType string, err error, i int64) (int, error) {
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError), err
	} else if i == 0 {
		return src.LogError(queryType, WrongParam, http.StatusBadRequest), WrongParam
	}
	return 0, nil
}

func (src *Router) emptyOrErrorNotFound(queryType string, err error, i int) (int, error) {
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError), err
	} else if i == 0 {
		return src.LogError(queryType, NotFound, http.StatusNotFound), WrongParam
	}
	return 0, nil
}

func (src *Router) GetUser(database *source.Database,
	resp http.ResponseWriter,
	req *http.Request,
) int {
	queryType := "GetUsers"
	params := req.URL.Query()
	id := params.Get("id")

	i, err := strconv.ParseInt(id, 10, 64)
	if code, errFunc := src.emptyOrErrorWrongID(queryType, err, i); errFunc != nil {
		return code
	}

	data, err := database.GetUser(i)
	if code, errFunc := src.emptyOrErrorNotFound(queryType, err, data.ID); errFunc != nil {
		return code
	}

	bytes, err := data.MarshalJSON()
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}
	resp.Write(bytes)
	return http.StatusOK
}

func (src *Router) GetUserStat(database *source.Database,
	resp http.ResponseWriter,
	req *http.Request,
) int {
	queryType := "GetUserStat"
	params := req.URL.Query()
	id := params.Get("id")

	i, err := strconv.ParseInt(id, 10, 64)
	if code, errFunc := src.emptyOrErrorWrongID(queryType, err, i); errFunc != nil {
		return code
	}

	data, err := database.GetUserStat(i)
	if code, errFunc := src.emptyOrErrorNotFound(queryType, err, data.ID); errFunc != nil {
		return code
	}

	bytes, err := data.MarshalJSON()
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}
	resp.Write(bytes)
	return http.StatusOK
}

func (src *Router) InsertUser(database *source.Database,
	resp http.ResponseWriter,
	req *http.Request,
) int {
	user := database.NewUser()
	queryType := "InsertUsers"
	bodyBytes, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		return src.LogError(queryType, err, http.StatusBadRequest)
	}

	err = user.UnmarshalJSON(bodyBytes)
	if err != nil {
		return src.LogError(queryType, err, http.StatusBadRequest)
	}

	err = database.InsertUser(user)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}

	bytes, err := user.MarshalJSON()
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}
	resp.Write(bytes)
	return http.StatusOK
}

func (src *Router) UpdateUser(database *source.Database,
	_ http.ResponseWriter,
	req *http.Request,
) int {
	user := database.NewUser()
	queryType := "UpdateUsers"
	bodyBytes, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		return src.LogError(queryType, err, http.StatusBadRequest)
	}

	err = user.UnmarshalJSON(bodyBytes)
	if err != nil {
		return src.LogError(queryType, err, http.StatusBadRequest)
	}

	err = database.UpdateUser(user)
	if errors.Is(err, source.NoRowsAffected) {
		return src.LogError(queryType, err, http.StatusNotFound)
	} else if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}

	return http.StatusOK
}

func (src *Router) DeleteUser(database *source.Database,
	_ http.ResponseWriter,
	req *http.Request,
) int {
	queryType := "UpdateUsers"
	params := req.URL.Query()
	id := params.Get("id")
	defer req.Body.Close()
	i, err := strconv.ParseInt(id, 10, 64)
	if code, errFunc := src.emptyOrErrorWrongID(queryType, err, i); errFunc != nil {
		return code
	}

	err = database.DeleteUser(i)
	if errors.Is(err, source.NoRowsAffected) {
		return src.LogError(queryType, err, http.StatusNotFound)
	} else if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}

	return http.StatusOK
}

// resp.WriteHeader(http.StatusInternalServerError)

func (src *Router) GetUsersList(database *source.Database,
	resp http.ResponseWriter,
	_ *http.Request,
) int {
	queryType := "GetUsersList"
	data, err := database.GetUsersList()
	if code, errFunc := src.emptyOrErrorWrongID(queryType, err, int64(len(data))); errFunc != nil {
		return code
	}

	bytes, err := database.MarshalJSONUsers(data)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}
	resp.Write(bytes)
	return http.StatusOK
}

func (src *Router) GetUsersStat(database *source.Database,
	resp http.ResponseWriter,
	_ *http.Request,
) int {
	queryType := "GetUsers"
	data, err := database.GetUsersStat()
	if code, errFunc := src.emptyOrErrorWrongID(queryType, err, int64(len(data))); errFunc != nil {
		return code
	}

	bytes, err := database.MarshalJSONUsersStat(data)
	if err != nil {
		return src.LogError(queryType, err, http.StatusInternalServerError)
	}
	resp.Write(bytes)
	return http.StatusOK
}
