package ioc

import (
	"log/slog"

	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/common/ulog"
)

var loggerSingleton = singleton.SingletonFactory(func() *slog.Logger {
	return slog.New(ulog.NewHandler(0))
	// return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}, singleton.Eager)

func GetLogHandler() *slog.Logger {
	return loggerSingleton.Get()
}
