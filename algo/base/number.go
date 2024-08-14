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

// NumberType is a custom type representing different types of numbers. It has two constants:
// - Integer: Represents an integer number type.
// - Float: Represents a floating-point number type.
// This type is used in conjunction with the Number struct to perform mathematical operations and comparisons
// on numeric values. See the Number struct documentation for more details.
type NumberType int

// Integer and Float are two constants of type NumberType.
// Integer represents an integer number type.
// Float represents a floating-point number type.
const (
	Integer NumberType = iota
	Float
)

// Number represents a numeric value that can be either an integer or a float.
//
// Fields:
// - intValue: A pointer to a big.Int value representing the integer part of the number.
// - floatValue: A pointer to a big.Float value representing the floating-point part of the number.
// - numType: The type of the number, which can be either Integer or Float.
//
// This type is used to perform mathematical operations and comparisons on numeric values.
// It provides methods for comparison, equality check, hashing, string representation, size calculation,
// checking whether the number is an integer, and converting the number to an Atom value.
//
// It also requires the following types for usage:
//   - NumberType: A type declaration representing the type of a Number (integer or float).
//   - Object: An interface representing common methods for objects, including comparison, equality check, hashing,
//     string representation, and size calculation.
//
// See the provided code examples for usage of this type.
type Number struct {
	intValue   *big.Int
	floatValue *big.Float
	numType    NumberType
}

// NewInteger returns a new Number object with the specified value as an integer.
func NewInteger(value int64) *Number {
	return &Number{
		intValue: big.NewInt(value),
		numType:  Integer,
	}
}

// NewFloat creates a new Number object with the specified float64 value.
// The Number object represents a floating-point number and is used for
// mathematical operations. The value parameter is the float64 value to
// initialize the Number object with. The returned Number object is a
// pointer to the newly created object.
func NewFloat(value float64) *Number {
	return &Number{
		floatValue: big.NewFloat(value),
		numType:    Float,
	}
}

// Compare compares the current Number with another Object and returns an integer
// result based on the comparison. The method panics if the other Object is not a Number.
// If both Numbers are of type Integer, their integer values are compared using the
// Cmp method of big.Int. If one of the Numbers is a Float, both Numbers are converted
// to big.Float and compared using the Cmp method. If an unsupported type is encountered,
// the method panics. The returned value is:
// -1 if the current Number is less than the other Object
// 0 if the current Number is equal to the other Object
// 1 if the current Number is greater than the other Object.
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

// Equals checks if the current Number object is equal to the given Object. The method compares the values of the
// Number objects based on their numeric types (Integer or Float) and returns true if they are equal,
// false otherwise. If the given object is not of type *Number, the method returns false.
//
// Numeric types are compared using the Cmp method of the big.Int and big.Float types. If the numeric types of the
// two objects are the same (both Integer or both Float), the Cmp method is used directly. If the numeric types
// are different, the method converts the integer value to a big.Float and then performs the comparison.
//
// If the numeric types of both objects are Integer, the method compares the values by calling the Cmp method of the
// big.Int type and returns true if the comparison result is zero.
//
// If the numeric types of both objects are Float, the method compares the values by calling the Cmp method of the
// big.Float type and returns true if the comparison result is zero.
//
// If the numeric types of the two objects are different, the method converts the integer value of the first object
// to a big.Float and compares it with the second object's float value using the Cmp method of the big.Float type. If
// the comparison result is zero, the method returns true.
//
// If the numeric types are unsupported (neither Integer nor Float), the method panics with the message "Unsupported
// type for equality check".
//
// The method uses type assertion to check if the given object is of type *Number. If the assertion fails, the method
// returns false.
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

// Hash calculates the hash value for the Number object.
// It uses the FNV-1a hash algorithm to create the hash value.
// The hash value is an integer representation of the Number object.
// The hash value is based on the string representation of the Number obtained with the String method.
// The returned hash value is an int.
func (n *Number) Hash() int {
	h := fnv.New32a()
	h.Write([]byte(n.String()))
	return int(h.Sum32())
}

// String returns the string representation of the Number. If the Number
// represents an integer, it returns the integer value as a string. If the
// Number represents a floating-point value, it returns the float value as a
// string. If the Number represents an unsupported type, it panics.
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

// Size returns the size of the Number object. The size of the Number object
// depends on its internal representation. If the number is of type Integer,
// the size is the length of its byte representation. If the number is of type
// Float, the size is the length of its string representation using the 'g'
// format with unlimited precision.
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

// IsInteger returns true if the Number object represents an integer value, and false otherwise.
// It checks the numType field of the Number object. If the numType is set to Integer, it means
// that the Number object represents an integer value and returns true. Otherwise, it returns false.
//
// Note that the Number object can represent either an integer value or a floating-point value.
// The numType field of the Number object is used to indicate the type of value it represents.
//
// Example usage can be found in the TestNewInteger and TestNewFloat functions.
func (n *Number) IsInteger() bool {
	return n.numType == Integer
}

// ToAtom converts the Number to an Atom. If the Number represents an integer value, it creates a new Atom with the
// integer value obtained from converting the intValue field to an int64. If the Number represents a floating-point
// value, it panics with the message "Cannot convert a float Number to an Atom".
func (n *Number) ToAtom() *Atom {
	if n.IsInteger() {
		return NewAtom(n.intValue.Int64())
	}
	panic("Cannot convert a float Number to an Atom")
}
