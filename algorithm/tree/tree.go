package tree

import (
	"github.com/emirpasic/gods/containers"
)

const (
	preOrder  = "PRE_ORDER"
	inOrder   = "IN_ORDER"
	postOrder = "POST_ORDER"
)

type Tree interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) (value interface{}, exist bool)
	Remove(key interface{})
	containers.Container
}
