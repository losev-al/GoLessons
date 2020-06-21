package pricetagvisitor

//import (
//	"fmt"
//)
//
//type unitWeightProductVisitor interface {
//	PrintUnitWeightProduct(p unitWeightProduct)
//}
//
//type unitWeightProduct interface {
//	Name() string
//	PricePerUnit() float32
//	UnitWeight() float32
//	PrintPriceTag(visitor unitWeightProductVisitor)
//	String() string
//}
//
//type weightProductVisitor interface {
//	PrintWeightProduct(p weightProduct)
//}
//
//type weightProduct interface {
//	Name() string
//	PricePerKg() float32
//	PrintPriceTag(visitor weightProductVisitor)
//	String() string
//}
//
//// ExpensiveGoodsPrinter ...
//type ExpensiveGoodsPrinter interface {
//	PrintUnitWeightProduct(p unitWeightProduct)
//	PrintWeightProduct(p weightProduct)
//}
//
//type expensiveGoodsPrinter struct {}
//
//// PrintUnitWeightProduct ...
//func (e *expensiveGoodsPrinter) PrintUnitWeightProduct(p unitWeightProduct) {
//	price := p.PricePerUnit() / p.UnitWeight() / 10
//	fmt.Printf("Товар: %v(%v кг)\nЦена: %v р/шт\nЦена за 100 гр:%v\n---\n", p.Name, p.UnitWeight, p.PricePerUnit, price)
//}
//
//// PrintWeightProduct ...
//func (e *expensiveGoodsPrinter) PrintWeightProduct(p weightProduct) {
//	fmt.Printf("Товар: %v\nЦена: %v р за 100 гр\n---\n", p.Name(), p.PricePerKg() / 10)
//}
//
//// NewExpensiveGoodsPrinter ...
//func NewExpensiveGoodsPrinter() ExpensiveGoodsPrinter {
//	return &expensiveGoodsPrinter{}
//}




