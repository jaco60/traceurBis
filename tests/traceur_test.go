package traceurBis

import (
	"testing"
)

import (
	. "github.com/jaco60/traceurBis"
)

// go test -v pour voir le Printf
func TestNorth(t *testing.T) {
	Init()
	Down("black")
	Forward(10)
	Right()
	Forward(5)
	North()
	if AngleToReturnToNorth != 0 {
		t.Errorf("Forward failed - expected 0 - got x=%d", AngleToReturnToNorth)
	}
}
