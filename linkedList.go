package Data_Structures___Algorithms

import "fmt"

type Node struct {
	Value int
	nextNode *Node
}


func addNode(t *Node, v int) int{
	if root == nil{
		t = &Node{v, nil}
		root = t
		return 0
	}

	if v == t.Value{
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.nextNode == nil{
		t.nextNode = &Node{v, nil}
		return -2
	}

	return addNode(t.nextNode, v)
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
