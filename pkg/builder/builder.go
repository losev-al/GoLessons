package builder

import (
	"github.com/losev-al/GoLessons/pkg/facade"
	"github.com/losev-al/GoLessons/pkg/person"
)

// Builder describes interface for create character
type Builder interface {
	SetNewPersonCreateFunction(newPersonCreateFunction func() person.Person) Builder
	SetBodiesClothes() Builder
	SetPants() Builder
	GetResult() *facade.Hero
}
