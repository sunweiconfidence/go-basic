package tree
import (
	"fmt"
)
type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

func (node TreeNode) Print(){
	fmt.Print(node.Value)
}

func (node *TreeNode) SetValue(value int){
	if node==nil {
		fmt.Println("set value to nil,Ignored")
		return
	}
	node.Value = value
}

func (node *TreeNode) Traverse(){
	if node == nil {
		return
	}
	if node.Left!=nil {
		node.Left.Traverse()
	}
	node.Print()
	node.Right.Traverse()
}

//工厂函数
func CreateNode(value int) *TreeNode{
	return & TreeNode{Value:value}
}


