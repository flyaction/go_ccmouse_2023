package main

import (
	"fmt"

	"imooc.com/ccmouse/learngo/tree"
)

func main() {
	var root tree.TreeNode
	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()

	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	root.SetValue(10)
	root.Print()
	root2 := &root
	root2.Print()
	root2.SetValue(100)
	root.SetValue(200)
	root.Print()
	root2.Print()

	fmt.Println()
	var root3 *tree.TreeNode
	root3.SetValue(300)

}
