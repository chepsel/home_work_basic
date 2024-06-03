package config

import (
	"fmt"
	"log/slog"
	"reflect"
)

type Empty struct{}

func (src *Config) LogDBRequest(reqType string, req string, value string) {
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

func (src *Config) LogError(reqType string, err error) error {
	logger := *src.Logger

	pkgName := reflect.TypeOf(Empty{}).PkgPath()
	msg := fmt.Sprintf("%s.%s", pkgName, reqType)

	logger.Error(msg, slog.String("error", err.Error()))
	return fmt.Errorf("%s.%s: %w", msg, reqType, err)
}
