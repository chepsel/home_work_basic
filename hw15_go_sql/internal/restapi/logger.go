package restapi

import (
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type Empty struct{}

func (src *Router) LogResult(reqType string, from time.Time, code int, req *http.Request) {
	now := time.Now()
	diff := now.UnixMilli() - from.UnixMilli()
	logger := *src.Logger

	pkgName := reflect.TypeOf(Empty{}).PkgPath()
	msg := fmt.Sprintf("%s.%s", pkgName, reqType)

	logger.Info(msg, slog.String("request", req.URL.String()))
	logger.Debug(
		msg,
		slog.String("client", req.RemoteAddr),
		slog.String("url", req.URL.String()),
		slog.String("method", req.Method),
		slog.String("code", strconv.Itoa(code)),
		slog.String("processing time, ms", strconv.Itoa(int(diff))),
	)
}

func (src *Router) LogError(reqType string, err error, code int) int {
	logger := *src.Logger

	pkgName := reflect.TypeOf(Empty{}).PkgPath()
	msg := fmt.Sprintf("%s.%s", pkgName, reqType)

	logger.Error(msg, slog.String("error", err.Error()))
	return code
}
