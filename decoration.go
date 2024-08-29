package baboon

import (
	"fmt"
	"reflect"
	"strings"
)

type Decoration struct {
	AlignSelf           string `json:"align-self"`
	All                 string `json:"align-content"`
	AspectRatio         string `json:"aspect-ratio"`
	BackdropFilter      string `json:"backdrop-filter"`
	Background          string `json:"background"`
	BackgroundColor     string `json:"background-color"`
	Border              string `json:"border"`
	BorderRadius        string `json:"border-radius"`
	BorderColor         string `json:"border-color"`
	Bottom              string `json:"bottom"`
	BoxShadow           string `json:"box-shadow"`
	Color               string `json:"color"`
	ColumnFill          string `json:"column-fill"`
	ColumnGap           string `json:"column-gap"`
	ColumnRule          string `json:"column-rule"`
	ColumnRuleColor     string `json:"column-rule-color"`
	ColumnSpan          string `json:"column-span"`
	ColumnWidth         string `json:"column-width"`
	Columns             string `json:"columns"`
	Content             string `json:"content"`
	Cursor              string `json:"cursor"`
	Display             string `json:"display"`
	Height              string `json:"height"`
	Flex                string `json:"flex"`
	Grid                string `json:"grid"`
	GridTemplateColumns string `json:"grid-template-columns"`
	Gap                 string `json:"gap"`
	GridRow             string `json:"grid-row"`
	FontSize            string `json:"font-size"`
	Padding             string `json:"padding"`
	TextAlign           string `json:"text-align"`
	Margin              string `json:"margin"`
	Width               string `json:"width"`
}

func (decoration Decoration) Style() string {
	var cssParts []string

	val := reflect.ValueOf(decoration)
	typ := reflect.TypeOf(decoration)
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		if value.Interface() != "" {
			cssParts = append(cssParts, fmt.Sprintf("%s:%s", field.Tag.Get("json"), value))
		}
	}
	return strings.Join(cssParts, "; ")
}
