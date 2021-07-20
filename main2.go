package main

import "fmt"

type Node struct {
	data  string
	left  *Node
	right *Node
}
func (root *Node) PrintPreorder() {
	fmt.Println(root.data)
	if root.left != nil {
		root.left.PrintPreorder()
	}
	if root.right != nil {
		root.right.PrintPreorder()
	}
}


func (root *Node) PrintPostorder() {
	if root.left != nil {
		root.left.PrintPostorder()
	}
	if root.right != nil {
		root.right.PrintPostorder()
	}
	fmt.Println(root.data)
}
func main() {
	nodeE := Node{data: "c", left: nil, right: nil}
	nodeD := Node{data: "b", left: nil, right: nil}
	nodeC := Node{data: "-", left: &nodeD, right: &nodeE}
	nodeB := Node{data: "a", left: nil, right: nil}
	nodeA := Node{data: "+", left: &nodeB, right: &nodeC}
	fmt.Println("Preorder")
	nodeA.PrintPreorder()
	fmt.Println("Postorder")
	nodeA.PrintPostorder()
}
