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

/// lecture 2 facets

type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position         string
	AnnualIncome                  int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}
func (it *PersonBuilder) Build() *Person {
	return it.person
}

type PersonAddressBuilder struct {
	PersonBuilder
}
type PersonJobBuilder struct {
	PersonBuilder
}

func (it *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*it}
}
func (it *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*it}
}
func (pjb *PersonJobBuilder) At(
	companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(
	position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(
	annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

func (it *PersonAddressBuilder) At(
	streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(
	city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) WithPostcode(
	postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

// lecture 3 builder Parameter

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("not email, no @ in from")
	}

	eb.email.from = from
	return eb
}
func (eb *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("not email, no @ in to")
	}
	eb.email.to = to
	return eb
}
func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.email.subject = subject
	return eb
}
func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.email.body = body
	return eb
}

type build func(*EmailBuilder)

func SendEmailImplementation(email *email) {
	fmt.Println(email)
}

func SendEmailAction(action build) {
	builder := EmailBuilder{}
	action(&builder)
	SendEmailImplementation(&builder.email)
}

// lecture 4 build with actions

type PersonLc4 struct {
	firstname, lastname, address string
}

type PersonLc4Mod func(*PersonLc4)

type PersonLc4Builder struct {
	actions []PersonLc4Mod
}

func (pb *PersonLc4Builder) SetFirstName(firstname string) *PersonLc4Builder {
	pb.actions = append(pb.actions, func(pl *PersonLc4) {
		pl.firstname = firstname
	})
	return pb
}
func (pb *PersonLc4Builder) SetLastName(lastname string) *PersonLc4Builder {
	pb.actions = append(pb.actions, func(pl *PersonLc4) {
		pl.lastname = lastname
	})
	return pb
}
func (pb *PersonLc4Builder) SetAddress(address string) *PersonLc4Builder {
	pb.actions = append(pb.actions, func(pl *PersonLc4) {
		pl.address = address
	})
	return pb
}
func (pb *PersonLc4Builder) Build() *PersonLc4 {
	p := PersonLc4{}
	for _, action := range pb.actions {
		action(&p)
	}
	return &p
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

	pb := NewPersonBuilder()
	pb.
		Works().
		At("Factory").
		AsA("Developer").
		Lives().
		At("Spb").
		In("Test")
	Person := pb.Build()
	fmt.Println(Person)

	SendEmailAction(func(eb *EmailBuilder) {
		eb.From("test@test.ru").
			To("test2@test.ru").
			Subject("test").
			Body("body")
	})

	pb2 := PersonLc4Builder{}
	pb2.SetAddress("address").SetFirstName("FName").SetLastName("Lname").SetFirstName("RewriteFName")
	person := pb2.Build()
	fmt.Println(person)
}
