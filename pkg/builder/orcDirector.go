package builder

import (
	"github.com/losev-al/GoLessons/pkg/person"
)

// orcDirector describes realization Director for orc character (orc have 70 hp and one pants)
type orcDirector struct {
	b Builder
}

func (d *orcDirector) Construct() {
	d.b.SetNewPersonCreateFunction(person.NewOrc)
	d.b.SetPants()
}
