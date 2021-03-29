package Sets

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