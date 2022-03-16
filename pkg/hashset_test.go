package hashset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type coo struct {
	s string
}

func TestHashSet(t *testing.T) {

	c := NewHash[string]()
	c.Append("hello")
	assert.Equal(t, 1, c.Len())
	assert.True(t, c.Contains("hello"))
	c.RemoveAll("hello")
	assert.Equal(t, 0, c.Len())

}

func TestHashSetStruct(t *testing.T) {
	coo1 := coo{s: "1"}
	coo2 := coo{s: "2"}
	coo3 := coo{s: "2"}
	coo4 := coo{s: "4"}
	c := NewHash[coo]()
	c.Append(coo1, coo2, coo3, coo4)
	assert.Equal(t, 3, c.Len())
	assert.True(t, c.Contains(coo1))

	c.Remove(coo1)
	sl := c.AsSlice()
	assert.True(t, c.Contains(coo2))
	assert.Len(t, sl, 2)
	assert.Contains(t, sl, coo2)

	ch := c.AsChan()
	nb := 0
	for _ = range ch {
		nb++
	}
	assert.Equal(t, 2, nb)
	c.Clear()
	assert.Equal(t, 0, c.Len())
	assert.True(t, c.IsEmpty())
}
