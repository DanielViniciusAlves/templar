# Templar

Templar is a simple and minimalist template composer library designed to facilitate template creation. You probably don't need this; I just created it to simplify my template creation process. So, when I want to undertake a server-side rendering project, I don't need to worry about how to access data in the template and how to render each template.

## Usage Example

To use Templar, create a map where the keys represent file names and the values contain the data that the corresponding template requires. Ensure that the order of keys in the map reflects the significance of the templates, from least to most significant.

**The order passed must be from the least significant to the most significant**

Here's an example implementation using Echo:

```
type PageData struct {
    Title string
}

func home(c echo.Context) error {
    // Create a map with template data
    homeMap := make(map[string]interface{})
    homeMap["ComponentFileName"] = PageData{Title: "Component"}
    homeMap["Home"] = PageData{Title: "Radar"}
    
    // Define the order of template rendering
    order := []string{"Component", "Home"}

    // Create a Templar instance
    templar := templar.New("View/templates/", homeMap, order)

    // Generate HTML output
    output := templar.ParseHTML()

    // Send output to the client using the preferred library
    return c.Render(res, "base", output)
}
```

In the HTML files, access template data with .Data.(data name) and the corresponding component with .Components.(component name):


```
{{define "Home"}}
  <h1>{{.Data.Title}}</h1>
{{.Components.Component}}
{{end}}

```

```
{{define "Component"}}
  <h1>{{.Data.Title}}</h1>
{{end}}
```
