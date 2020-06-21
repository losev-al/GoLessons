package pricetagvisitor

import (
	"fmt"
)

type shortUnitWeightProduct interface {
	Name() string
	PricePerUnit() float32
	UnitWeight() float32
}

type shortWeightProduct interface {
	Name() string
	PricePerKg() float32
}

// CheapGoodsPrinter ...
type CheapGoodsPrinter interface {
	PrintUnitWeightProduct(p shortUnitWeightProduct)
	PrintWeightProduct(p shortWeightProduct)
}

type cheapGoodsPrinter struct {}

// PrintUnitWeightProduct ...
func (e *cheapGoodsPrinter) PrintUnitWeightProduct(p shortUnitWeightProduct) {
	price := p.PricePerUnit() / p.UnitWeight()
	fmt.Printf("Товар: %v(%v кг)\nЦена: %v р/шт\nЦена за 1 кг: %v\n---\n", p.Name(), p.UnitWeight(), p.PricePerUnit(), price)
}

// PrintWeightProduct ...
func (e *cheapGoodsPrinter) PrintWeightProduct(p shortWeightProduct) {
	fmt.Printf("Товар: %v\nЦена: %v р/кг\n---\n", p.Name(), p.PricePerKg())
}

// NewCheapGoodsPrinter ...
func NewCheapGoodsPrinter() CheapGoodsPrinter {
	return &cheapGoodsPrinter{}
}
