package main

import "fmt"

var m map[int]int


func main()  {

	var numbers = []int{}
	ids := []int{45, 67, 89, 54, 90, 67, 89, 12, 10}
	m = make(map[int]int)

	//Add ids to numbers list
	for _, value := range ids{
		numbers = append(numbers, value)
		m[value] = 0
	}

	fmt.Println("Array: ", numbers)
	fmt.Println("Map: ", m)


}

