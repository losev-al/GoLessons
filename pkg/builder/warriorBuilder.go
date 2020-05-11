package builder

import (
	"github.com/losev-al/GoLessons/pkg/clothes"
	"github.com/losev-al/GoLessons/pkg/facade"
	"github.com/losev-al/GoLessons/pkg/person"
)

// WarriorBuilder describes builder for create mag character
type WarriorBuilder struct {
	newPersonCreateFunction func() person.Person
	bodyClothesName         string
	bodyClothesArmor        int
	pantsName               string
	pantsArmor              int
}

// SetNewPersonCreateFunction set function for create person part of character
func (b *WarriorBuilder) SetNewPersonCreateFunction(newPersonCreateFunction func() person.Person) Builder {
	b.newPersonCreateFunction = newPersonCreateFunction
	return b
}

// SetBodiesClothes set params of bodies clothes
func (b *WarriorBuilder) SetBodiesClothes() Builder {
	b.bodyClothesName = "leather cuirass"
	b.bodyClothesArmor = 10
	return b
}

// SetPants set params of pants
func (b *WarriorBuilder) SetPants() Builder {
	b.pantsName = "leather pants"
	b.pantsArmor = 5
	return b
}

// GetResult create Hero by params
func (b *WarriorBuilder) GetResult() *facade.Hero {
	hero := &facade.Hero{}
	hero.Person = b.newPersonCreateFunction()
	if b.bodyClothesArmor != 0 && b.bodyClothesName != "" {
		bc := &clothes.BodiesClothes{Name: b.bodyClothesName}
		bc.SetArmor(b.bodyClothesArmor)
		hero.WearBody(bc)
	}
	if b.pantsArmor != 0 && b.pantsName != "" {
		p := &clothes.Pants{Name: b.pantsName}
		p.SetArmor(b.pantsArmor)
		hero.WearLegs(p)
	}
	return hero
}
