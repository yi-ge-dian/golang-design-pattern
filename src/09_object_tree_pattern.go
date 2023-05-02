package main

import "fmt"

// Node 定义树结构中的一个节点
type Node struct {
	Value    string
	Children []*Node
}

// AddChild 添加一个子节点
func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

// Print 打印树结构
func (n *Node) Print(indent string) {
	fmt.Println(indent + n.Value)
	for _, child := range n.Children {
		child.Print(indent + " ")
	}
}

func CallObjectTreePattern() {
	root := &Node{Value: "Root"}
	//添加一级子节点
	child1 := &Node{Value: "Child 1"}
	child2 := &Node{Value: "Child 2"}
	root.AddChild(child1)
	root.AddChild(child2)
	//添加二级子节点
	grandchild1 := &Node{Value: "Grandchild 1"}
	grandchild2 := &Node{Value: "Grandchild 2"}
	child1.AddChild(grandchild1)
	child2.AddChild(grandchild2)
	//打印树结构
	root.Print("")
}
