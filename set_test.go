package set

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	s0 := New(0, 1, 2)
	s1 := New("Apple", "Blueberry", "Strawberry")

	if s0.data == nil || s1.data == nil {
		t.Errorf("")
	}
}

func TestAddRemove(t *testing.T) {
	s0 := New("Apple", "Blueberry", "Strawberry")

	s0.Remove("Apple")
	if !s0.Equals(New("Blueberry", "Strawberry")) {
		t.Errorf("")
	}

	s0.Add("Apple")
	s0.Remove("Blueberry")
	if !s0.Equals(New("Strawberry", "Apple")) {
		t.Errorf("")
	}

	s0.Add("Blueberry")
	s0.Remove("Strawberry")
	if !s0.Equals(New("Blueberry", "Apple")) {
		t.Errorf("")
	}

	s0.Remove("Blueberry")
	s0.Remove("Apple")
	if !s0.Equals(New[string]()) {
		t.Errorf("")
	}

	s1 := New(0, 1, 2)

	s1.Remove(0)
	if !s1.Equals(New(1, 2)) {
		t.Errorf("")
	}

	s1.Add(0)
	s1.Remove(1)
	if !s1.Equals(New(2, 0)) {
		t.Errorf("")
	}

	s1.Add(1)
	s1.Remove(2)
	if !s1.Equals(New(1, 0)) {
		t.Errorf("")
	}

	s1.Remove(0)
	s1.Remove(1)
	if !s1.Equals(New[int]()) {
		t.Errorf("")
	}
}

func FuzzAddRemoveString(f *testing.F) {
	// Test cases
	tcs0 := New("Apple", "Blueberry", "Strawberry")
	for tc := range tcs0.Elems() {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		s := New[string]()
		s.Add(orig)
		s.Remove(orig)
		if !s.Equals(New[string]()) {
			t.Errorf("")
		}
	})
}

func FuzzAddRemoveInt(f *testing.F) {
	tcs1 := New(0, 1, 2)
	for tc := range tcs1.Elems() {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig int) {
		s := New[int]()
		s.Add(orig)
		s.Remove(orig)
		if !s.Equals(New[int]()) {
			t.Errorf("")
		}
	})
}

func TestContainsAllAny(t *testing.T) {
	s0 := New("Apple", "Blueberry", "Strawberry")
	if !s0.Contains("Apple") || !s0.Contains("Blueberry") || !s0.Contains("Strawberry") {
		t.Errorf("")
	}
	if !s0.All("Apple", "Blueberry", "Strawberry") {
		t.Errorf("")
	}
	if !s0.Any("Apple", "Watermelon", "Strawberry") {
		t.Errorf("")
	}

	s1 := New(0, 1, 2)
	if !s1.Contains(0) || !s1.Contains(1) || !s1.Contains(2) {
		t.Errorf("")
	}
	if !s1.All(0, 1, 2) {
		t.Errorf("")
	}
	if !s1.Any(0, 2) {
		t.Errorf("")
	}
}

func TestUnion(t *testing.T) {
	s0, s1 := New("Apple", "Blueberry"), New("Strawberry")
	res0 := s0.Union(s1)
	if !res0.Equals(New("Apple", "Blueberry", "Strawberry")) {
		t.Errorf("")
	}

	s2, s3 := New(0, 2), New(1)
	res1 := s2.Union(s3)
	if !res1.Equals(New(0, 1, 2)) {
		t.Errorf("")
	}
}

func TestIntersection(t *testing.T) {
	s0, s1 := New("Apple", "Blueberry"), New("Strawberry", "Blueberry")
	res0 := s0.Intersection(s1)
	if !res0.Equals(New("Blueberry")) {
		t.Errorf("")
	}

	s2, s3 := New(0, 2), New(3)
	res1 := s2.Intersection(s3)
	if !res1.Equals(New[int]()) {
		t.Errorf("")
	}
}

