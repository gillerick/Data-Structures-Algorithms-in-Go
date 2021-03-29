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