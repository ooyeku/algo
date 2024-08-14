package main

import (
	"fmt"
	"github.com/ooyeku/algo/algo"
	"github.com/ooyeku/algo/algo/searching"
	"github.com/ooyeku/algo/algo/sorting"
)

func main() {
	list := algo.GenerateList(100, 1, 100000000000)
	bs := sorting.BubbleSort(list)
	ms := sorting.MergeSort(list)
	qs := sorting.QuickSort(list)
	hs := sorting.HeapSort(list)
	ins := sorting.IntroSort(list)

	// convert []int to []interface{}
	bsInterface := make([]interface{}, len(bs))
	msInterface := make([]interface{}, len(ms))
	qsInterface := make([]interface{}, len(qs))
	hsInterface := make([]interface{}, len(hs))
	insInterface := make([]interface{}, len(ins))

	for i, v := range bs {
		bsInterface[i] = v
	}
	for i, v := range ms {
		msInterface[i] = v
	}
	for i, v := range qs {
		qsInterface[i] = v
	}
	for i, v := range hs {
		hsInterface[i] = v
	}
	for i, v := range ins {
		insInterface[i] = v
	}

	// compare lists
	fmt.Println(algo.CompareLists(bsInterface, msInterface, qsInterface, hsInterface, insInterface))

	listString := algo.GenerateListString(100, 65, 90)
	bsString := sorting.BubbleSortString(listString)
	msString := sorting.MergeSortString(listString)
	qsString := sorting.QuickSortString(listString)
	hsString := sorting.HeapSortString(listString)
	insString := sorting.IntroSortString(listString)

	bsStringInterface := make([]interface{}, len(bsString))
	msStringInterface := make([]interface{}, len(msString))
	qsStringInterface := make([]interface{}, len(qsString))
	hsStringInterface := make([]interface{}, len(hsString))
	insStringInterface := make([]interface{}, len(insString))

	for i, v := range bsString {
		bsStringInterface[i] = v
	}
	for i, v := range msString {
		msStringInterface[i] = v
	}
	for i, v := range qsString {
		qsStringInterface[i] = v
	}
	for i, v := range hsString {
		hsStringInterface[i] = v
	}
	for i, v := range insString {
		insStringInterface[i] = v
	}

	// compare lists
	fmt.Println(algo.CompareLists(bsStringInterface, msStringInterface, qsStringInterface, hsStringInterface, insStringInterface))

	listGeneric := algo.GenerateListGeneric(100, 1, 100000000000)
	bsGeneric := sorting.BubbleSortGeneric(listGeneric, func(i, j int) bool {
		return listGeneric[i].(int) < listGeneric[j].(int)
	})
	msGeneric := sorting.MergeSortGeneric(listGeneric, func(i, j int) bool {
		return listGeneric[i].(int) < listGeneric[j].(int)
	})
	qsGeneric := sorting.QuickSortGeneric(listGeneric, func(i, j int) bool {
		return listGeneric[i].(int) < listGeneric[j].(int)
	})
	hsGeneric := sorting.HeapSortGeneric(listGeneric, func(i, j int) bool {
		return listGeneric[i].(int) < listGeneric[j].(int)
	})
	insGeneric := sorting.IntroSortGeneric(listGeneric, func(i, j int) bool {
		return listGeneric[i].(int) < listGeneric[j].(int)
	})

	bsGenericInterface := make([]interface{}, len(bsGeneric))
	copy(bsGenericInterface, bsGeneric)
	msGenericInterface := make([]interface{}, len(msGeneric))
	copy(msGenericInterface, msGeneric)
	qsGenericInterface := make([]interface{}, len(qsGeneric))
	copy(qsGenericInterface, qsGeneric)
	hsGenericInterface := make([]interface{}, len(hsGeneric))
	copy(hsGenericInterface, hsGeneric)
	insGenericInterface := make([]interface{}, len(insGeneric))
	copy(insGenericInterface, insGeneric)

	// compare lists
	fmt.Println(algo.CompareLists(bsGenericInterface, msGenericInterface, qsGenericInterface, hsGenericInterface, insGenericInterface))

	// Example usage of BinarySearch
	sortedList := sorting.QuickSort(algo.GenerateList(100, 1, 1000))
	target := sortedList[50] // Choose a random element as the target
	index := searching.BinarySearch(sortedList, target)
	fmt.Printf("Binary Search: Target %d found at index %d\n", target, index)

	// Example usage of BinarySearchGeneric
	sortedStringList := sorting.QuickSortString(algo.GenerateListString(100, 65, 90))
	targetString := sortedStringList[25] // Choose a random element as the target
	indexString := searching.BinarySearchGeneric(sortedStringList, targetString, func(a, b string) bool {
		return a < b
	})
	fmt.Printf("Binary Search Generic: Target '%s' found at index %d\n", targetString, indexString)

	// Example usage of LinearSearch
	targetInt := sortedList[75] // Choose a random element as the target
	indexInt := searching.LinearSearch(sortedList, targetInt)
	fmt.Printf("Linear Search: Target %d found at index %d\n", targetInt, indexInt)

	// Example usage of LinearSearchString
	targetStringLinear := sortedStringList[75] // Choose a random element as the target
	indexStringLinear := searching.LinearSearchString(sortedStringList, targetStringLinear)
	fmt.Printf("Linear Search String: Target '%s' found at index %d\n", targetStringLinear, indexStringLinear)

	// Example usage of LinearSearchGeneric
	targetGeneric := listGeneric[75] // Choose a random element as the target
	indexGeneric := searching.LinearSearchGeneric(listGeneric, targetGeneric)
	fmt.Printf("Linear Search Generic: Target %v found at index %d\n", targetGeneric, indexGeneric)

	// Example usage of JumpSearch
	targetJump := sortedList[60] // Choose a random element as the target
	indexJump := searching.JumpSearch(sortedList, targetJump)
	fmt.Printf("Jump Search: Target %d found at index %d\n", targetJump, indexJump)

	// Example usage of JumpSearchGeneric
	targetJumpGeneric := sortedStringList[60] // Choose a random element as the target
	indexJumpGeneric := searching.JumpSearchGeneric(sortedStringList, targetJumpGeneric, func(a, b string) bool {
		return a < b
	})
	fmt.Printf("Jump Search Generic: Target '%s' found at index %d\n", targetJumpGeneric, indexJumpGeneric)

	// Example usage of CompareSearchAlgorithms
	sortedList = sorting.QuickSort(algo.GenerateList(1000, 1, 1_000_000_000))
	searchBenchmark := algo.CompareSearchAlgorithms(sortedList)
	fmt.Println("Search Benchmark Results:")
	for _, result := range searchBenchmark.Results {
		fmt.Printf("%s: Time: %v, Memory: %d bytes\n", result.Algorithm, result.Time, result.Memory)
	}
	fmt.Printf("Fastest: %s\n", searchBenchmark.Fastest)

	// Example usage of CompareSearchAlgorithmsGeneric
	unsortedGenericList := algo.GenerateListGeneric(1_000_000, 1, 1_000_000_000)
	sortedGenericList := sorting.QuickSortGeneric(unsortedGenericList, func(i, j int) bool {
		return unsortedGenericList[i].(int) < unsortedGenericList[j].(int)
	})
	target = sortedGenericList[500000].(int)
	searchBenchmarkGeneric := algo.CompareSearchAlgorithmsGeneric(sortedGenericList, target, func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
	fmt.Println("\nGeneric Search Benchmark Results:")
	for _, result := range searchBenchmarkGeneric.Results {
		fmt.Printf("%s: Time: %v, Memory: %d bytes\n", result.Algorithm, result.Time, result.Memory)
	}
	fmt.Printf("Fastest: %s\n", searchBenchmarkGeneric.Fastest)

	// Example usage of CompareSortAlgorithms
	sortedList = sorting.QuickSort(algo.GenerateList(100000, 1, 1000000))
	sortBenchmark := algo.CompareSortAlgorithms(sortedList)
	fmt.Println("\nSort Benchmark Results:")
	for _, result := range sortBenchmark.Results {
		fmt.Printf("%s: Time: %v, Memory: %d bytes\n", result.Algorithm, result.Time, result.Memory)
	}
	fmt.Printf("Fastest: %s\n", sortBenchmark.Fastest)
	fmt.Printf("Most Memory Efficient: %s\n", sortBenchmark.MostMemoryEfficient)
}
