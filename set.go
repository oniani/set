// Set is a very useful abstraction, yet it is not supported by Go. This package implements the set
// based on `map` data structure already defined in the language. In addition to the standard set
// operations, functions for mapping and filtering are also provided. The implementation seeks to be
// minimal, yet productive.
package set

import "reflect"

// Set represents a collection of unique elements
type Set[T comparable] struct {
	data map[T]struct{}
}

// New returns a new set containing the provided elements.
//
// Complexity: O(n), where n is the number of provided elements.
func New[T comparable](elems ...T) Set[T] {
	data := make(map[T]struct{})
	for _, elem := range elems {
		data[elem] = struct{}{}
	}
	return Set[T]{data: data}
}

// Add adds an element to the set.
//
// Complexity: O(1).
func (s *Set[T]) Add(elem T) {
	s.data[elem] = struct{}{}
}

// Remove removes an element from the set.
//
// Complexity: O(1).
func (s *Set[T]) Remove(elem T) {
	delete(s.data, elem)
}

// Contains returns true if the set contains the provided element, false otherwise.
//
// Complexity: O(1).
func (s *Set[T]) Contains(elem T) bool {
	_, ok := s.data[elem]
	return ok
}

// All returns true if the set contains all provided elements, false otherwise.
//
// Complexity: O(n), where n is the number of elements.
func (s *Set[T]) All(elems ...T) bool {
	for _, elem := range elems {
		if !s.Contains(elem) {
			return false
		}
	}
	return true
}

// Any returns true if the set contains any of the provided elements, false otherwise.
//
// Complexity: O(n), where n is the number of elements.
func (s *Set[T]) Any(elems ...T) bool {
	for _, elem := range elems {
		if s.Contains(elem) {
			return true
		}
	}
	return false
}

// Elems returns set values to iterate over.
//
// Complexity: O(1).
func (s *Set[T]) Elems() map[T]struct{} {
	return s.data
}

// Union returns the union with the other set as a new set.
//
// Complexity: O(n + m), where n and m are numbers of elements in sets s and o, respectively.
func (s *Set[T]) Union(o Set[T]) Set[T] {
	res := New[T]()
	for elem := range s.Elems() {
		res.data[elem] = struct{}{}
	}
	for elem := range o.Elems() {
		res.data[elem] = struct{}{}
	}
	return res
}

// Intersect is a helper function that returns the intersection of two sets.
func intersect[T comparable](s Set[T], o Set[T]) Set[T] {
	res := New[T]()
	for elem := range s.Elems() {
		if o.Contains(elem) {
			res.data[elem] = struct{}{}
		}
	}
	return res
}

// Intersection returns the intersection with the other set as a new set.
//
// Complexity: O(min(n, m)), where n and m are numbers of elements in sets s and o, respectively.
func (s *Set[T]) Intersection(o Set[T]) Set[T] {
	if s.Len() < o.Len() {
		return intersect(*s, o)
	}
	return intersect(o, *s)
}

// Difference returns the difference with the other set as a new set.
//
// Complexity: O(m), where m is the number of elements in the other set.
func (s *Set[T]) Difference(o Set[T]) Set[T] {
	res := New[T]()
	for elem := range s.Elems() {
		if !o.Contains(elem) {
			res.data[elem] = struct{}{}
		}
	}
	return res
}

// SymmetricDifference returns the symmetric difference with the other set as a new set.
//
// Complexity: O(n + m), where n and m are numbers of elements in sets s and o, respectively.
func (s *Set[T]) SymmetricDifference(o Set[T]) Set[T] {
	res := s.Union(o)
	return res.Difference(s.Intersection(o))
}

// IsSubset returns true if another set contains the set, false otherwise
//
// Complexity: O(n), where n is the number of elements in the set.
func (s *Set[T]) IsSubset(o Set[T]) bool {
	for elem := range s.Elems() {
		if !o.Contains(elem) {
			return false
		}
	}
	return true
}

// Len returns the number of elements in the set.
//
// Complexity: O(1).
func (s *Set[T]) Len() int {
	return len(s.data)
}

// Clear removes all elements from the set.
//
// Complexity: O(1).
func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

// Returns a deep copy of the set.
//
// Complexity: O(n), where n is the number of elements in the set.
func (s *Set[T]) Clone() Set[T] {
	res := Set[T]{make(map[T]struct{})}
	for elem := range s.Elems() {
		res.data[elem] = struct{}{}
	}
	return res
}

// Returns true if the set equals to the other set, false otherwise.
//
// Complexity: O(1).
func (s *Set[T]) Equals(o Set[T]) bool {
	return reflect.DeepEqual(s.data, o.data)
}

// ToSlice returns a slice containing the elements of the set.
//
// Complexity: O(n), where n is the number of elements in the set.
func (s *Set[T]) ToSlice() []T {
	res, idx := make([]T, s.Len()), 0
	for elem := range s.Elems() {
		// Faster than doing res = append(res, elem)
		res[idx] = elem
		idx += 1
	}
	return res
}
