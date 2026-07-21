package main

import (
	"fmt"
	"strings"
)

const (
	size = 4
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", size*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", size*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>", i, e.name))
	sb.WriteString("\n")

	return sb.String()
}

func (el *HtmlElement) AddChild(childName, childText string) *HtmlElement {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	el.elements = append(el.elements, e)
	return el
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	b := HtmlBuilder{
		rootName,
		HtmlElement{rootName, "", []HtmlElement{}},
	}
	return &b
}
func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func (b *HtmlBuilder) AddChildChild(childName, childText string) *HtmlBuilder {
	len := len(b.root.elements)
	if len > 0 {
		b.root.elements[len-1].AddChild(childName, childText)
	}
	return b
}

func main() {
	sb := strings.Builder{}
	sb.WriteString("<ul>")
	sb.WriteString("hello")
	sb.WriteString("</ul>")

	fmt.Println(sb.String())

	hb := NewHtmlBuilder("div")
	hb.AddChild("li", "").
		AddChildChild("span", "div").
		AddChildChild("span", "div").
		AddChildChild("span", "div").
		AddChildChild("span", "div").
		AddChild("li", "").
		AddChildChild("span", "div").
		AddChild("li", "").
		AddChildChild("span", "div").
		AddChild("li", "").
		AddChildChild("span", "div")

	fmt.Println(hb.String())
}
