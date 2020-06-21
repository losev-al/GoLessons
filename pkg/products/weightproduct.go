package products

import (
	"fmt"
	)

type weightProductVisitor interface {
	PrintWeightProduct(p shortWeightProduct)
}

type shortWeightProduct interface {
	Name() string
	PricePerKg() float32
}

type WeightProduct interface {
	Name() string
	PricePerKg() float32
	PrintPriceTag(visitor weightProductVisitor)
	String() string
}

type weightProduct struct {
	name string
	pricePerKg float32
}

// Name ...
func (p *weightProduct) Name() string {
	return p.name
}

// PricePerKg ...
func (p *weightProduct) PricePerKg() float32 {
	return p.pricePerKg
}

// PricePerKg ...
func (p *weightProduct) PrintPriceTag(visitor weightProductVisitor) {
	visitor.PrintWeightProduct(p)
}

// String ...
func (p *weightProduct) String() string {
	return fmt.Sprintf("Товар: %v\nЦена: %v р/кг\n", p.Name(), p.PricePerKg())
}

// NewWeightProduct
func NewWeightProduct(name string, pricePerKg float32) WeightProduct {
	return &weightProduct{
		name:       name,
		pricePerKg: pricePerKg,
	}
}

