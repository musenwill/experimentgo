package binarytree

import (
	"testing"

	"github.com/emirpasic/gods/utils"
)

func TestPreorderValuesWithoutRecursion(t *testing.T) {
	tree := New(utils.IntComparator)
	cases := [][]interface{}{
		{8, "8", true},
		{4, "4", true},
		{2, "2", true},
		{1, "1", true},
		{3, "3", true},
		{6, "6", true},
		{5, "5", true},
		{7, "7", true},
		{11, "11", true},
		{10, "10", true},
		{9, "9", true},
		{12, "12", true},
		{13, "13", true},
	}
	for _, c := range cases {
		key, val, exist := c[0], c[1], c[2].(bool)
		if exist {
			tree.Put(key, val)
		}
	}

	expValues := make([]interface{}, 0, tree.Size()+1)
	tree.preOrderValues(tree.root, &expValues)
	actValues := tree.preorderValuesWithoutRecursion()
	if len(expValues) != len(actValues) {
		t.Errorf("got %v expected %v", actValues, expValues)
	}
	for i := 0; i < len(expValues); i++ {
		if exp, act := expValues[i], actValues[i]; exp != act {
			t.Errorf("got %v expected %v", actValues, expValues)
		}
	}
}

func TestInorderValuesWithoutRecursion(t *testing.T) {
	tree := New(utils.IntComparator)
	cases := [][]interface{}{
		{8, "8", true},
		{4, "4", true},
		{2, "2", true},
		{1, "1", true},
		{3, "3", true},
		{6, "6", true},
		{5, "5", true},
		{7, "7", true},
		{11, "11", true},
		{10, "10", true},
		{9, "9", true},
		{12, "12", true},
		{13, "13", true},
	}
	for _, c := range cases {
		key, val, exist := c[0], c[1], c[2].(bool)
		if exist {
			tree.Put(key, val)
		}
	}

	expValues := make([]interface{}, 0, tree.Size()+1)
	tree.inOrderValues(tree.root, &expValues)
	actValues := tree.inorderValuesWithoutRecursion()
	if len(expValues) != len(actValues) {
		t.Errorf("got %v expected %v", actValues, expValues)
	}
	for i := 0; i < len(expValues); i++ {
		if exp, act := expValues[i], actValues[i]; exp != act {
			t.Errorf("got %v expected %v", actValues, expValues)
		}
	}
}
