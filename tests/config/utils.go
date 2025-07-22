package config

import (
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func UseRunner(t *testing.T, r interface{}, dependencies ...interface{}) {
	app := fxtest.New(
		t,
		fx.Provide(dependencies...),
		fx.Invoke(r),
		fx.NopLogger,
	)

	defer app.RequireStop()
	app.RequireStart()
}
