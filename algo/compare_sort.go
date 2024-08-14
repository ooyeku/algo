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

type SortResult struct {
	Algorithm string
	Time      time.Duration
	Memory    uint64
}

type SortBenchmark struct {
	Results  []SortResult
	ListSize int
	Fastest  string
}

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
