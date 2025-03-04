package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	tmpl, err := template.ParseGlob("views/*.html")
	if err != nil {
		log.Fatalf("Unable to parse glob %e\n", err)
	}

	e := echo.New()
	e.Renderer = &TemplateRenderer {
		templates: tmpl,
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.GET("/navbar/:page", func(c echo.Context) error {
		name := c.Param("page")
		return c.Render(http.StatusOK, "navbar.html", name)
	})

	e.GET("/books", func(c echo.Context) error {
		return c.Render(http.StatusOK, "books.html", nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
