package ioc

import (
	"log/slog"

	"github.com/universalmacro/common/auth"
	"github.com/universalmacro/common/config"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/common/ulog"
)

var GetLogHandler = singleton.EagerSingleton(func() *slog.Logger {
	return slog.New(ulog.NewHandler(0))
	// return slog.New(slog.NewJSONHandler(os.Stdout, nil))
})

var GetJwtSigner = singleton.EagerSingleton(createJwtSignerSingleton)

func createJwtSignerSingleton() *auth.JwtSigner {
	return auth.NewJwtSigner([]byte(config.GetString("jwt.secret")))
}
