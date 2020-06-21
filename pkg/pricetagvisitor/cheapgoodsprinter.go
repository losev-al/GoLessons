package pricetagvisitor

import (
	"fmt"

	v1 "github.com/losev-al/GoLessons/pkg/api/v1"
)

// CheapGoodsPrinter ...
type CheapGoodsPrinter interface {
	PrintUnitWeightProduct(p *v1.UnitWeightProduct)
	PrintWeightProduct(p *v1.WeightProduct)
}

type cheapGoodsPrinter struct {}

// PrintUnitWeightProduct ...
func (e *cheapGoodsPrinter) PrintUnitWeightProduct(p *v1.UnitWeightProduct) {
	price := p.PricePerUnit / p.UnitWeight
	fmt.Printf("Товар: %v(%v кг)\nЦена: %v р/шт\nЦена за 1 кг: %v\n---\n", p.Name, p.UnitWeight, p.PricePerUnit, price)
}

// PrintWeightProduct ...
func (e *cheapGoodsPrinter) PrintWeightProduct(p *v1.WeightProduct) {
	fmt.Printf("Товар: %v\nЦена: %v р/кг\n---\n", p.Name, p.PricePerKg)
}

// NewCheapGoodsPrinter ...
func NewCheapGoodsPrinter() CheapGoodsPrinter {
	return &cheapGoodsPrinter{}
}
