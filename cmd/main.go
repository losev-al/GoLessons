package main

import (
	"fmt"
	"os"

	"github.com/losev-al/GoLessons/pkg/facade"
	"github.com/losev-al/GoLessons/pkg/orderservice"
	"github.com/losev-al/GoLessons/pkg/warehouseservice"
)

func main() {
	orderService := orderservice.NewService()
	warehouse := warehouseservice.NewService()
	// Добавляем товары в хранилище
	warehouse.Add(1, 2)
	warehouse.Add(2, 1)
	warehouse.Add(3, 2)
	// Создаем он-лайн магазин
	market := facade.NewOnlineMarket(orderService, warehouse)
	// Создаем два заказа, на которые есть товары
	order1, err := market.CreateOrder([]int{1, 2, 3})
	if err != nil {
		fmt.Printf("Ошибка создания первого заказа, такого быть не должно: %v\n", err)
		os.Exit(1)
	}
	order2, err := market.CreateOrder([]int{1, 2, 3})
	if err != nil {
		fmt.Printf("Ошибка создания второго заказа, такого быть не должно: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Созданы два заказа")
	// Пытаемся подтвердить два заказа и выясняется, что на второй уже товаров не хватает
	err = market.ConfirmOrder(order1.Id)
	if err != nil {
		fmt.Printf("Ошибка подтверждения первого заказа: %v\n", err)
	}
	err = market.ConfirmOrder(order2.Id)
	if err != nil {
		fmt.Printf("Ошибка подтверждения второго заказа: %v\n", err)
	}
}
