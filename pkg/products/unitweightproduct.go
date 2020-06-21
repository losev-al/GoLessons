package products

import (
	"fmt"

	v1 "github.com/losev-al/GoLessons/pkg/api/v1"
)

type unitWeightProductVisitor interface {
	PrintUnitWeightProduct(p *v1.UnitWeightProduct)
}

// UnitWeightProduct ...
type UnitWeightProduct interface {
	Name() string
	PricePerUnit() float32
	UnitWeight() float32
	PrintPriceTag(visitor unitWeightProductVisitor)
	String() string
}

type unitWeightProduct struct {
	v1.UnitWeightProduct
}

// Name ...
func (p *unitWeightProduct) Name() string {
	return p.UnitWeightProduct.Name
}

// PricePerUnit ...
func (p *unitWeightProduct) PricePerUnit() float32 {
	return p.UnitWeightProduct.PricePerUnit
}

// UnitWeight ...
func (p *unitWeightProduct) UnitWeight() float32 {
	return p.UnitWeightProduct.UnitWeight
}

// PrintPriceTag ...
func (p *unitWeightProduct) PrintPriceTag(visitor unitWeightProductVisitor) {
	visitor.PrintUnitWeightProduct(&p.UnitWeightProduct)
}

// String ...
func (p *unitWeightProduct) String() string {
	return fmt.Sprintf("Товар: %v(%v кг)\nЦена: %v р/шт\n", p.Name(), p.UnitWeight(), p.PricePerUnit())
}

// NewUnitWeightProduct ...
func NewUnitWeightProduct(p v1.UnitWeightProduct) UnitWeightProduct {
	return &unitWeightProduct{
		UnitWeightProduct: p,
	}
}