// Package facade based on https://eax.me/golang-unit-testing/
// Что то придумывать было лень, читал про модульные тесты статью чуть выше,
// ну и взял идею за основу, по хорошему, то что ниже фасадом можно назвать
// с большой натяжкой, но абстрактное делать не хочется, а тут еще и
// синтаксис с тестами потренируюсь
package facade

import (
	"errors"

	"github.com/losev-al/GoLessons/pkg/clothes"
	"github.com/losev-al/GoLessons/pkg/person"
)

// Hero is a facade for interacting with the character in the game
type Hero struct {
	person.Person
	// Вот это место мне не понятно, надо чтобы все предметы одежды
	// поддерживали получение брони, но нельзя было надеть штаны на тело,
	// сделал, через интерфейс для брони, а дальше отдельные структуры и проверки
	// на правильный тип структуры везде где можно, недостаток, у каждой
	// структуры приходится дублировать методы, как правильно?
	bodiesClothes clothes.Clothes
	pants         clothes.Clothes
}

// IsDead return true if hero is dead
func (h *Hero) IsDead() bool {
	return h.Person.HP() == 0
}

// CanWearLegs return true if clothes can wear on Legs
func (h *Hero) CanWearLegs(c clothes.Clothes) bool {
	switch c.(type) {
	case *clothes.Pants:
		return true
	default:
		return false
	}
}

// CanWearBody return true if clothes can wear on Body
func (h *Hero) CanWearBody(c clothes.Clothes) bool {
	switch c.(type) {
	case *clothes.BodiesClothes:
		return true
	default:
		return false
	}
}

// WearLegs allows wear pants
func (h *Hero) WearLegs(c clothes.Clothes) error {
	if h.CanWearLegs(c) {
		h.pants = c
		return nil
	}
	return errors.New("Wrong type of clothes")
}

// WearBody allows wear Bodies Clothes
func (h *Hero) WearBody(c clothes.Clothes) error {
	if h.CanWearBody(c) {
		h.bodiesClothes = c
		return nil
	}
	return errors.New("Wrong type of clothes")
}

// LegsArmor return pants armor if it exist else 0
func (h *Hero) LegsArmor() int {
	if h.pants != nil {
		return h.pants.Armor()
	}
	return 0
}

// BodyArmor return bodies clothes armor if it exist else 0
func (h *Hero) BodyArmor() int {
	if h.bodiesClothes != nil {
		return h.bodiesClothes.Armor()
	}
	return 0
}

// FullArmor returns the full armor of the worn clothing if they exists else 0
func (h *Hero) FullArmor() int {
	return h.BodyArmor() + h.LegsArmor()
}

// TakeDamage calculates the damage done to the player
func (h *Hero) TakeDamage(damage int) {
	fullArmor := h.FullArmor()
	if damage > fullArmor {
		hp := h.HP()
		hp -= (damage - fullArmor)
		if hp < 0 {
			hp = 0
		}
		h.SetHP(hp)
	}
}
