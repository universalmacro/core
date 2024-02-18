package ioc

import (
	"log/slog"

	"github.com/universalmacro/common/auth"
	"github.com/universalmacro/common/config"
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

var jwtSignerSingleton = singleton.SingletonFactory[auth.JwtSigner](createJwtSignerSingleton, singleton.Eager)

func GetJwtSigner() *auth.JwtSigner {
	return jwtSignerSingleton.Get()
}

func createJwtSignerSingleton() *auth.JwtSigner {
	return auth.NewJwtSigner([]byte(config.GetString("jwt.secret")))
}
