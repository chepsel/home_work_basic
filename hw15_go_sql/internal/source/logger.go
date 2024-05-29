package source

import (
	"fmt"
	"log/slog"
	"reflect"
)

type Empty struct{}

func (src *Database) LogDBRequest(reqType string, req string, value string) {
	logger := *src.Logger

	pkgName := reflect.TypeOf(Empty{}).PkgPath()
	msg := fmt.Sprintf("%s.%s", pkgName, reqType)

	logger.Info(msg, slog.String("request", "execute"))
	logger.Debug(
		msg,
		slog.String("query", req),
		slog.String("value", value),
	)
}

func (src *Database) LogDBResult(reqType string, res string) {
	logger := *src.Logger

	pkgName := reflect.TypeOf(Empty{}).PkgPath()
	msg := fmt.Sprintf("%s.%s", pkgName, reqType)

	logger.Info(msg, slog.String("request", "done"))
	logger.Debug(
		msg,
		slog.String("result", res),
	)
}

func (src *Database) LogError(reqType string, err error) error {
	logger := *src.Logger

	pkgName := reflect.TypeOf(Empty{}).PkgPath()
	msg := fmt.Sprintf("%s.%s", pkgName, reqType)

	logger.Error(msg, slog.String("error", err.Error()))
	return fmt.Errorf("%s: %w", msg, err)
}
