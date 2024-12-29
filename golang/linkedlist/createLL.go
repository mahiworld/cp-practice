package main

import "fmt"

// Node defines ll node
type Node struct {
	Val  int
	Next *Node
}

var head *Node = nil

func main() {
	n := 10

	for i := 0; i < n; i++ {
		head = createLL(head, i)
	}

	printLL(head)
}

func createLL(head *Node, ele int) *Node {
	newNode := &Node{
		Val:  ele,
		Next: nil,
	}

	if head == nil {
		head = newNode
	} else {
		temp := head
		for temp.Next != nil {
			temp = temp.Next
		}

		temp.Next = newNode
	}

	return head
}

func printLL(head *Node) {
	temp := head
	for temp != nil {
		fmt.Printf("%d -> ", temp.Val)
		temp = temp.Next
	}
	fmt.Println("nil")
}
