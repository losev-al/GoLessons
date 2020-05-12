package builder

import (
	"reflect"
	"testing"

	"github.com/losev-al/GoLessons/pkg/person"
)

func TestCreateMage(t *testing.T) {
	t.Parallel()
	b := MageBuilder{}
	character := b.SetNewPersonCreateFunction(person.NewHuman).SetBodiesClothes().SetPants().GetResult()
	wantRace := reflect.TypeOf((*person.Human)(nil)).Elem()
	gotRace := reflect.TypeOf(character.Person).Elem()
	if gotRace == wantRace {
		t.Logf("Set Race success. Want = %v, Got = %v", wantRace, gotRace)
	} else {
		t.Errorf("Set Race failed. Want = %v, Got = %v", wantRace, gotRace)
	}
	wantArmor := 7
	gotArmor := character.FullArmor()
	if gotArmor == wantArmor {
		t.Logf("Set clothes success. Want = %v, Got = %v", wantArmor, gotArmor)
	} else {
		t.Errorf("Set clothes failed. Want = %v, Got = %v", wantArmor, gotArmor)
	}
}