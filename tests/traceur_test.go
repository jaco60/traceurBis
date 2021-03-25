package traceurBis

import (
	"testing"
)

import (
	. "github.com/jaco60/traceurBis"
)

func TestNorth(t *testing.T) {
	h // https://goplay.space/#rqe8MaGVWd3
	Init()
	Down("red")
	Forward(5)
	Say("Je tourne à 3O° et j'avance de 2 pas")
	Pivote(30)
	Forward(2)
	Say("Je repars vers le nord et j'avance de 3 pas")
	North()
	if AngleToReturnToNorth != 0 {
		t.Errorf("Forward failed - expected 0 - got x=%d", AngleToReturnToNorth)
	}
}

func TestSouth(t *testing.T) {
	Init()
	Right()
	Forward(5)
	South()
	if AngleToReturnToNorth != -180 {
		t.Errorf("Forward failed - expected 0 - got x=%d", AngleToReturnToNorth)
	}
}
