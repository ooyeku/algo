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

import (
	"github.com/ooyeku/algo/algo/sorting"
	"runtime"
	"time"
)

// SortResult represents the result of a sorting operation. It contains information about the algorithm used,
// the time taken to sort, and the memory used during the sorting operation.
type SortResult struct {
	Algorithm string
	Time      time.Duration
	Memory    uint64
}

// SortBenchmark represents the results of benchmarking different sorting algorithms.
// It stores the sorting results, the size of the input list, and the name of the fastest algorithm.
type SortBenchmark struct {
	Results  []SortResult
	ListSize int
	Fastest  string
}

// CompareSortAlgorithms benchmarks multiple sorting algorithms on a given list and returns a SortBenchmark
// containing information about the results. The function compares the performance of Bubble Sort, Merge Sort,
// Quick Sort, Heap Sort, and Intro Sort. It measures the time taken by each algorithm and the memory usage.
// The fastest sort algorithm is determined based on the time taken. The function returns a SortBenchmark
// struct with the list size, results of the benchmarks, and the name of the fastest algorithm.
// SortBenchmark is a struct that contains results of sorting benchmarks, the size of the list, and the name
// of the fastest sort algorithm.
func CompareSortAlgorithms(list []int) SortBenchmark {
	benchmark := SortBenchmark{
		ListSize: len(list),
	}

	// Benchmark Bubble Sort
	benchmark.Results = append(benchmark.Results, benchmarkSort("Bubble Sort", func() []int {
		return sorting.BubbleSort(append([]int(nil), list...))
	}))

	// Benchmark Merge Sort
	benchmark.Results = append(benchmark.Results, benchmarkSort("Merge Sort", func() []int {
		return sorting.MergeSort(append([]int(nil), list...))
	}))

	// Benchmark Quick Sort
	benchmark.Results = append(benchmark.Results, benchmarkSort("Quick Sort", func() []int {
		return sorting.QuickSort(append([]int(nil), list...))
	}))

	// Benchmark Heap Sort
	benchmark.Results = append(benchmark.Results, benchmarkSort("Heap Sort", func() []int {
		return sorting.HeapSort(append([]int(nil), list...))
	}))

	// Benchmark Intro Sort
	benchmark.Results = append(benchmark.Results, benchmarkSort("Intro Sort", func() []int {
		return sorting.IntroSort(append([]int(nil), list...))
	}))

	// get the fastest sort algorithm
	fastest := benchmark.Results[0]
	for _, result := range benchmark.Results {
		if result.Time < fastest.Time {
			fastest = result
		}
	}
	benchmark.Fastest = fastest.Algorithm

	return benchmark
}

// benchmarkSort is a function that measures the time and memory usage of a sorting algorithm.
// The function takes a name (string) and a sortFunc (function that returns a []int) as parameters.
// It measures the memory usage before and after executing the sortFunc, and calculates the duration of the execution.
// The function returns a SortResult struct that contains the algorithm name, execution time, and memory usage.
func benchmarkSort(name string, sortFunc func() []int) SortResult {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memBefore := m.Alloc

	start := time.Now()
	sortFunc()
	duration := time.Since(start)

	runtime.ReadMemStats(&m)
	memAfter := m.Alloc

	return SortResult{
		Algorithm: name,
		Time:      duration,
		Memory:    memAfter - memBefore,
	}
}

