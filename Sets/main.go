package Sets

type Set struct {
	integerMap map[int]bool
 }

 //Create the map of integer and bool
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}