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

// CompareLists compares two or more lists and returns true if they are equal.
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

// CompareListsGeneric compares two or more lists and returns true if they are equal.
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

func less(i, j int) bool {
	return i < j
}