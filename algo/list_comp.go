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

package algo

// CompareLists compares multiple lists and returns true if all the lists have the same elements in the same order,
// and false otherwise. The lists should be passed as variadic arguments, with each list represented as []interface{}.
// If no lists are provided, the function returns true. The function uses a nested loop to compare the elements of each
// list. If the lengths of the lists are different, the function immediately returns false. If any of the elements at the
// same index in different lists are different, the function also returns false. If all the elements in all the lists
// are equal, the function returns true.
func CompareLists(lists ...[]interface{}) bool {
	if len(lists) == 0 {
		return true
	}
	for i := 1; i < len(lists); i++ {
		if len(lists[i]) != len(lists[0]) {
			return false
		}
	}
	for i := range lists[0] {
		for j := 1; j < len(lists); j++ {
			if lists[j][i] != lists[0][i] {
				return false
			}
		}
	}
	return true
}

// CompareListsGeneric compares multiple lists of interfaces and returns true if they are the same.
// It first checks if all the lists have the same length, and if not, it returns false.
// Then it iterates through each element of the first list and compares it with the corresponding
// elements in the other lists. It uses the "less" function to compare the elements, so the elements
// must be of type int. If any comparison fails, it returns false. If all comparisons pass, it returns true.
func CompareListsGeneric(lists ...[]interface{}) bool {
	if len(lists) == 0 {
		return true
	}
	for i := 1; i < len(lists); i++ {
		if len(lists[i]) != len(lists[0]) {
			return false
		}
	}
	for i := range lists[0] {
		for j := 1; j < len(lists); j++ {
			val1, ok1 := lists[j][i].(int)
			val2, ok2 := lists[0][i].(int)
			if !ok1 || !ok2 || !less(val1, val2) {
				return false
			}
		}
	}
	return true
}

// less is a function that compares two integers and returns true if the first integer is less than the second integer.
func less(i, j int) bool {
	return i < j
}
