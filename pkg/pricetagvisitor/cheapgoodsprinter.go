package pricetagvisitor

import (
	"fmt"
)

type unitWeightProductVisitor interface {
	PrintUnitWeightProduct(p unitWeightProduct)
}

type unitWeightProduct interface {
	Name() string
	PricePerUnit() float32
	UnitWeight() float32
	PrintPriceTag(visitor unitWeightProductVisitor)
	String() string
}

type weightProductVisitor interface {
	PrintWeightProduct(p weightProduct)
}

type weightProduct interface {
	Name() string
	PricePerKg() float32
	PrintPriceTag(visitor weightProductVisitor)
	String() string
}

// CheapGoodsPrinter ...
type CheapGoodsPrinter interface {
	PrintUnitWeightProduct(p unitWeightProduct)
	PrintWeightProduct(p weightProduct)
}

type cheapGoodsPrinter struct {}

// PrintUnitWeightProduct ...
func (e *cheapGoodsPrinter) PrintUnitWeightProduct(p unitWeightProduct) {
	price := p.PricePerUnit() / p.UnitWeight()
	fmt.Printf("Товар: %v(%v кг)\nЦена: %v р/шт\nЦена за 1 кг: %v\n---\n", p.Name(), p.UnitWeight(), p.PricePerUnit(), price)
}

// PrintWeightProduct ...
func (e *cheapGoodsPrinter) PrintWeightProduct(p weightProduct) {
	fmt.Printf("Товар: %v\nЦена: %v р/кг\n---\n", p.Name(), p.PricePerKg())
}

// NewCheapGoodsPrinter ...
func NewCheapGoodsPrinter() CheapGoodsPrinter {
	return &cheapGoodsPrinter{}
}
