package baboon

import (
	"fmt"
	"strings"
)

type Content interface{}

type Container struct {
	Style    Decoration
	Children []Content
}

func (container Container) Load() string {
	styleStr := container.Style.Style()
	var html string
	if container.Style.Style() != "" {
		html = fmt.Sprintf(`<div style="%s">`, styleStr)
	} else {
		html = "<div>"
	}

	for _, child := range container.Children {
		switch v := child.(type) {
		case string:
			html += v
		case Container:
			html += v.Load()
		case Component:
			html += v.Load()
		}
	}
	html += "</div>"
	return html
}

type Component struct {
	Containers []Container
}

func (c Component) Load() string {
	var html strings.Builder
	for _, container := range c.Containers {
		html.WriteString(container.Load())
	}
	return html.String()
}
