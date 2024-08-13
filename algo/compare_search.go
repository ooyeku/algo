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

type SearchResult struct {
	Algorithm string
	Target    interface{}
	Index     int
	Time      time.Duration
	Memory    uint64
}

type SearchBenchmark struct {
	Results  []SearchResult
	ListSize int
	Fastest  string
}

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
