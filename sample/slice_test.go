package sample

import "testing"

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if exp, act := 3, len(s[:3]); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}
}
