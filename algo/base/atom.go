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

// Atom represents a basic object that implements the Object interface.
type Atom struct {
	value interface{}
}

// NewAtom creates a new Atom with the given value.
func NewAtom(value interface{}) *Atom {
	return &Atom{value: value}
}

// Compare compares the Atom with another Object.
// It returns -1 if the Atom is less than the other Object, 0 if they are equal, and 1 if the Atom is greater.
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

// Equals checks if the Atom is equal to another Object.
func (a *Atom) Equals(other Object) bool {
	otherAtom, ok := other.(*Atom)
	if !ok {
		return false
	}
	return a.value == otherAtom.value
}

// Hash returns a hash code for the Atom.
func (a *Atom) Hash() int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", a.value)))
	return int(h.Sum32())
}

// String returns a string representation of the Atom.
func (a *Atom) String() string {
	return fmt.Sprintf("%v", a.value)
}

// Size returns the size of the Atom in bytes.
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
