// 代码生成时间: 2025-09-21 02:26:12
package main

import (
    "fmt"
    "sort"
    "math/rand"
    "time"
)

// Sorter interface defines a method to sort a slice
type Sorter interface {
    Sort([]int) []int
}

// BubbleSort is a struct that implements the Sorter interface
type BubbleSort struct {}

// Sort implements the Sorter interface by sorting a slice using bubble sort
func (bs *BubbleSort) Sort(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
}

// QuickSort is a struct that implements the Sorter interface
type QuickSort struct {}

// Sort implements the Sorter interface by sorting a slice using quick sort
func (qs *QuickSort) Sort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    left, right := 0, len(arr)-1
    for left < right {
        for left < right && arr[right] >= arr[0] {
            right--
        }
        arr[left], arr[right] = arr[right], arr[left]
        for left < right && arr[left] <= arr[0] {
            left++
        }
        arr[left], arr[right] = arr[right], arr[left]
    }
    arr[0], arr[left] = arr[left], arr[0]
    return arr
}

// GenerateRandomSlice generates a random slice of integers
func GenerateRandomSlice() []int {
    rand.Seed(time.Now().UnixNano())
    size := rand.Intn(100) + 1
    numbers := make([]int, size)
    for i := 0; i < size; i++ {
        numbers[i] = rand.Intn(100)
    }
    return numbers
}

func main() {
    // Generate a random slice
    slice := GenerateRandomSlice()

    // Print the original slice
    fmt.Println("Original slice: ", slice)

    // Create instances of sorters
    bubbleSorter := &BubbleSort{}
    quickSorter := &QuickSort{}

    // Sort the slice using bubble sort
    sortedSlice := bubbleSorter.Sort(slice)
    fmt.Println("Sorted slice (Bubble Sort): ", sortedSlice)

    // Generate another random slice
    slice = GenerateRandomSlice()

    // Sort the slice using quick sort
    sortedSlice = quickSorter.Sort(slice)
    fmt.Println("Sorted slice (Quick Sort): ", sortedSlice)
}
