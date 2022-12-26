package set

import "testing"

func TestNew(t *testing.T) {
	s := New("Apple", "Blueberry", "Cherry", "Strawberry")

	if s.data == nil {
		t.Errorf("Must not be empty!")
	}
}

func TestAdd(t *testing.T) {
	s := New[int]()

	s.Add(-1)
	s.Add(1)
	s.Add(12)

	if s.data == nil {
		t.Errorf("Must not be empty!")
	}
}