// CompareSortAlgorithmsGeneric benchmarks several generic sorting algorithms on a given list.
// It takes a list of values and a comparison function as input and returns a SortBenchmark struct
// containing the results of the benchmarking.
//
// The benchmark compares the following sorting algorithms:
// - Bubble Sort Generic
// - Merge Sort Generic
// - Quick Sort Generic
// - Heap Sort Generic
// - Intro Sort Generic
//
// The fastest sorting algorithm is determined based on the execution time and stored in the Fastest field of the SortBenchmark struct.
//
// Each sorting algorithm is benchmarked using the benchmarkSortGeneric function, which measures the execution time
// and memory consumption of a given sorting function.
//
// The SortResult struct contains information about the sorting algorithm, execution time, and memory consumption.
//
// The benchmarkSortGeneric function utilizes the runtime and time packages to measure the memory consumption and execution time.
//
// The sorting functions used in the benchmark (BubbleSortGeneric, MergeSortGeneric, QuickSortGeneric, HeapSortGeneric, IntroSortGeneric)
// implement the corresponding sorting algorithms with generic support for different types of lists.
//
// The mergeGeneric, quickSortGeneric, heapifyGeneric, and partitionGeneric functions are helper functions used by the sorting algorithms.
//
// The SortBenchmark struct contains the Results field, which is a slice of SortResult structs storing the results for each sorting algorithm.
// The ListSize field represents the size of the original list that was sorted.
//
// Example usage:
// list := []interface{}{5, 3, 8, 2, 1}
//
//	less := func(i, j int) bool {
//	    return list[i].(int) < list[j].(int)
//	}
//
// benchmark := CompareSortAlgorithmsGeneric(list, less)
// fmt.Println(benchmark.Fastest) // Output: Bubble Sort Generic
// fmt.Println(benchmark.ListSize) // Output: 5
// fmt.Println(benchmark.Results[0].Algorithm) // Output: Bubble Sort Generic
// fmt.Println(benchmark.Results[0].Time) // Output: 1.042µs
// fmt.Println(benchmark.Results[0].Memory) // Output: 0
func CompareSortAlgorithmsGeneric(list []interface{}, less func(i, j int) bool) SortBenchmark {
	benchmark := SortBenchmark{
		ListSize: len(list),
	}

	// Benchmark Bubble Sort Generic
	benchmark.Results = append(benchmark.Results, benchmarkSortGeneric("Bubble Sort Generic", func() []interface{} {
		return sorting.BubbleSortGeneric(append([]interface{}(nil), list...), less)
	}))

	// Benchmark Merge Sort Generic
	benchmark.Results = append(benchmark.Results, benchmarkSortGeneric("Merge Sort Generic", func() []interface{} {
		return sorting.MergeSortGeneric(append([]interface{}(nil), list...), less)
	}))

	// Benchmark Quick Sort Generic
	benchmark.Results = append(benchmark.Results, benchmarkSortGeneric("Quick Sort Generic", func() []interface{} {
		return sorting.QuickSortGeneric(append([]interface{}(nil), list...), less)
	}))

	// Benchmark Heap Sort Generic
	benchmark.Results = append(benchmark.Results, benchmarkSortGeneric("Heap Sort Generic", func() []interface{} {
		return sorting.HeapSortGeneric(append([]interface{}(nil), list...), less)
	}))

	// Benchmark Intro Sort Generic
	benchmark.Results = append(benchmark.Results, benchmarkSortGeneric("Intro Sort Generic", func() []interface{} {
		return sorting.IntroSortGeneric(append([]interface{}(nil), list...), less)
	}))

	// get the fastest sort algorithm
	fastest := benchmark.Results[0]
	for _, result := range benchmark.Results {
		if result.Time < fastest.Time {
			fastest = result
		}
	}
	benchmark.Fastest = fastest.Algorithm

	return benchmark
}

// benchmarkSortGeneric takes the name of a sorting algorithm and a sorting function, and benchmarks the performance
// of the sorting function. It measures the time taken to sort the data, the memory allocated before and after sorting,
// and returns a SortResult struct with the algorithm name, time, and memory information.
func benchmarkSortGeneric(name string, sortFunc func() []interface{}) SortResult {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memBefore := m.Alloc

	start := time.Now()
	sortFunc()
	duration := time.Since(start)

	runtime.ReadMemStats(&m)
	memAfter := m.Alloc

	return SortResult{
		Algorithm: name,
		Time:      duration,
		Memory:    memAfter - memBefore,
	}
}
