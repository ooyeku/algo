package algo

import "math/rand"

// GenerateList generates a list of random integers with the specified length, minimum value, and maximum value.
func GenerateList(length int, min int, max int) []int {
	list := make([]int, length)
	for i := 0; i < length; i++ {
		list[i] = rand.Intn(max-min) + min
	}
	return list
}

// GenerateListString takes in the length, minimum value, and maximum value and returns
// a list of strings with random values between the minimum and maximum values.
// The length parameter specifies the size of the list to be generated.
// The min parameter specifies the minimum value for the generated strings.
// The max parameter specifies the maximum value for the generated strings.
// The returned list will contain random strings between the specified minimum and maximum values.
func GenerateListString(length int, min int, max int) []string {
	list := make([]string, length)
	for i := 0; i < length; i++ {
		list[i] = string(rune(rand.Intn(max-min) + min))
	}
	return list
}

// GenerateListGeneric generates a list of specified length containing random integers
// between the specified minimum and maximum values. The list is of type []interface{},
// allowing it to hold any type of value. The length parameter specifies the number of
// elements in the list. The min parameter specifies the minimum value that can be
// generated, inclusive. The max parameter specifies the maximum value that can be
// generated, exclusive. The function returns the generated list.
func GenerateListGeneric(length int, min int, max int) []interface{} {
	list := make([]interface{}, length)
	for i := 0; i < length; i++ {
		list[i] = rand.Intn(max-min) + min
	}
	return list
}
