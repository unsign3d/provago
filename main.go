package main

import (
  "github.com/labstack/echo"
  "github.com/unsign3d/provago/middlewares"

)

func main() {
	e := echo.New()

	e.Get("/", middlewares.Welcome)
	e.Run(":1234")
}
