package main

import (
	"fmt"
	v1 "github.com/losev-al/GoLessons/pkg/api/v1"
	"github.com/losev-al/GoLessons/pkg/pricetagvisitor"
	"github.com/losev-al/GoLessons/pkg/products"
)

func main() {
	fmt.Println("Создаем штучно-весовой товар")
	uw := products.NewUnitWeightProduct(v1.UnitWeightProduct{
		Name:         "Пирожок с капустой",
		PricePerUnit: 25,
		UnitWeight:   0.08,
	})
	fmt.Printf("СОздан штучно-весовой товар:\n%v", uw)

	fmt.Println("Создаем весовой товар")
	w := products.NewWeightProduct(v1.WeightProduct{
		Name:       "Яблоки",
		PricePerKg: 130,
	})
	fmt.Printf("Создан весовой товар:\n%v", w)

	fmt.Println("--- Печать ценников если товары считать дорогими")
	egp := pricetagvisitor.NewExpensiveGoodsPrinter()
	uw.PrintPriceTag(egp)
	w.PrintPriceTag(egp)

	fmt.Println("--- Печать ценников если товары считать дешевыми")
	сgp := pricetagvisitor.NewCheapGoodsPrinter()
	uw.PrintPriceTag(сgp)
	w.PrintPriceTag(сgp)

}
