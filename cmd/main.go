package main

import (
	"Avatar/pkg/clothes"
	"Avatar/pkg/facade"
	person2 "Avatar/pkg/person"
	"fmt"
)

func main() {
	person, err := person2.NewPerson("Ivan", 50)
	if err != nil {
		fmt.Println(err)
		return
	}

	jacket,err := clothes.NewClothes(clothes.BodyClothing, 10)
	if err != nil {
		fmt.Println("Ошибка создания куртки: ", err)
	}
	_, err = clothes.NewClothes(clothes.BodyClothing, -10)
	if err != nil {
		fmt.Println("Ошибка создания кольчуги: ", err)
	}

	leatherPants,err  := clothes.NewClothes(clothes.Pants, 5)
	if err != nil {
		fmt.Println("Ошибка создания кожанных штанов: ", err)
	}

	// Создаем фасад
	h := facade.NewHero(person, jacket, leatherPants)
	fmt.Printf("Создан герой по имени: %v, полная броня: %v\n", h.Name(), h.FullArmor())

	fmt.Println("Пытаемся надеть штаны на тело. Броня до: ", h.FullArmor())
	err = h.WearBody(leatherPants)
	if err != nil {
		fmt.Println("Ошибка надевания штанов: ", err)
	}
	fmt.Println("завершена попытка надеть штаны на тело. Броня после: ", h.FullArmor())

	fmt.Println("Пытаемся снять куртку. Броня до: ", h.FullArmor())
	err = h.WearBody(nil)
	if err != nil {
		fmt.Println("Ошибка снятия куртки: ", err)
	}
	fmt.Println("завершена попытка снять куртку. Броня после: ", h.FullArmor())

	fmt.Println("Пытаемся нанести урон. HP до: ", h.HP())
	h.TakeDamage(30)
	fmt.Println("завершена попытка нанести урон. HP после: ", h.HP())

	fmt.Println("Пытаемся надеть куртку. Броня до: ", h.FullArmor())
	err = h.WearBody(jacket)
	if err != nil {
		fmt.Println("Ошибка надевания куртки: ", err)
	}
	fmt.Println("завершена попытка надеть куртку. Броня после: ", h.FullArmor())

	fmt.Println("Пытаемся нанести урон. HP до: ", h.HP())
	h.TakeDamage(30)
	fmt.Println("завершена попытка нанести урон. HP после: ", h.HP())

}
