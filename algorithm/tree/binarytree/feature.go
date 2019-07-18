package binarytree

import (
	"errors"

	"github.com/emirpasic/gods/stacks/arraystack"
)

/*
非递归前序打印二叉树，本质与递归一样，需要用到栈。
需要注意的是，为了保证左子树先处理，右子树后处理，所以右子树节点先入栈，左子树后入栈，后入先出。
*/
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

/*
非递归中序打印二叉树
*/
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

/*
给定前序与中序遍历结果，输出后序遍历结果
*/
func calcPostOrder(preOrder, inOrder []interface{}) ([]interface{}, error) {
	preLen := len(preOrder)
	inLen := len(inOrder)
	if preLen != inLen {
		return nil, errors.New("invalid input, length of pre order and in order not match")
	}
	if preLen < 1 {
		return nil, nil
	}

	root, err := constructTree(preOrder, inOrder)
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, 0, preLen)
	postOrder(root, &values)

	return values, nil
}

func postOrder(root *Node, values *[]interface{}) {
	if root == nil {
		return
	}
	postOrder(root.Left, values)
	postOrder(root.Right, values)
	*values = append(*values, root.Key)
}

func constructTree(preOrder, inOrder []interface{}) (*Node, error) {
	root := &Node{Key: preOrder[0]}

	if len(preOrder) == 1 {
		return root, nil
	}

	rootIndex := 0
	for ; rootIndex < len(inOrder) && inOrder[rootIndex] != preOrder[0]; rootIndex++ {
	}
	if rootIndex >= len(inOrder) {
		return nil, errors.New("invalid input")
	}

	inOrderLeft := inOrder[0:rootIndex]
	inOrderRight := inOrder[rootIndex+1:]

	preOrderLeft := preOrder[1 : rootIndex+1]
	preOrderRight := preOrder[rootIndex+1:]

	var leftRoot, rightRoot *Node
	var err error
	if len(preOrderLeft) > 0 {
		leftRoot, err = constructTree(preOrderLeft, inOrderLeft)
		if err != nil {
			return nil, err
		}
	}
	if len(preOrderRight) > 0 {
		rightRoot, err = constructTree(preOrderRight, inOrderRight)
		if err != nil {
			return nil, err
		}
	}

	if leftRoot != nil {
		leftRoot.Parent = root
		root.Left = leftRoot
	}
	if rightRoot != nil {
		rightRoot.Parent = root
		root.Right = rightRoot
	}

	return root, nil
}
