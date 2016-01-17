package middlewares

import(
  "net/http"
  "github.com/labstack/echo"
  "github.com/unsign3d/provago/models"
)


func Welcome(c *echo.Context) error {

  wm :=  models.Welcome{
    Id: "123456",
    Message: "Hello world",
  }

	return c.String(http.StatusOK, wm.Message)
}
