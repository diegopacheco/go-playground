package main

import "fmt"

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func main() {
	node := Node[int]{
		Value: 1,
		Left: &Node[int]{
			Value: 2,
			Left: &Node[int]{
				Value: 4,
			},
			Right: &Node[int]{
				Value: 5,
			},
		},
		Right: &Node[int]{
			Value: 3,
		},
	}
	fmt.Printf("root node %d \n", node.Value)
	fmt.Printf("root node.left %d \n", node.Left.Value)
	fmt.Printf("root node.left.left %d \n", node.Left.Left.Value)
	fmt.Printf("root node.left.right %d \n", node.Left.Right.Value)
	fmt.Printf("root node.right %d \n", node.Right.Value)
}
