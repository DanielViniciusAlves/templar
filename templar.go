package templar

import (
	"bytes"
	"html/template"

	"github.com/labstack/echo/v4"
)

type compMap map[string]template.HTML

type templar struct {
	info        []TemplateInfo
	components  compMap
	defaultPath string
}

type TemplateInfo struct {
	TemplateName string
	TemplateData interface{}
	FileName     string
}

type Element struct {
	Comp compMap
	Data interface{}
}

func ParseAndRender(res int, defaultPath string, info []TemplateInfo, c echo.Context) error {
	templar := &templar{
		defaultPath: defaultPath,
		components:  make(compMap),
		info:        info,
	}

	p, err := templar.ParseHTML()
	if err != nil {
		return err
	}

	return c.HTML(res, p)
}

func (t *templar) ParseHTML() (string, error) {
	last := len(t.info) - 1

	p, err := t.parseHelper()
	if err != nil {
		return "", err
	}

	for i, info := range t.info {
		var buf bytes.Buffer

		element := &Element{
			Comp: t.components,
			Data: info.TemplateData,
		}

		err := p.ExecuteTemplate(&buf, info.TemplateName, element)
		if err != nil {
			return "", &CustomError{
				Message:      "Error executing the templates using ParseAndRender.",
				DefaultError: err,
			}
		}

		if i == last {
			return buf.String(), nil
		}

		t.components[info.TemplateName] = template.HTML(buf.String())
	}

	return "", &CustomError{
		Message: "Unexpected result from HTML parsing in Templar.",
	}
}

func (t *templar) parseHelper() (*template.Template, error) {
	var files []string

	for _, info := range t.info {
		if len(info.FileName) > 0 {
			files = append(files, t.defaultPath+info.FileName+".html")
		} else {
			files = append(files, t.defaultPath+info.TemplateName+".html")
		}
	}

	parsedTemplates, err := template.ParseFiles(files...)
	if err != nil {
		return nil, &CustomError{
			Message:      "Error parsing the template files.",
			DefaultError: err,
		}
	}

	return parsedTemplates, nil
}
