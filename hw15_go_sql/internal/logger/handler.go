package logger

import (
	"context"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

const (
	Debug = slog.LevelDebug
	Info  = slog.LevelInfo
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"
	start := []byte(string('\n'))
	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b, err := yaml.Marshal(fields)
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[2006-01-02T15:04:05.000Z]")
	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, level, msg, color.WhiteString(string(append(start, b...))))

	return nil
}

func NewPrettyHandler(out io.Writer, opts PrettyHandlerOptions) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}

func New(level slog.Level) *slog.Logger {
	handler := NewPrettyHandler(os.Stdout, PrettyHandlerOptions{SlogOpts: slog.HandlerOptions{Level: level}})
	result := slog.New(handler)
	return result
}
