package main

import "fmt"

type Product struct {
	Name  string
	Color Color
	Size  Size
}

type Color int
type Size int

const (
	Red Color = iota
	Blue
	Green
)

const (
	Small Size = iota
	Medium
	Large
)

type Filter struct{}

func (f *Filter) filterProductsByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.Color == color {
			result = append(result, &products[i])
		}
	}
	return result
}
func (f *Filter) filterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.Size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySizeAndColor(
	products []Product, size Size,
	color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.Size == size && v.Color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecifecation struct {
	Color Color
}

func (spec ColorSpecifecation) IsSatisfied(p *Product) bool {
	return p.Color == spec.Color
}

type SizeSpecification struct {
	Size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.Size == spec.Size
}

type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

type OrSpecification struct {
	first, second Specification
}

func (spec OrSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) || spec.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (bf *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"apple", Green, Small}
	tree := Product{"tree", Green, Large}
	house := Product{"house", Blue, Large}

	products := []Product{apple, tree, house}
	filter := Filter{}
	filtered := filter.filterProductsByColor(products, Green)

	for fProduct := range filtered {
		fmt.Println("old", filtered[fProduct])
	}

	Specification := ColorSpecifecation{Green}
	BetterFilter := BetterFilter{}
	bfiltered := BetterFilter.Filter(products, Specification)
	for fProduct := range bfiltered {
		fmt.Println("new green", bfiltered[fProduct])
	}

	LargeSpec := SizeSpecification{Large}
	SpecificationAnd := AndSpecification{Specification, LargeSpec}
	bfiltered2 := BetterFilter.Filter(products, SpecificationAnd)
	for fProduct := range bfiltered2 {
		fmt.Println("new green and large", bfiltered2[fProduct])
	}
}
