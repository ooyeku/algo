
# Algo

This repo contains implementations of various simple algorithms and data structures in Go.


## 🌟 Features

- 🔍 Searching algorithms: Binary, Linear, Jump
- 🔢 Sorting algorithms: Bubble, Merge, Quick, Heap, Intro
- 🌳 Data structures: Binary Search Tree
- 🏎️ Performance benchmarking
- 🧠 Generic implementations for maximum flexibility

## 🚀 Quick Start

To install the package, run:

```bash
go get github.com/ooyeku/algo
```

## Usage

```go
package main

import (
    "fmt"
	"github.com/ooyeku/algo"
	"github.com/ooyeku/algo/searching"
	"github.com/ooyeku/algo/sorting"
)

func main() {
    // Generate a list of 100 random integers
    list := algo.GenerateList(100, 1, 1000000)
    // Sort the list using QuickSort
    sortedList := sorting.QuickSort(list)
    // Search for a value using BinarySearch
    target := sortedList[50]
    index := searching.BinarySearch(sortedList, target)
    fmt.Printf("Binary Search: Target %d found at index %d\n", target, index)
    // Benchmark sorting algorithms
    benchmark := algo.CompareSortAlgorithms(list)
    fmt.Printf("Fastest sorting algorithm: %s\n", benchmark.Fastest)
}
```


## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

