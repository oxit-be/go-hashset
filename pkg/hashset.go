package hashset

// HashSet implementation of a set
type HashSet[K comparable] struct {
	set map[K]bool
}

func (h *HashSet[K]) Clear() {
	h.set = make(map[K]bool, 10)
}

//NewHash Creates a new Set
func NewHash[K comparable]() Set[K] {
	m := make(map[K]bool, 10)
	return &HashSet[K]{set: m}
}

func (h *HashSet[K]) Append(s ...K) {
	for i := range s {
		h.set[s[i]] = true
	}
}

// Contains Return true if the given value is present
func (h *HashSet[K]) Contains(value K) bool {
	return h.set[value]
}

//IsEmpty Returns true if the set empty
func (h *HashSet[K]) IsEmpty() bool {
	return len(h.set) == 0
}

// Len Returns the size of the current set
func (h *HashSet[K]) Len() int {
	return len(h.set)
}

// AsSlice Returns the current set in a slice
func (h *HashSet[K]) AsSlice() []K {
	slice := make([]K, 0, h.Len())
	for k := range h.set {
		slice = append(slice, k)
	}
	return slice
}

func (h *HashSet[K]) AsChan() chan K {
	ch := make(chan K)
	go func() {
		defer close(ch)
		for k := range h.set {
			ch <- k
		}
	}()
	return ch
}

//Remove the given value from the set
func (h *HashSet[K]) Remove(val K) bool {
	defer delete(h.set, val)
	return h.set[val]
}

func (h *HashSet[K]) RemoveAll(s ...K) bool {
	ok := false
	for i := range s {
		ok = ok || h.Remove(s[i])
	}
	return ok
}
