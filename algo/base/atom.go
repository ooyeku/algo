// Copyright 2024 olayeku
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base

import (
	"fmt"
	"hash/fnv"
)

// Atom represents a generic atomic value that can be compared, hashed,
// converted to a string, and determine its size.
//
// The value of Atom is stored as an interface{} type, allowing it to hold
// values of any type.
type Atom struct {
	value interface{}
}

// NewAtom creates a new Atom object initialized with the given value.
func NewAtom(value interface{}) *Atom {
	return &Atom{value: value}
}

// Compare compares the Atom object with another object and returns an integer
// representing their relationship.
//
// The provided 'other' object must be of type *Atom. If it is not, a panic is
// triggered with the message "Cannot compare Atom with non-Atom object".
//
// The comparison is performed based on the value type of the Atom's value.
// - If the value is an int, the comparison is done numerically:
//   - If the Atom's value is less than the other Atom's value, -1 is returned.
//   - If the Atom's value is greater than the other Atom's value, 1 is returned.
//   - If the Atom's value is equal to the other Atom's value, 0 is returned.
//   - If the value is a float64, the comparison is done numerically in the same
//     way as for int values.
//   - If the value is a string, the comparison is done lexicographically:
//   - If the Atom's value is lexicographically less than the other Atom's value,
//     -1 is returned.
//   - If the Atom's value is lexicographically greater than the other Atom's
//     value, 1 is returned.
//   - If the Atom's value is equal to the other Atom's value, 0 is returned.
//   - If the value is of any other type, a panic is triggered with the message
//     "Unsupported type for comparison".
func (a *Atom) Compare(other Object) int {
	otherAtom, ok := other.(*Atom)
	if !ok {
		panic("Cannot compare Atom with non-Atom object")
	}

	switch v := a.value.(type) {
	case int:
		ov := otherAtom.value.(int)
		if v < ov {
			return -1
		} else if v > ov {
			return 1
		}
		return 0
	case float64:
		ov := otherAtom.value.(float64)
		if v < ov {
			return -1
		} else if v > ov {
			return 1
		}
		return 0
	case string:
		ov := otherAtom.value.(string)
		if v < ov {
			return -1
		} else if v > ov {
			return 1
		}
		return 0
	default:
		panic("Unsupported type for comparison")
	}
}

// Equals checks if the current Atom object is equal to the given object.
// It returns true if the objects are equal and false otherwise.
// The comparison is based on the value field of the Atom objects.
// The given object must be of type Atom, otherwise it returns false.
// This method is case-sensitive for string comparison.
// Example usage can be found in the TestAtom_Equals function.
func (a *Atom) Equals(other Object) bool {
	otherAtom, ok := other.(*Atom)
	if !ok {
		return false
	}
	return a.value == otherAtom.value
}

// Hash returns the hash code for the Atom value. It computes the hash value by
// converting the value to a string and then hashing the string using the
// FNV-1a 32-bit algorithm. The resulting hash code is converted to an int and
// returned.
func (a *Atom) Hash() int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", a.value)))
	return int(h.Sum32())
}

// String returns a string representation of the Atom's value.
func (a *Atom) String() string {
	return fmt.Sprintf("%v", a.value)
}

// Size returns the size in bytes of the Atom's value.
// If the value is an int, it assumes a size of 8 bytes (assuming 64-bit int).
// If the value is a float64, it assumes a size of 8 bytes (assuming 64-bit float).
// If the value is a string, it returns the length of the string.
// If the value is of any other type, it panics with the message "Unsupported type for size calculation".
func (a *Atom) Size() int {
	switch v := a.value.(type) {
	case int:
		return 8 // assuming 64-bit int
	case float64:
		return 8 // assuming 64-bit float
	case string:
		return len(v)
	default:
		panic("Unsupported type for size calculation")
	}
}
