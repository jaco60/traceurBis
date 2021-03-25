package traceurBis

import (
	"testing"
)

import (
	. "github.com/jaco60/traceurBis"
)

func TestNorth(t *testing.T) {
	Init()
	Right()
	Forward(5)
	North()
	if AngleToReturnToNorth != 0 {
		t.Errorf("Forward failed - expected 0 - got x=%d", AngleToReturnToNorth)
	}
	Pivote(30)
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
