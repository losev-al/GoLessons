package facade

import (
	"Avatar/pkg/clothes"
	"fmt"
)

type person interface {
	Name() string
	HP() int
	SetHP(hp int) error
}

type elementOfClothes interface {
	Armor() int
	// Во это место мне не понятно, я интервейс приватный объявляю,
	// но возврщаться должен тип из пакета, все равно возникает зависимость,
	// как в этом случае правильно посутпать?
	Type() clothes.Type
}

// Hero is a facade for interacting with the character in the game
type Hero interface {
	HP() int
	Name() string
	IsDead() bool
	CanWearLegs(c elementOfClothes) bool
	CanWearBody(c elementOfClothes) bool
	WearLegs(c elementOfClothes) error
	WearBody(c elementOfClothes) error
	LegsArmor() int
	BodyArmor() int
	FullArmor() int
	TakeDamage(damage int)
}


type hero struct {
	person
	bodiesClothes elementOfClothes
	pants         elementOfClothes
}

// HP get HP of hero
func (h *hero) HP() int {
	return h.person.HP()
}

// Name get name of hero
func (h *hero) Name() string {
	return h.person.Name()
}

// IsDead return true if hero is dead
func (h *hero) IsDead() bool {
	return h.HP() == 0
}

// CanWearLegs return true if clothes can wear on Legs
func (h *hero) CanWearLegs(c elementOfClothes) bool {
	// Аналогично, мне нужна константа из зависимого пакета. Как правильно поступать?
	return c.Type() == clothes.Pants
}

// CanWearBody return true if clothes can wear on Body
func (h *hero) CanWearBody(c elementOfClothes) bool {
	return c.Type() == clothes.BodyClothing
}

// WearLegs allows wear pants
func (h *hero) WearLegs(c elementOfClothes) error {
	if c == nil || h.CanWearLegs(c) {
		h.pants = c
		return nil
	}
	return fmt.Errorf("wrong type of clothes. Need: %v, passed: %v", clothes.Pants, c.Type())
}

// WearBody allows wear Bodies Clothes
func (h *hero) WearBody(c elementOfClothes) error {
	if c == nil || h.CanWearBody(c) {
		h.bodiesClothes = c
		return nil
	}
	return fmt.Errorf("wrong type of clothes. Need: %v, passed: %v", clothes.BodyClothing, c.Type())
}

// LegsArmor return pants armor if it exist else 0
func (h *hero) LegsArmor() int {
	if h.pants != nil {
		return h.pants.Armor()
	}
	return 0
}

// BodyArmor return bodies clothes armor if it exist else 0
func (h *hero) BodyArmor() int {
	if h.bodiesClothes != nil {
		return h.bodiesClothes.Armor()
	}
	return 0
}

// FullArmor returns the full armor of the worn clothing if they exists else 0
func (h *hero) FullArmor() int {
	return h.BodyArmor() + h.LegsArmor()
}

// TakeDamage calculates the damage done to the player
func (h *hero) TakeDamage(damage int) {
	fullArmor := h.FullArmor()
	if damage > fullArmor {
		hp := h.HP()
		hp -= damage - fullArmor
		if hp < 0 {
			hp = 0
		}
		h.SetHP(hp)
	}
}

// NewHero create hero
func NewHero(p person, bodiesClothes elementOfClothes, pants elementOfClothes) Hero {
	return &hero{person:p,bodiesClothes: bodiesClothes,pants: pants}
}
