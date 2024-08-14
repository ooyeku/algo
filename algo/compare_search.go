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
	"github.com/ooyeku/algo/algo/searching"
	"math/rand"
	"runtime"
	"time"
)

// SearchResult represents the result of a search operation.
type SearchResult struct {
	Algorithm string
	Target    interface{}
	Index     int
	Time      time.Duration
	Memory    uint64
}

// SearchBenchmark represents the result of a search algorithm benchmark.
// It contains an array of SearchResult, the size of the list being searched,
// and the name of the fastest search algorithm.
type SearchBenchmark struct {
	Results  []SearchResult
	ListSize int
	Fastest  string
}

// CompareSearchAlgorithms compares the performance of three different search algorithms:
// Binary Search, Linear Search, and Jump Search. It takes a list of integers as input
// and returns a SearchBenchmark struct that contains the results of each algorithm,
// the size of the list, and the name of the fastest algorithm.
//
// The function first selects a random target from the list. Then, it benchmarks the
// Binary Search algorithm, Linear Search algorithm, and Jump Search algorithm by
// calling the benchmarkSearch function for each algorithm. The results of each algorithm
// are appended to the Results field of the SearchBenchmark struct.
//
// After benchmarking is completed, the function determines the fastest algorithm
// by comparing the execution time of each algorithm in the benchmark results. The
// algorithm with the shortest execution time is set as the Fastest field of the
// SearchBenchmark struct.
//
// The function returns the SearchBenchmark struct with all the benchmark results.
//
// SearchBenchmark struct:
//
//	type SearchBenchmark struct {
//	    Results  []SearchResult
//	    ListSize int
//	    Fastest  string
//	}
//
// benchmarkSearch function:
//
//	func benchmarkSearch(name string, searchFunc func() int) SearchResult {
//	    ...
//	}
//
// BinarySearch function:
//
//	func BinarySearch(arr []int, target int) int {
//	    ...
//	}
//
// LinearSearch function:
//
//	func LinearSearch(slice []int, target int) int {
//	    ...
//	}
//
// JumpSearch function:
//
//	func JumpSearch(arr []int, x int) int {
//	    ...
//	}
func CompareSearchAlgorithms(list []int) SearchBenchmark {
	benchmark := SearchBenchmark{
		ListSize: len(list),
	}

	// Choose a random target from the list
	target := list[rand.Intn(len(list))]

	// Benchmark Binary Search
	benchmark.Results = append(benchmark.Results, benchmarkSearch("Binary Search", func() int {
		return searching.BinarySearch(list, target)
	}))

	// Benchmark Linear Search
	benchmark.Results = append(benchmark.Results, benchmarkSearch("Linear Search", func() int {
		return searching.LinearSearch(list, target)
	}))

	// Benchmark Jump Search
	benchmark.Results = append(benchmark.Results, benchmarkSearch("Jump Search", func() int {
		return searching.JumpSearch(list, target)
	}))

	// get the fastest search algorithm
	fastest := benchmark.Results[0]
	for _, result := range benchmark.Results {
		if result.Time < fastest.Time {
			fastest = result
		}
	}
	benchmark.Fastest = fastest.Algorithm

	return benchmark
}

// benchmarkSearch measures the performance of a search algorithm by timing its execution
// and measuring its memory usage.
//
// It takes a name string and a searchFunc function as input, where the searchFunc function
// performs the search algorithm. The searchFunc function should return the index of the target
// if found, or -1 if not found.
//
// benchmarkSearch records the memory usage before and after executing the search algorithm,
// as well as the time taken to execute the algorithm. It returns a SearchResult struct
// that contains the algorithm name, index, time, and memory usage.
//
// The SearchResult struct has the following fields:
// - Algorithm: The name of the search algorithm.
// - Target: The target value being searched.
// - Index: The index of the target if found, or -1 if not found.
// - Time: The duration of time taken to execute the search algorithm.
// - Memory: The difference in memory usage before and after executing the search algorithm.
func benchmarkSearch(name string, searchFunc func() int) SearchResult {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memBefore := m.Alloc

	start := time.Now()
	index := searchFunc()
	duration := time.Since(start)

	runtime.ReadMemStats(&m)
	memAfter := m.Alloc

	return SearchResult{
		Algorithm: name,
		Index:     index,
		Time:      duration,
		Memory:    memAfter - memBefore,
	}
}

// CompareSearchAlgorithmsGeneric compares the performance of three generic search algorithms:
// Binary Search Generic, Linear Search Generic, and Jump Search Generic. It takes a slice of
// any type, a target element, and a less function as parameters. The less function determines
// the ordering of elements in the slice. It returns a SearchBenchmark struct that contains
// information about the benchmark results, including the list size, the fastest search algorithm,
// and an array of SearchResult structs that contain information about each search algorithm's
// name, index, time, and memory usage.
func CompareSearchAlgorithmsGeneric(list []interface{}, target interface{}, less func(a, b interface{}) bool) SearchBenchmark {
	benchmark := SearchBenchmark{
		ListSize: len(list),
	}

	// Benchmark Binary Search Generic
	benchmark.Results = append(benchmark.Results, benchmarkSearchGeneric("Binary Search Generic", func() int {
		return searching.BinarySearchGeneric(list, target, less)
	}))

	// Benchmark Linear Search Generic
	benchmark.Results = append(benchmark.Results, benchmarkSearchGeneric("Linear Search Generic", func() int {
		return searching.LinearSearchGeneric(list, target)
	}))

	// Benchmark Jump Search Generic
	benchmark.Results = append(benchmark.Results, benchmarkSearchGeneric("Jump Search Generic", func() int {
		return searching.JumpSearchGeneric(list, target, less)
	}))

	// get the fastest search algorithm
	fastest := benchmark.Results[0]
	for _, result := range benchmark.Results {
		if result.Time < fastest.Time {
			fastest = result
		}
	}
	benchmark.Fastest = fastest.Algorithm

	return benchmark
}

// benchmarkSearchGeneric benchmarks the performance of a search algorithm by measuring
// the time taken and memory used to execute the search function.
//
// Parameters:
//   - name: The name of the search algorithm being benchmarked.
//   - searchFunc: The function that performs the search operation.
//
// Returns:
//
//	A SearchResult struct containing the benchmark results, including the algorithm name,
//	the index of the target if found, the time taken to execute the search function,
//	and the memory used before and after the search operation.
//
// SearchResult struct:
//   - Algorithm: The name of the search algorithm.
//   - Index: The index of the target if found, or -1 if not found.
//   - Time: The duration of the search operation.
//   - Memory: The difference in memory usage before and after the search operation.
//
// Usage Example:
//
//	benchmarkSearchGeneric("Binary Search Generic", func() int {
//	  return searching.BinarySearchGeneric(list, target, less)
//	})
func benchmarkSearchGeneric(name string, searchFunc func() int) SearchResult {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memBefore := m.Alloc

	start := time.Now()
	index := searchFunc()
	duration := time.Since(start)

	runtime.ReadMemStats(&m)
	memAfter := m.Alloc

	return SearchResult{
		Algorithm: name,
		Index:     index,
		Time:      duration,
		Memory:    memAfter - memBefore,
	}
}
