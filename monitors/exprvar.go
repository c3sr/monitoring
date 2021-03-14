package monitors

// https://github.com/divan/expvarmon

import (
	"expvar"

	"github.com/labstack/echo/v4"
)

type Expvar struct{}

var (
	expvar_ Expvar
)

func (p Expvar) Routes(e *echo.Echo) error {
	if !memberQ(Config.Monitors, "expvar") {
		return nil
	}
	e.GET("/debug/vars", echo.WrapHandler(expvar.Handler()))
	return nil
}
