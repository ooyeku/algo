package algo

import "math/rand"

func GenerateList(length int, min int, max int) []int {
	list := make([]int, length)
	for i := 0; i < length; i++ {
		list[i] = rand.Intn(max-min) + min
	}
	return list
}

func GenerateListString(length int, min int, max int) []string {
	list := make([]string, length)
	for i := 0; i < length; i++ {
		list[i] = string(rune(rand.Intn(max-min) + min))
	}
	return list
}

func GenerateListGeneric(length int, min int, max int) []interface{} {
	list := make([]interface{}, length)
	for i := 0; i < length; i++ {
		list[i] = rand.Intn(max-min) + min
	}
	return list
}
