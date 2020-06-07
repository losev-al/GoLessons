package clothes

import "fmt"

// ElementOfClothes describes the behavior of protective clothing
type ElementOfClothes interface {
	Armor() int
	Type() Type
}

type elementOfClothes struct {
	armor int
	clothesType Type
}

// Armor get armor of clothes
func (c *elementOfClothes) Armor() int {
	return c.armor
}

// Type get type of clothes
func (c *elementOfClothes) Type() Type {
	return c.clothesType
}

// NewClothes create clothing of the specified type
func NewClothes(t Type, armor int) (ElementOfClothes, error) {
	if armor < 0 {
		return nil, fmt.Errorf("armor cannot be less than 0. The passed value %v", armor)
	}
	return &elementOfClothes{armor: armor,clothesType: t}, nil
}
