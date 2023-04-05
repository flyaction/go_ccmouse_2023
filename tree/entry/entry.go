package main

import (
	"fmt"

	"imooc.com/ccmouse/learngo/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()

	nodes := []tree.Node{
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
	var root3 *tree.Node
	root3.SetValue(300)

}
