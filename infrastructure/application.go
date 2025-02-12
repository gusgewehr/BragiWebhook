package infrastructure

import (
	"go.uber.org/zap"
)

type Application struct {
	Env    *Env
	Amqp   *Amqp
	Logger *zap.Logger
}

func App(envPath string) *Application {

	app := &Application{}

	app.Logger = NewLogger()

	app.Env = NewEnv(envPath, app.Logger)

	app.Amqp = NewAmqp(*app.Env)

	return app
}
