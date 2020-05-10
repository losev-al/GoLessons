package facade

import (
	"errors"
	"testing"

	"github.com/losev-al/GoLessons/pkg/clothes"
	"github.com/losev-al/GoLessons/pkg/person"
)

func TestIsDead(t *testing.T) {
	t.Parallel()
	tests := map[int]bool{0: true, 1: false}
	p := person.NewPersonmock(t)
	h := Hero{Person: p}
	for hp, want := range tests {
		h.SetHP(hp)
		got := h.IsDead()
		if got == want {
			t.Logf("IsDead() success. Want = %v, Got = %v (HP = %v)", want, got, h.HP())
		} else {
			t.Errorf("IsDead() failed. Want = %v, Got = %v (HP = %v)", want, got, h.HP())
		}
	}
}

func TestCanWearLegs(t *testing.T) {
	t.Parallel()
	tests := map[clothes.Clothes]bool{&clothes.Pants{}: true, &clothes.BodiesClothes{}: false}
	h := Hero{}
	for clothes, want := range tests {
		got := h.CanWearLegs(clothes)
		if got == want {
			t.Logf("\"CanWearLegs(%T) success. Want = %v, Got = %v", clothes, want, got)
		} else {
			t.Errorf("\"CanWearLegs(%T) failed. Want = %v, Got = %v", clothes, want, got)
		}
	}
}

func TestCanWearBody(t *testing.T) {
	t.Parallel()
	tests := map[clothes.Clothes]bool{&clothes.BodiesClothes{}: true, &clothes.Pants{}: false}
	h := Hero{}
	for clothes, want := range tests {
		got := h.CanWearBody(clothes)
		if got == want {
			t.Logf("\"CanWearBody(%T) success. Want = %v, Got = %v", clothes, want, got)
		} else {
			t.Errorf("\"CanWearBody(%T) failed. Want = %v, Got = %v", clothes, want, got)
		}
	}
}

func TestWearLegs(t *testing.T) {
	t.Parallel()
	tests := map[clothes.Clothes]error{&clothes.Pants{}: nil, &clothes.BodiesClothes{}: errors.New("Error")}
	h := Hero{}
	for clothes, want := range tests {
		got := h.WearLegs(clothes)
		// Вот здесь: https://stackoverflow.com/questions/42035104/how-to-unit-test-go-errors
		// не рекомендуют чекать текст ошибки, а чисто nil или не nil есть, какие-то внутренние
		// или внешние соглашения на эту тему?
		if got == nil && want == nil || got != nil && want != nil {
			t.Logf("\"WearLegs(%T) success. Want = %v, Got = %v", clothes, want, got)
		} else {
			t.Errorf("\"WearLegs(%T) failed. Want = %v, Got = %v", clothes, want, got)
		}
	}
}

func TestWearBody(t *testing.T) {
	t.Parallel()
	tests := map[clothes.Clothes]error{&clothes.BodiesClothes{}: nil, &clothes.Pants{}: errors.New("Error")}
	h := Hero{}
	for clothes, want := range tests {
		got := h.WearBody(clothes)
		if got == nil && want == nil || got != nil && want != nil {
			t.Logf("\"WearLegs(%T) success. Want = %v, Got = %v", clothes, want, got)
		} else {
			t.Errorf("\"WearLegs(%T) failed. Want = %v, Got = %v", clothes, want, got)
		}
	}
}

func TestLegsArmorWithoutPants(t *testing.T) {
	t.Parallel()
	h := Hero{}
	want := 0
	got := h.LegsArmor()
	if want == got {
		t.Logf("\"LegsArmor() success. Want = %v, Got = %v", want, got)
	} else {
		t.Errorf("\"LegsArmor() failed. Want = %v, Got = %v", want, got)
	}

}

func TestLegsArmorWithPants(t *testing.T) {
	t.Parallel()
	c10 := &clothes.Pants{}
	c10.SetArmor(10)
	h := Hero{}
	h.WearLegs(c10)
	want := 10
	got := h.LegsArmor()
	if want == got {
		t.Logf("\"LegsArmor() success. Want = %v, Got = %v", want, got)
	} else {
		t.Errorf("\"LegsArmor() failed. Want = %v, Got = %v", want, got)
	}
}

func TestBodyArmorWithoutBodiesClothes(t *testing.T) {
	t.Parallel()
	h := Hero{}
	want := 0
	got := h.BodyArmor()
	if want == got {
		t.Logf("\"BodyArmor() success. Want = %v, Got = %v", want, got)
	} else {
		t.Errorf("\"BodyArmor() failed. Want = %v, Got = %v", want, got)
	}

}

func TestBodyArmorWithBodiesClothes(t *testing.T) {
	t.Parallel()
	c10 := &clothes.BodiesClothes{}
	c10.SetArmor(10)
	h := Hero{}
	h.WearBody(c10)
	want := 10
	got := h.BodyArmor()
	if want == got {
		t.Logf("\"BodyArmor() success. Want = %v, Got = %v", want, got)
	} else {
		t.Errorf("\"BodyArmor() failed. Want = %v, Got = %v", want, got)
	}
}

func TestFullArmor(t *testing.T) {
	bc10 := &clothes.BodiesClothes{}
	bc10.SetArmor(10)
	p5 := &clothes.Pants{}
	p5.SetArmor(5)
	h := Hero{}
	h.WearBody(bc10)
	h.WearLegs(p5)
	want := 15
	got := h.FullArmor()
	if want == got {
		t.Logf("\"FullArmor() success. Want = %v, Got = %v", want, got)
	} else {
		t.Errorf("\"FullArmor() failed. Want = %v, Got = %v", want, got)
	}
}

func TestTakeDamage(t *testing.T) {
	tests := map[int]int{10: 100, 20: 95, 200: 0}
	bc10 := &clothes.BodiesClothes{}
	bc10.SetArmor(10)
	p5 := &clothes.Pants{}
	p5.SetArmor(5)
	h := Hero{Person: person.NewPersonmock(t)}
	h.WearBody(bc10)
	h.WearLegs(p5)
	for damage, want := range tests {
		h.SetHP(100)
		h.TakeDamage(damage)
		got := h.HP()
		if want == got {
			t.Logf("\"TakeDamage(%v) success. Want = %v, Got = %v", damage, want, got)
		} else {
			t.Errorf("\"TakeDamage(%v) failed. Want = %v, Got = %v", damage, want, got)
		}
	}
}
