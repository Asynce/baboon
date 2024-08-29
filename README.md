# Baboon [![GoDoc](https://godoc.org/github.com/microcosm-cc/bluemonday?status.png)](https://godoc.org/github.com/microcosm-cc/bluemonday)

---
## About
Baboon is a HTML UI component library for go webassembly or HTML tempalte projects.

## Installation
``` sh
go get github.com/Asynce/baboon
```
## Usage
If you're using a go webassembly project & need to render a beautiful div onto HTML page then you can simply utilise like this:
``` go
func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("runapp", js.FuncOf(runapp))
}

func runapp(this js.Value, i []js.Value) any {
	body := js.Global().Get("document").Call("getElementById", "body")
	body.Set("innerHTML", buildSimpleContainer())
	return nil
}

func buildSimpleContainer() string {
	return models.Container{
		Style: models.CSS{
			BackgroundColor: "#edf6f9",
			Padding:         "10px",
		},
		Children: []models.Content{
			models.Container{
				Style: models.CSS{
					BackgroundColor: "#83c5be",
					Padding:         "20px 0",
					FontSize:        "30px",
					TextAlign:       "center",
					Margin:          "10px 0",
					BorderRadius:    "12px",
				},
				Children: []models.Content{
					"Example Container",
				},
			},
		},
	}.Load()
}



```