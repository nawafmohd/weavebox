package main

import (
	"net/http"

	"github.com/weavebox/weavebox"
)

func main() {
	app := weavebox.New()

	app.Get("/hello/:name", helloHandler)

	app.Serve(3000)
}

func helloHandler(c *weavebox.Context) error {
	return c.Text(http.StatusOK, c.Param("name"))
}
