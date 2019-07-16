package routinemodel

import (
	"testing"
)

func TestMonitor(t *testing.T) {
	m := newMonitor()
	m.set(9)
	if exp, act := 9, m.get(); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}
}
