# 二叉查找树

## Remove

1. 假设待删除节点为 A，找到 A 节点的左邻居或者右邻居，左邻居或右邻居可能为空

把二叉查找树按照中序遍历打印出来，将得到一个排序数组，所谓左邻居，是指该数组中左边紧邻要删除节点
的元素，而右邻居是指数组中右边紧邻要删除节点的元素。如下面一棵树中:

```
        ┌── 13
    ┌── 12
┌── 11
|   └── 10
|       └── 9
8       ┌── 7
│   ┌── 6
|   |   └── 5
└── 4
    │   ┌── 3
    └── 2
        └── 1
```

中序遍历结果: 1 2 3 4 5 6 7 8 9 10 11 12 13
那么节点 8 的左邻居是 7，右邻居是 9；
节点 11 的左邻居是 10，右邻居是 12。
可见一个节点的左邻居或是有邻居一定是位于叶子节点上的，而且满足:
左邻居在左子树的最右侧叶子节点上，右邻居在右子树的最左侧叶子节点上。

故左邻居或是右邻居的查找方法特别简单:

```
func (t *Tree) leftNeighbor(node *Node) *Node {
	neighbor := node.Left
	for neighbor != nil && neighbor.Right != nil {
		neighbor = neighbor.Right
	}
	return neighbor
}
```

```
func (t *Tree) rightNeighbor(node *Node) *Node {
	neighbor := node.Right
	for neighbor != nil && neighbor.Left != nil {
		neighbor = neighbor.Left
	}
	return neighbor
}
```

2. 找到邻居即是找到了候选者，称为 C，用于填补删除 A 节点留下来的空缺
3. 摘除邻居节点
4. 删除 A 节点
5. 将邻居节点 C 放置在 A 节点留下的空位上，设置好前后关联的指针

该方法对任意 A 节点都有效，包括叶子节点和根节点。


## 边界条件

1. 根节点为空
2. 左右子树为空
