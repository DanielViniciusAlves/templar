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
	Templates   []string
}

type Element struct {
	Components map[string]template.HTML
	Data       interface{}
}

func New(defaultPath string, data map[string]interface{}, order []string) *Templar {
	return &Templar{
		defaultPath: defaultPath,
		Components:  make(map[string]template.HTML),
		Data:        data,
		Templates:   order,
	}
}

func (t *Templar) ParseHTML() interface{} {
	var files []string

	temp := &Templar{
		Components: make(map[string]template.HTML),
		Data:       t.Data,
		Templates:  t.Templates,
	}

	for _, chave := range temp.Templates {
		fmt.Print(chave)
		files = append(files, t.defaultPath+chave+".html")
	}

	tmpl, _ := template.ParseFiles(files...)

	fmt.Println(files)
	for _, fileName := range files {
		name := getFileName(fileName)
		var buf bytes.Buffer

		element := &Element{
			Components: temp.Components,
			Data:       temp.Data[name],
		}

		err := tmpl.ExecuteTemplate(&buf, name, element)
		if err != nil {
			log.Fatal(err)
		}

		output := template.HTML(buf.String())
		err = temp.addComponent(name, output)
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
		}
	}

	return temp.Components
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
