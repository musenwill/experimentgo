package binarytree

import (
	"github.com/emirpasic/gods/stacks/arraystack"
)

func (t *Tree) preorderValuesWithoutRecursion() []interface{} {
	values := make([]interface{}, 0, t.size+1)
	stack := arraystack.New()

	if t.root == nil {
		return values
	}
	stack.Push(t.root)

	for !stack.Empty() {
		element, _ := stack.Pop()
		node := element.(*Node)
		values = append(values, node.Value)

		if node.Right != nil {
			stack.Push(node.Right)
		}

		if node.Left != nil {
			stack.Push(node.Left)
		}
	}

	return values
}

func (t *Tree) inorderValuesWithoutRecursion() []interface{} {
	values := make([]interface{}, 0, t.size+1)
	stack := arraystack.New()

	if t.root == nil {
		return values
	}
	stack.Push(t.root)

	push := true
	for !stack.Empty() {
		element, _ := stack.Peek()
		node := element.(*Node)

		if node.Left != nil && push {
			stack.Push(node.Left)
			push = true
			continue
		}

		element, _ = stack.Pop()
		node = element.(*Node)
		values = append(values, node.Value)
		push = false

		if node.Right != nil {
			stack.Push(node.Right)
			push = true
			continue
		}
	}

	return values
}
