package tree
import (
	"fmt"
)
type Node struct {
	Value int
	Left, Right *Node
}

func (node Node) Print(){
	fmt.Print(node.Value)
}

func (node *Node) SetValue(value int){
	if node==nil {
		fmt.Println("set value to nil,Ignored")
		return
	}
	node.Value = value
}



//工厂函数
func CreateNode(value int) *Node{
	return & Node{Value:value}
}


