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
	"hash/fnv"
	"math/big"
)

// NumberType is an enumeration to differentiate between integer and float types.
type NumberType int

const (
	Integer NumberType = iota
	Float
)

// Number represents a numeric value that can be either an integer or a float.
type Number struct {
	intValue   *big.Int
	floatValue *big.Float
	numType    NumberType
}

// NewInteger creates a new Number with an integer value.
func NewInteger(value int64) *Number {
	return &Number{
		intValue: big.NewInt(value),
		numType:  Integer,
	}
}

// NewFloat creates a new Number with a float value.
func NewFloat(value float64) *Number {
	return &Number{
		floatValue: big.NewFloat(value),
		numType:    Float,
	}
}

// Compare compares the Number with another Object.
// It returns -1 if the Number is less than the other Object, 0 if they are equal, and 1 if the Number is greater.
func (n *Number) Compare(other Object) int {
	otherNumber, ok := other.(*Number)
	if !ok {
		panic("Cannot compare Number with non-Number object")
	}

	switch n.numType {
	case Integer:
		if otherNumber.numType == Integer {
			return n.intValue.Cmp(otherNumber.intValue)
		}
		nFloat := new(big.Float).SetInt(n.intValue)
		return nFloat.Cmp(otherNumber.floatValue)
	case Float:
		if otherNumber.numType == Float {
			return n.floatValue.Cmp(otherNumber.floatValue)
		}
		otherFloat := new(big.Float).SetInt(otherNumber.intValue)
		return n.floatValue.Cmp(otherFloat)
	default:
		panic("Unsupported type for comparison")
	}
}

// Equals checks if the Number is equal to another Object.
func (n *Number) Equals(other Object) bool {
	otherNumber, ok := other.(*Number)
	if !ok {
		return false
	}

	switch n.numType {
	case Integer:
		if otherNumber.numType == Integer {
			return n.intValue.Cmp(otherNumber.intValue) == 0
		}
		nFloat := new(big.Float).SetInt(n.intValue)
		return nFloat.Cmp(otherNumber.floatValue) == 0
	case Float:
		if otherNumber.numType == Float {
			return n.floatValue.Cmp(otherNumber.floatValue) == 0
		}
		otherFloat := new(big.Float).SetInt(otherNumber.intValue)
		return n.floatValue.Cmp(otherFloat) == 0
	default:
		panic("Unsupported type for equality check")
	}
}

// Hash returns a hash code for the Number.
func (n *Number) Hash() int {
	h := fnv.New32a()
	h.Write([]byte(n.String()))
	return int(h.Sum32())
}

// String returns a string representation of the Number.
func (n *Number) String() string {
	switch n.numType {
	case Integer:
		return n.intValue.String()
	case Float:
		return n.floatValue.String()
	default:
		panic("Unsupported type for string conversion")
	}
}

// Size returns the size of the Number in bytes.
func (n *Number) Size() int {
	switch n.numType {
	case Integer:
		return len(n.intValue.Bytes())
	case Float:
		return len(n.floatValue.Text('g', -1))
	default:
		panic("Unsupported type for size calculation")
	}
}

// IsInteger checks if the Number is an integer.
func (n *Number) IsInteger() bool {
	return n.numType == Integer
}

// ToAtom converts the Number to an Atom if it is an integer.
func (n *Number) ToAtom() *Atom {
	if n.IsInteger() {
		return NewAtom(n.intValue.Int64())
	}
	panic("Cannot convert a float Number to an Atom")
}
