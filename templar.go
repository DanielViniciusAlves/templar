package templar

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"regexp"
)

type Render interface {
	TemplarRender() string
}

type Templar struct {
	defaultPath string
	Components  map[string]template.HTML
	Data        map[string]interface{}
}

type Element struct {
	Components map[string]template.HTML
	Data       interface{}
}

func New(defaultPath string, data map[string]interface{}) *Templar {
	return &Templar{
		defaultPath: defaultPath,
		Components:  make(map[string]template.HTML),
		Data:        data,
	}
}

func (t *Templar) ParseHTML(templ ...string) interface{} {
	var files []string

	for chave := range t.Data {
		fmt.Print(chave)
		files = append(files, t.defaultPath+chave+".html")
	}

	tmpl, _ := template.ParseFiles(files...)

	for _, fileName := range files {
		name := getFileName(fileName)
		var buf bytes.Buffer

		element := &Element{
			Components: t.Components,
			Data:       t.Data[name],
		}

		err := tmpl.ExecuteTemplate(&buf, name, element)
		if err != nil {
			log.Fatal(err)
		}

		output := template.HTML(buf.String())
		err = t.addComponent(name, output)
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
		}
	}

	return t.Components
}

func getFileName(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return ""
	}

	htmlContent := string(content)

	re := regexp.MustCompile(`{{define "([^"]+)"}}`)

	matches := re.FindStringSubmatch(htmlContent)

	if len(matches) == 2 {
		valorDefinido := matches[1]
		return valorDefinido
	} else {
		fmt.Println("Padrão não encontrado no arquivo.")
		return ""
	}
}

func (t *Templar) addComponent(name string, output template.HTML) error {
	t.Components[name] = output

	return nil
}
