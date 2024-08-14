package base

import (
	"testing"
)

func TestAtom_Compare(t *testing.T) {
	type testCase struct {
		a, b         *Atom
		expectedComp int
	}

	table := []testCase{
		{NewAtom(5), NewAtom(3), 1},
		{NewAtom(3), NewAtom(5), -1},
		{NewAtom(5), NewAtom(5), 0},
		{NewAtom("abc"), NewAtom("aba"), 1},
		{NewAtom("aba"), NewAtom("abc"), -1},
		{NewAtom("abc"), NewAtom("abc"), 0},
		{NewAtom(2.5), NewAtom(2.3), 1},
		{NewAtom(2.3), NewAtom(2.5), -1},
		{NewAtom(2.5), NewAtom(2.5), 0},
	}

	for _, c := range table {
		if got := c.a.Compare(c.b); got != c.expectedComp {
			t.Errorf("Unexpected comparison result between %v and %v, got: %v, want: %v",
				c.a, c.b, got, c.expectedComp)
		}
	}
}

func TestAtom_Equals(t *testing.T) {
	type testCase struct {
		a, b    *Atom
		isEqual bool
	}

	table := []testCase{
		{NewAtom(5), NewAtom(3), false},
		{NewAtom(5), NewAtom(5), true},
		{NewAtom("abc"), NewAtom("aba"), false},
		{NewAtom("abc"), NewAtom("abc"), true},
		{NewAtom(2.5), NewAtom(2.3), false},
		{NewAtom(2.5), NewAtom(2.5), true},
	}

	for _, c := range table {
		if got := c.a.Equals(c.b); got != c.isEqual {
			t.Errorf("Unexpected equality result between %v and %v, got: %v, want: %v",
				c.a, c.b, got, c.isEqual)
		}
	}
}

func TestAtom_Hash(t *testing.T) {
	a := NewAtom(5)
	b := NewAtom(5)
	if a.Hash() != b.Hash() {
		t.Errorf("Hash function is not consistent for same value")
	}
}

func TestAtom_String(t *testing.T) {
	a := NewAtom("test")
	if a.String() != "test" {
		t.Errorf("String function is not consistent with value: got %v want %v", a.String(), "test")
	}
}

func TestAtom_Size(t *testing.T) {
	type testCase struct {
		a    *Atom
		size int
	}

	table := []testCase{
		{NewAtom(5), 8},
		{NewAtom(3.14), 8},
		{NewAtom("abc"), 3},
	}

	for _, c := range table {
		if got := c.a.Size(); got != c.size {
			t.Errorf("Unexpected size result for %v, got: %v, want: %v",
				c.a, got, c.size)
		}
	}
}
