// 代码生成时间: 2025-07-31 14:45:45
package main

import (
    "fmt"
    "sort"
    "revel"
)

// SortingService contains methods to perform sorting operations.
type SortingService struct{}

// BubbleSort performs a bubble sort on the given slice of integers.
func (service SortingService) BubbleSort(numbers []int) ([]int, error) {
    n := len(numbers)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if numbers[j] > numbers[j+1] {
                // Swap the elements
                numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
            }
        }
    }
    return numbers, nil
}

// InsertionSort performs an insertion sort on the given slice of integers.
func (service SortingService) InsertionSort(numbers []int) ([]int, error) {
    for i := 1; i < len(numbers); i++ {
        key := numbers[i]
        j := i - 1
        // Move elements of numbers[0..i-1], that are greater than key, to one position ahead of their current position
        for j >= 0 && numbers[j] > key {
            numbers[j+1] = numbers[j]
            j--
        }
        numbers[j+1] = key
    }
    return numbers, nil
}

// SelectionSort performs a selection sort on the given slice of integers.
func (service SortingService) SelectionSort(numbers []int) ([]int, error) {
    n := len(numbers)
    for i := 0; i < n-1; i++ {
        // Find the minimum element in the remaining unsorted array
        minIdx := i
        for j := i + 1; j < n; j++ {
            if numbers[j] < numbers[minIdx] {
                minIdx = j
            }
        }
        // Swap the found minimum element with the first element of the unsorted part
        numbers[i], numbers[minIdx] = numbers[minIdx], numbers[i]
    }
    return numbers, nil
}

// AppController handles sorting operations.
type AppController struct {
    *revel.Controller
}

// SortAction handles sorting request and returns the sorted array.
func (c AppController) SortAction(method string) revel.Result {
    var numbers []int
    var err error

    // Example input: [5, 3, 8, 4, 2]
    numbers = []int{5, 3, 8, 4, 2}

    // Choose the sorting method based on the request parameter
    switch method {
    case "bubble":
        numbers, err = SortingService{}.BubbleSort(numbers)
    case "insertion":
        numbers, err = SortingService{}.InsertionSort(numbers)
    case "selection":
        numbers, err = SortingService{}.SelectionSort(numbers)
    default:
        err = fmt.Errorf("invalid sorting method")
    }

    if err != nil {
        return c.RenderError(err)
    }

    return c.RenderJSON(map[string]interface{}{"sortedNumbers": numbers})
}

func init() {
    // Filter is a type of middleware that preprocesses requests.
    revel.Filters.Append(revel.PanicFilter, revel.RecoverFilter)
}

func main() {
    revel.Run()
}