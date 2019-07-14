package binarytree

import (
	"testing"

	"github.com/emirpasic/gods/utils"
)

func TestTreeEmpty(t *testing.T) {
	tree := New(utils.IntComparator)

	if exp, act := 0, tree.Size(); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}

	if exp, act := true, tree.Empty(); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}

	if exp, act := 0, len(tree.Values()); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}

	_, exist := tree.Get(0)
	if exp, act := false, exist; exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}

	tree.Remove(0)
}

func TestTreeSpecifiedInt(t *testing.T) {
	tree := New(utils.IntComparator)

	// 10 elements
	cases := [][]interface{}{
		{6, "a", true},
		{5, "b", true},
		{4, "c", true},
		{3, "d", true},
		{2, "e", true},
		{1, "f", true},
		{7, "g", true},
		{8, "h", true},
		{9, "i", true},
		{10, "j", true},
		{20, "k", false},
		{21, "l", false},
		{22, "m", false},
	}

	testCases(t, tree, cases)
}

func TestTreeSpecifiedString(t *testing.T) {
	tree := New(utils.StringComparator)

	// 10 elements
	cases := [][]interface{}{
		{"a", 1, true},
		{"b", 2, true},
		{"c", 3, true},
		{"d", 4, true},
		{"e", 5, true},
		{"f", 6, true},
		{"g", 7, true},
		{"h", 8, true},
		{"i", 9, true},
		{"j", 10, true},
		{"k", 11, false},
		{"l", 12, false},
		{"m", 13, false},
	}

	testCases(t, tree, cases)
}

func testCases(t *testing.T, tree *Tree, cases [][]interface{}) {
	for _, c := range cases {
		key, val, exist := c[0], c[1], c[2].(bool)
		if exist {
			tree.Put(key, val)
		}
	}

	if exp, act := 10, tree.Size(); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}

	if exp, act := false, tree.Empty(); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}

	if exp, act := 10, len(tree.Values()); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}

	for _, c := range cases {
		key, expVal, expExist := c[0], c[1], c[2].(bool)
		actVal, actExist := tree.Get(key)
		if expExist != actExist || expExist == true && expVal != actVal {
			t.Errorf("exist got %v expected %v, value got %v expected %v", actExist, expExist, actVal, expVal)
		}
	}

	size := tree.Size()
	for _, c := range cases {
		key, exist := c[0], c[2].(bool)
		tree.Remove(key)
		size--
		_, actExist := tree.Get(key)
		if exp, act := false, actExist; exp != act {
			t.Errorf("got %v expected %v", act, exp)
		}
		if exist {
			if exp, act := size, tree.Size(); exp != act {
				t.Errorf("got %v expected %v", act, exp)
			}
		}
	}

	if exp, act := 0, tree.Size(); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}

	if exp, act := true, tree.Empty(); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}

	if exp, act := 0, len(tree.Values()); exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}
}

func BenchmarkTreeRandomInt(b *testing.B) {
	tree := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	for i := 0; i < b.N; i++ {
		key := i * 3662551 % 8347337
		tree.Put(key, i)
	}
	for i := 0; i < b.N; i++ {
		key := i * 3662551 % 8347337
		tree.Get(key)
	}
}
