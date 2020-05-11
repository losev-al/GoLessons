package builder

import (
	"github.com/losev-al/GoLessons/pkg/person"
)

// humanDirector describes realization Director for human character (human have 50 hp and two closes)
type humanDirector struct {
	b Builder
}

func (d *humanDirector) Construct() {
	d.b.SetNewPersonCreateFunction(person.NewHuman)
	d.b.SetBodiesClothes()
	d.b.SetPants()
}
