# Templar


Templar is a simple and minimalistic template composer library designed to facilitate template creation. While it is still a work in progress, you might find it useful for streamlining your template handling. Although creating templates without a library is entirely possible, Templar aims to provide a structured approach for managing template data.

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
