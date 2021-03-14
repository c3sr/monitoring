package monitors

import "github.com/labstack/echo/v4"

func AddRoutes(e *echo.Echo) error {
	pprof_.Routes(e)
	expvar_.Routes(e)
	return nil
}
