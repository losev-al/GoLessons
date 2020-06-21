package products

import (
	"fmt"

	v1 "github.com/losev-al/GoLessons/pkg/api/v1"
)

type weightProductVisitor interface {
	PrintWeightProduct(p *v1.WeightProduct)
}

type WeightProduct interface {
	Name() string
	PricePerKg() float32
	PrintPriceTag(visitor weightProductVisitor)
	String() string
}

type weightProduct struct {
	v1.WeightProduct
}

// Name ...
func (p *weightProduct) Name() string {
	return p.WeightProduct.Name
}

// PricePerKg ...
func (p *weightProduct) PricePerKg() float32 {
	return p.WeightProduct.PricePerKg
}

// PricePerKg ...
func (p *weightProduct) PrintPriceTag(visitor weightProductVisitor) {
	visitor.PrintWeightProduct(&p.WeightProduct)
}

// String ...
func (p *weightProduct) String() string {
	return fmt.Sprintf("Товар: %v\nцена: %v р/кг\n", p.Name(), p.PricePerKg())
}

// NewWeightProduct
func NewWeightProduct(p v1.WeightProduct) WeightProduct {
	return &weightProduct{
		WeightProduct: p,
	}
}

