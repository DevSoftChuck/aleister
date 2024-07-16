package main

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func always200(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hola Mundo UAA!!")
	// return c.JSON(http.StatusOK, "If you see this page, the tiny web server is successfully installed and working.")
}

func info(c echo.Context) error {
	return c.JSON(http.StatusOK, os.Environ())
}

func foo(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "Mi socio no encontre tu archivo, dejatena de joder las bolas")
}
func main() {
	e := echo.New()
	e.GET("/*", always200)
	e.GET("/info", info)
	e.GET("/aleister", foo)
	e.Logger.Fatal(e.Start(":8080"))
}
