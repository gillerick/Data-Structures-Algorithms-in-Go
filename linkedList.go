package main

import "fmt"

//Node class
type Node struct {
	Value int
	nextNode *Node
}

//LinkedList class
type LinkedList struct{
	headNode *Node
}


func (LinkedList *LinkedList) addNode(v int) int{
	var node = Node{}
	node.Value = v
	if node.nextNode != nil{
		node.nextNode = LinkedList.headNode
	}
	LinkedList.headNode = &node
	return v
}

func traverse(t *Node){
	if t == nil{
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil{

		fmt.Printf("%d -> ", t.Value)
		t =t.nextNode
	}
	fmt.Println()
}

var root = new(Node)

func main()  {

}
