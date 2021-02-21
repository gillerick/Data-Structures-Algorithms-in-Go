package Data_Structures___Algorithms

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


func () addNode(t *Node, v int) int{

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
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, 1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 10)
	addNode(root, 0)
	addNode(root, 0)
	traverse(root)
	addNode(root, 100)
	traverse(root)
}
