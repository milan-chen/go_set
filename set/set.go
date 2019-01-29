package set

type Set interface {
	Add(e interface{}) bool
	Clear()
	Has(e interface{}) bool
	Equals(other Set) bool
	IsEmpty() bool
	Remove(e interface{})
	Len() int
	Elements() []interface{}
	ToString() string
	ToArray() []interface{}
}

// Add Set b to Set a.
func AddSet(a Set, b Set) {
	if nil == a || nil == b || 0 == b.Len() {
		return
	}
	for _, v := range b.Elements() {
		a.Add(v)
	}
}

// Determine if the set a is a superset of the set b.
func IsSuperset(a Set, b Set) bool {
	if nil == a || b == nil {
		return false
	}
	if 0 == a.Len() || a.Len() <= b.Len() {
		return false
	}
	if 0 < a.Len() && 0 == b.Len() {
		return true
	}
	for _, v := range b.Elements() {
		if !b.Has(v) {
			return false
		}
	}
	return true
}

// Return the intersection of set a and set b
func Intersect(a Set, b Set) Set {
	if nil == a || nil == b {
		return nil
	}
	hashSet := NewSimpleSet()
	if 0 == a.Len() || 0 == b.Len() {
		return hashSet
	}
	if a.Len() < b.Len() {
		for _, v := range a.Elements() {
			if a.Has(v) {
				hashSet.Add(v)
			}
		}
	} else {
		for _, v := range b.Elements() {
			if a.Has(v) {
				hashSet.Add(v)
			}
		}
	}
	return hashSet
}

// Return the union of the Set a and the Set b.
func Union(a Set, b Set) Set {
	if  nil == a || nil == b {
		return nil
	}
	hashSet := NewSimpleSet()
	AddSet(hashSet, a)
	AddSet(hashSet, b)
	return hashSet
}

// Return the difference between the Set a and the Set b.
func Difference(a Set, b Set) Set {
	if nil == a {
		return nil
	}
	hashSet := NewSimpleSet()
	if nil == b || b.Len() == 0 {
		AddSet(hashSet, a)
		return hashSet
	}
	for _, v := range a.Elements() {
		if !b.Has(v) {
			hashSet.Add(v)
		}
	}
	return hashSet
}

// Return the symmetric difference set of the Set a and the Set b.
func SymmetricDifference(a Set, b Set) Set {
	diffA := Difference(a, b)
	if nil == b || 0 == b.Len() {
		return diffA
	}
	diffB := Difference(b, a)
	return Union(diffA, diffB)
}

// Determine if a given e is a collection and return bool result.
func IsSet(e interface{}) bool {
	if _, ok := e.(Set); ok {
		return true
	}
	return false
}

// Return the example of HashSet.
func NewSimpleSet() Set {
	return NewHashSet()
}