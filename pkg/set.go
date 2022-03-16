package hashset

//Set interface of a Set
type Set[K any] interface {

	//Append a value to a set
	Append(s ...K)

	//Contains Returns true if the set contains the given value
	Contains(s K) bool

	//Len Return the current size of the set
	Len() int

	//AsSlice Return the current set as a slice
	AsSlice() []K

	//AsChan Return the current set as a chan
	AsChan() chan K

	//Remove the given value from the set, returns true if the value was present
	Remove(s K) bool

	//RemoveAll remove the given values from the set, returns true if at least one value was present
	RemoveAll(s ...K) bool

	//IsEmpty Returns true if the current set is empty
	IsEmpty() bool

	//Clear removes all values from the set
	Clear()
}
