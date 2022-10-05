package health

import "github.com/labstack/echo/v4"

func Health(c echo.Context) error {
	c.Set("Content-Type", "application/json")
	return c.String(200, `{"myhealth": "100%"}`)
}
