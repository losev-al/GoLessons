package products

import (
	"fmt"
)

type unitWeightProductVisitor interface {
	PrintUnitWeightProduct(p UnitWeightProduct)
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
	name string
	pricePerUnit float32
	unitWeight float32
}

// Name ...
func (p *unitWeightProduct) Name() string {
	return p.name
}

// PricePerUnit ...
func (p *unitWeightProduct) PricePerUnit() float32 {
	return p.pricePerUnit
}

// UnitWeight ...
func (p *unitWeightProduct) UnitWeight() float32 {
	return p.unitWeight
}

// PrintPriceTag ...
func (p *unitWeightProduct) PrintPriceTag(visitor unitWeightProductVisitor) {
	visitor.PrintUnitWeightProduct(p)
}

// String ...
func (p *unitWeightProduct) String() string {
	return fmt.Sprintf("Товар: %v(%v кг)\nЦена: %v р/шт\n", p.Name(), p.UnitWeight(), p.PricePerUnit())
}

// NewUnitWeightProduct ...
func NewUnitWeightProduct(name string, pricePerUnit float32, unitWeight float32) UnitWeightProduct {
	return &unitWeightProduct{
		name:         name,
		pricePerUnit: pricePerUnit,
		unitWeight:   unitWeight,
	}
}