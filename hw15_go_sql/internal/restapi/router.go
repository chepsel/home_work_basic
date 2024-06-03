package restapi

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/chepsel/home_work_basic/hw15_go_sql/internal/source"
)

type Router struct {
	Logger *slog.Logger
}

type SentinelError string

func (err SentinelError) Error() string {
	return string(err)
}

const (
	WrongParam     SentinelError = "wrong param"
	MissingID      SentinelError = "id is missing"
	NoRowsAffected SentinelError = "none rows affected"
	NotFound       SentinelError = "not found"
)

func NewRouter(log *slog.Logger) *Router {
	return &Router{Logger: log}
}

func UsersHandler(src *Router, source *source.Database) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var code int
		now := time.Now()
		queryType := "UsersHandler"
		switch req.Method {
		case http.MethodGet:
			code = src.GetUser(source, resp, req)
		case http.MethodPost:
			code = src.InsertUser(source, resp, req)
		case http.MethodPatch:
			code = src.UpdateUser(source, resp, req)
		case http.MethodDelete:
			code = src.DeleteUser(source, resp, req)
		default:
			code = http.StatusMethodNotAllowed
		}
		resp.WriteHeader(code)
		go src.LogResult(queryType, now, code, req)
	}
}

func UsersListHandler(src *Router, source *source.Database) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var code int
		now := time.Now()
		queryType := "UsersListHandler"
		switch req.Method {
		case http.MethodGet:
			code = src.GetUsersList(source, resp, req)
		default:
			code = http.StatusMethodNotAllowed
		}
		resp.WriteHeader(code)
		go src.LogResult(queryType, now, code, req)
	}
}

func UserStatHandler(src *Router, source *source.Database) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var code int
		now := time.Now()
		queryType := "UserStatHandler"
		switch req.Method {
		case http.MethodGet:
			code = src.GetUserStat(source, resp, req)
		default:
			code = http.StatusMethodNotAllowed
		}
		resp.WriteHeader(code)
		go src.LogResult(queryType, now, code, req)
	}
}

func UsersStatHandler(src *Router, source *source.Database) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var code int
		now := time.Now()
		queryType := "UsersStatHandler"
		switch req.Method {
		case http.MethodGet:
			code = src.GetUsersStat(source, resp, req)
		default:
			code = http.StatusMethodNotAllowed
		}
		resp.WriteHeader(code)
		go src.LogResult(queryType, now, code, req)
	}
}

func ProductsHandler(src *Router, source *source.Database) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var code int
		now := time.Now()
		queryType := "ProductsHandler"
		switch req.Method {
		case http.MethodGet:
			code = src.GetProduct(source, resp, req)
		case http.MethodPost:
			code = src.InsertProduct(source, resp, req)
		case http.MethodPatch:
			code = src.UpdateProduct(source, resp, req)
		case http.MethodDelete:
			code = src.DeleteProduct(source, resp, req)
		default:
			code = http.StatusMethodNotAllowed
		}
		resp.WriteHeader(code)
		go src.LogResult(queryType, now, code, req)
	}
}

func ProductsListHandler(src *Router, source *source.Database) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var code int
		now := time.Now()
		queryType := "ProductsListHandler"
		switch req.Method {
		case http.MethodGet:
			code = src.GetProductsList(source, resp, req)
		default:
			code = http.StatusMethodNotAllowed
		}
		resp.WriteHeader(code)
		go src.LogResult(queryType, now, code, req)
	}
}

func OrersHandler(src *Router, source *source.Database) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var code int
		now := time.Now()
		queryType := "OrersHandler"
		switch req.Method {
		case http.MethodGet:
			code = src.GetOrder(source, resp, req)
		case http.MethodPost:
			code = src.InsertOrder(source, resp, req)
		case http.MethodDelete:
			code = src.DeleteOrder(source, resp, req)
		default:
			code = http.StatusMethodNotAllowed
		}
		resp.WriteHeader(code)
		go src.LogResult(queryType, now, code, req)
	}
}

func UserOrdersHandler(src *Router, source *source.Database) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var code int
		now := time.Now()
		queryType := "UserOrdersHandler"
		switch req.Method {
		case http.MethodGet:
			code = src.GetUserOrders(source, resp, req)
		default:
			code = http.StatusMethodNotAllowed
		}
		resp.WriteHeader(code)
		go src.LogResult(queryType, now, code, req)
	}
}
