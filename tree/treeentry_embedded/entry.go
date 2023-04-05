package main

import (
	"fmt"

	"imooc.com/ccmouse/learngo/tree"
)

type myTreeMode struct {
	*tree.Node // Embedding
}

func (myNode *myTreeMode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeMode{myNode.Left}
	left.postOrder()
	right := myTreeMode{myNode.Right}
	right.postOrder()
	myNode.Print()

}

func (myNode *myTreeMode) Traverse() {
	fmt.Println("this method is shadowed.")
}

func main() {

	root := myTreeMode{&tree.Node{Value: 3}}

	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
	root.Node.Traverse()
	fmt.Println()
	root.postOrder()
}
