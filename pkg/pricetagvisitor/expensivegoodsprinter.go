package pricetagvisitor

import (
	"fmt"

	v1 "github.com/losev-al/GoLessons/pkg/api/v1"
)

// ExpensiveGoodsPrinter ...
type ExpensiveGoodsPrinter interface {
	PrintUnitWeightProduct(p *v1.UnitWeightProduct)
	PrintWeightProduct(p *v1.WeightProduct)
}

type expensiveGoodsPrinter struct {}

// PrintUnitWeightProduct ...
func (e *expensiveGoodsPrinter) PrintUnitWeightProduct(p *v1.UnitWeightProduct) {
	price := p.PricePerUnit / p.UnitWeight / 10
	fmt.Printf("Товар: %v(%v кг)\nЦена: %v р/шт\nЦена за 100 гр:%v\n---\n", p.Name, p.UnitWeight, p.PricePerUnit, price)
}

// PrintWeightProduct ...
func (e *expensiveGoodsPrinter) PrintWeightProduct(p *v1.WeightProduct) {
	fmt.Printf("Товар: %v\nЦена: %v р за 100 гр\n---\n", p.Name, p.PricePerKg / 10)
}

// NewExpensiveGoodsPrinter ...
func NewExpensiveGoodsPrinter() ExpensiveGoodsPrinter {
	return &expensiveGoodsPrinter{}
}




