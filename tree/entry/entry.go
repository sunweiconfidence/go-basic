package main

import (
	"fmt"
	"github.com/sunweiconfidence/learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder(){
	if myNode ==nil || myNode.node ==nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.Node

	root = tree.Node{ Value:3 }
	root.Left = & tree.Node {}
	root.Right = & tree.Node{5,nil,nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	nodes :=[]tree.Node{
		{Value:3},
		{},
		{6,nil,&root},
	}
	fmt.Println(nodes)

	root.Print()
	fmt.Println()

	var pRoot *tree.Node
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.SetValue(300)
	fmt.Println()

	root.Traverse()
	fmt.Println()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()


}
