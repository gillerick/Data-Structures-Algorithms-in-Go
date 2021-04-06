package main

import "fmt"

type Set struct {
	integerMap map[int]bool
}

//Create the map of integer and bool
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

//The AddElement function adds the element to integerMap if the element is not in the set
func (set *Set) AddElement(element int)  {
	if !set.ContainsElement(element){
		set.integerMap[element] = true
	}

}

//Deletes an element from a set
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

//Checks if a set contains an element
func (set *Set) ContainsElement(element int) bool{
	var exists bool
	_, exists = set.integerMap[element]
	return exists

}

//Returns the set which intersects with another set
func (set *Set) IntersectionSet(anotherSet *Set)  *Set{
	var intersectSet = & Set{}
	intersectSet.New()
	var value int

	for value, _ = range set.integerMap{
		if anotherSet.ContainsElement(value){
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

//Returns a UnionSet that consists of a union between a set and another set
func (set *Set) Union(anotherSet *Set)  *Set{
	var unionSet = &Set{}
	unionSet.New()
	var value int
	for value, _ = range set.integerMap{
		unionSet.AddElement(value)
	}

	for value, _ = range anotherSet.integerMap{
		unionSet.AddElement(value)
	}
	return unionSet
}

func main()  {
	var set1 *Set
	var set2 *Set
	var set3 *Set

	//Set 1
	set1 = &Set{}
	set1.New()
	set1.AddElement(6)
	set1.AddElement(9)
	set1.AddElement(1)
	set1.AddElement(4)

	//Set 2
	set2 = &Set{}
	set2.New()
	set2.AddElement(3)
	set2.AddElement(4)
	set2.AddElement(2)
	set2.AddElement(7)

	//Set 3
	set3 = &Set{}
	set3.New()
	set3.AddElement(5)
	set3.AddElement(1)
	set3.AddElement(8)
	set3.AddElement(9)
	fmt.Println("Set 1: ", set1)
	fmt.Println("Set 2: ", set2)
	fmt.Println("Set 3: ", set3)
	fmt.Println("Intersect Set 1 & Set 2: ", set1.IntersectionSet(set2))
	//fmt.Println(set1.ContainsElement
	fmt.Println("Union Set 1 & Set 2: ", set1.Union(set2))

}