func TestDifference(t *testing.T) {
	s0, s1 := New("Apple", "Blueberry", "Strawberry"), New("Strawberry")
	res0 := s0.Difference(s1)
	if !res0.Equals(New("Apple", "Blueberry")) {
		t.Errorf("")
	}

	s2, s3 := New(0, 2), New(1)
	res1 := s2.Difference(s3)
	if !res1.Equals(New(0, 2)) {
		t.Errorf("")
	}
}

func TestSymmetricDifference(t *testing.T) {
	s0, s1 := New("Apple", "Blueberry"), New("Strawberry")
	res0 := s0.SymmetricDifference(s1)
	if !res0.Equals(New("Apple", "Blueberry", "Strawberry")) {
		t.Errorf("")
	}

	s2, s3 := New(0, 2, 3), New(1, 2, 0)
	res1 := s2.SymmetricDifference(s3)
	if !res1.Equals(New(3, 1)) {
		t.Errorf("")
	}
}

func TestIsSubset(t *testing.T) {
	s0, s1 := New("Strawberry"), New("Apple", "Blueberry", "Strawberry")
	if !s0.IsSubset(s1) {
		t.Errorf("")
	}

	s2, s3 := New(0, 1, 2), New(2, 0)
	if !s3.IsSubset(s2) {
		t.Errorf("")
	}
}

func TestLen(t *testing.T) {
	s0, s1 := New("Strawberry"), New("Apple", "Blueberry", "Strawberry")
	s2, s3 := New(0, 1, 2), New(2, 0)
	if s0.Len() != 1 || s1.Len() != 3 || s2.Len() != 3 || s3.Len() != 2 {
		t.Errorf("")
	}
}

func TestClear(t *testing.T) {
	s0, s1 := New("Strawberry"), New("Apple", "Blueberry", "Strawberry")
	s2, s3 := New(0, 1, 2), New(2, 0)
	if s0.Len() != 1 || s1.Len() != 3 || s2.Len() != 3 || s3.Len() != 2 {
		t.Errorf("")
	}

	s0.Clear()
	s1.Clear()
	s2.Clear()
	s3.Clear()
	if s0.Len() != 0 || s1.Len() != 0 || s2.Len() != 0 || s3.Len() != 0 {
		t.Errorf("")
	}
}

func TestClone(t *testing.T) {
	s0, s1 := New("Strawberry"), New("Apple", "Blueberry", "Strawberry")
	s2, s3 := New(0, 1, 2), New(2, 0)

	s4, s5 := s0.Clone(), s1.Clone()
	s6, s7 := s2.Clone(), s3.Clone()

	if !(s0.Equals(s4) && s1.Equals(s5) && s2.Equals(s6) && s3.Equals(s7)) {
		t.Errorf("")
	}
}

func counter[T comparable](slice []T) map[T]int {
	counts := make(map[T]int)
	for _, elem := range slice {
		if _, ok := counts[elem]; ok {
			counts[elem] += 1
		} else {
			counts[elem] = 1
		}
	}
	return counts
}

func TestToSlice(t *testing.T) {
	s0, s1 := New("Strawberry"), New("Apple", "Blueberry", "Strawberry")
	s2, s3 := New(0, 1, 2), New(2, 0)

	s0s, s1s := s0.ToSlice(), s1.ToSlice()
	s2s, s3s := s2.ToSlice(), s3.ToSlice()

	if !reflect.DeepEqual(counter(s0s), counter([]string{"Strawberry"})) {
		t.Errorf("")
	}

	if !reflect.DeepEqual(counter(s1s), counter([]string{"Apple", "Blueberry", "Strawberry"})) {
		t.Errorf("")
	}

	if !reflect.DeepEqual(counter(s2s), counter([]int{0, 1, 2})) {
		t.Errorf("")
	}

	if !reflect.DeepEqual(counter(s3s), counter([]int{2, 0})) {
		t.Errorf("")
	}
}
