package tree

func (node *Node) Traverse(){
	if node == nil {
		return
	}
	if node.Left!=nil {
		node.Left.Traverse()
	}
	node.Print()
	node.Right.Traverse()
}