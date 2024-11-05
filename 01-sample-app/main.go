package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		var icnt int = 1
		var txt = ""
		for i := 0; i < 100; i++ {
			txt += string(icnt) + "hogehogeほげ"
			//txt =
		}
		// return c.String(http.StatusOK, "Hello, World!")
		// return c.String(http.StatusOK, txt+"hello world")
		return c.File("index.html")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
