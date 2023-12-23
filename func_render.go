package templar

import (
	"bytes"
	"html/template"

	"github.com/labstack/echo/v4"
)

type render func() *FuncRenderStruct

type FuncRenderStruct struct {
	Data     interface{}
	Template string
}

func FuncRender(res int, render render, c echo.Context) error {

	funcRender := render()
	tmpl, err := template.New("").Parse(funcRender.Template)
	if err != nil {
		return &CustomError{
			Message:      "Error parsing the template using FuncRender.",
			DefaultError: err,
		}
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, funcRender.Data)
	if err != nil {
		return &CustomError{
			Message:      "Error executing the template using FuncRender.",
			DefaultError: err,
		}
	}

	return c.HTML(res, buf.String())
}
