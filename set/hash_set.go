package set

import (
	"bytes"
	"fmt"
)

// Define a HashSet, and depends on the map implementation.
type HashSet struct {
	m map[interface{}]bool
}

// Define a new Hash Set.
func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

// Add element e to s and return bool result.
func (s *HashSet) Add(e interface{}) bool {
	if !s.m[e] {
		s.m[e] = true
		return true
	}
	return false
}

// Clear Set s.
func (s *HashSet) Clear() {
	s.m = make(map[interface{}]bool)
}

// Determine whether s contains element e and return bool result.
func (s *HashSet) Has(e interface{}) bool {
	return s.m[e]
}

// Determine if a is equal to s, return bool result.
func (s *HashSet) Equals(a Set) bool {
	if nil == a {
		return false
	}
	if s.Len() != a.Len() {
		return false
	}
	for key := range s.m {
		if !a.Has(key) {
			return false
		}
	}
	return true
}

// Determine if the set s is empty, return bool result.
func (s *HashSet) IsEmpty() bool {
	if 0 == s.Len() {
		return true
	}
	return false
}

// Remove element e from collection Set s.
func (s *HashSet) Remove(e interface{}) {
	delete(s.m, e)
}

// Get the length of the Set s.
func (s *HashSet) Len() int {
	return len(s.m)
}

// Iterating over the elements in the output Set s.
func (s *HashSet) Elements() []interface{} {
	initialLen := len(s.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range s.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

// Convert Set s to string output.
func (s *HashSet) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range s.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

// Convert the Set s into a slice.
func (s *HashSet) ToArray() []interface{} {
	var arr []interface{}
	for item := range s.m {
		arr = append(arr, item)
	}
	return arr
}
