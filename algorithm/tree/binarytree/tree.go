package binarytree

import (
	"github.com/emirpasic/gods/utils"
	"github.com/musenwill/experimentgo/algorithm/tree"
)

func assert() {
	var _ tree.Tree = (*Tree)(nil)
}

type Node struct {
	Key                 interface{}
	Value               interface{}
	Left, Right, Parent *Node
}

type Tree struct {
	root       *Node
	size       int
	comparator utils.Comparator
}

func New(comparator utils.Comparator) *Tree {
	return &Tree{comparator: comparator}
}

func (t *Tree) Put(key interface{}, value interface{}) {
	in := &Node{Key: key, Value: value}
	if t.root == nil {
		t.root = in
		t.size++
		return
	}

	parent := t.root
	for {
		if t.comparator(in.Key, parent.Key) < 0 {
			if parent.Left == nil {
				in.Parent = parent
				parent.Left = in
				break
			}
			parent = parent.Left
		} else {
			if parent.Right == nil {
				in.Parent = parent
				parent.Right = in
				break
			}
			parent = parent.Right
		}
	}

	t.size++
}

func (t *Tree) Get(key interface{}) (value interface{}, exist bool) {
	node := t.get(key)
	if node == nil {
		return nil, false
	}
	return node.Value, true
}

func (t *Tree) Remove(key interface{}) {
	node := t.get(key)
	if node != nil {
		t.drop(node)
		t.size--
	}
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Empty() bool {
	return t.size == 0
}

func (t *Tree) Clear() {
	t.root = nil
	t.size = 0
}

func (t *Tree) Values() []interface{} {
	values := make([]interface{}, 0, t.size+1)
	t.inOrderValues(t.root, &values)
	return values
}

func (t *Tree) inOrderValues(root *Node, values *[]interface{}) {
	if root == nil {
		return
	}

	t.inOrderValues(root.Left, values)
	*values = append(*values, root.Value)
	t.inOrderValues(root.Right, values)
}

func (t *Tree) get(key interface{}) *Node {
	node := t.root

	for node != nil {
		diff := t.comparator(key, node.Key)
		if diff < 0 {
			node = node.Left
		} else if diff > 0 {
			node = node.Right
		} else {
			return node
		}
	}

	return nil
}

func (t *Tree) drop(node *Node) {
	if node == nil {
		return
	}

	parent := node.Parent
	var replace *Node

	ln := t.leftNeighbor(node)
	rn := t.rightNeighbor(node)

	if ln == nil && rn == nil {
		replace = nil
	} else if ln != nil {
		t.drop(ln)
		replace = ln
	} else if rn != nil {
		t.drop(rn)
		replace = rn
	} else {
		panic("unreachable branch")
	}

	if replace != nil {
		replace.Parent = parent
		replace.Left = node.Left
		replace.Right = node.Right
		if node.Left != nil {
			node.Left.Parent = replace
		}
		if node.Right != nil {
			node.Right.Parent = replace
		}
	}

	if parent == nil { // root
		t.root = replace
	} else {
		if t.comparator(node.Key, parent.Key) < 0 {
			parent.Left = replace
		} else {
			parent.Right = replace
		}
	}
}

// get a smaller neighbor of node
func (t *Tree) leftNeighbor(node *Node) *Node {
	neighbor := node.Left
	for neighbor != nil && neighbor.Right != nil {
		neighbor = neighbor.Right
	}
	return neighbor
}

// get a larger neighbor of node
func (t *Tree) rightNeighbor(node *Node) *Node {
	neighbor := node.Right
	for neighbor != nil && neighbor.Left != nil {
		neighbor = neighbor.Left
	}
	return neighbor
}
