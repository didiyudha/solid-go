package main

import "log"

// Color data type.
type Color int

// Size data type.
type Size int

const (
	red Color = iota
	green
	blue
)

const (
	small Size = iota
	medium
	large
)

// Product data type
type Product struct {
	name  string
	color Color
	size  Size
}

// Filter data type.
type Filter struct{}

// FilterProductByColor filter product by color.
func (f *Filter) FilterProductByColor(products []Product, color Color) []Product {

	filteredProducts := make([]Product, 0, len(products))

	for _, prod := range products {
		if prod.color == color {
			filteredProducts = append(filteredProducts, prod)
		}
	}

	return filteredProducts
}

// FilterProductBySize filter product by size.
func (f *Filter) FilterProductBySize(products []Product, size Size) []Product {
	filteredProducts := make([]Product, 0, len(products))

	if len(products) == 0 {
		return filteredProducts
	}

	for _, prod := range products {
		if prod.size == size {
			filteredProducts = append(filteredProducts, prod)
		}
	}

	return filteredProducts
}

// FilterProductBySizeAndColor - filter product by size and color.
func (f *Filter) FilterProductBySizeAndColor(products []Product, size Size, color Color) []Product {
	filteredProducts := make([]Product, 0, len(products))

	if len(products) == 0 {
		return filteredProducts
	}

	for _, prod := range products {
		if prod.size == size && prod.color == color {
			filteredProducts = append(filteredProducts, prod)
		}
	}

	return filteredProducts
}

// Better: Using OpenClosed principle.

// Specification - filter contract.
type Specification interface {
	IsSatisfied(p Product) bool
}

// ColorSpecification - color specification.
type ColorSpecification struct {
	color Color
}

// IsSatisfied - spec implementation by color specification.
func (s *ColorSpecification) IsSatisfied(p Product) bool {
	return s.color == p.color
}

// SizeSpecification - size sepecification.
type SizeSpecification struct {
	size Size
}

// IsSatisfied - spec implementation by size specification.
func (s *SizeSpecification) IsSatisfied(p Product) bool {
	return s.size == p.size
}

// AndSpecification - 'AND' specification.
type AndSpecification struct {
	specs []Specification
}

// IsSatisfied - 'AND' specification implementation.
func (s *AndSpecification) IsSatisfied(p Product) bool {

	for _, spec := range s.specs {

		if !spec.IsSatisfied(p) {
			return false
		}

	}
	return true
}

// BetterFilter - better filter data type.
type BetterFilter struct{}

// Filter - filter product implementation by better filter data type.
func (b *BetterFilter) Filter(products []Product, spec Specification) []Product {
	result := make([]Product, 0)

	for _, p := range products {
		if spec.IsSatisfied(p) {
			result = append(result, p)
		}
	}

	return result
}

func main() {

	products := []Product{
		{
			name:  "Apple",
			color: green,
			size:  small,
		},
		{
			name:  "Tree",
			color: green,
			size:  large,
		},
		{
			name:  "House",
			color: blue,
			size:  medium,
		},
	}

	filter := Filter{}
	betterFilter := BetterFilter{}

	greenProduct := filter.FilterProductByColor(products, green)
	log.Println("green product: ", greenProduct)

	greenSpec := ColorSpecification{
		color: blue,
	}
	blueProduct := betterFilter.Filter(products, &greenSpec)
	log.Println("blue product: ", blueProduct)

	largeSpec := SizeSpecification{size: large}
	largeProduct := betterFilter.Filter(products, &largeSpec)
	log.Println("large product: ", largeProduct)

	largeGreenSpec := AndSpecification{
		specs: []Specification{&greenSpec, &greenSpec},
	}
	largeGreenProduct := betterFilter.Filter(products, &largeGreenSpec)
	log.Println("large and green product: ", largeGreenProduct)

}
