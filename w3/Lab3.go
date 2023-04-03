// 64050229
// Vish Siriwatana
// Lab3
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numbers := rand.Perm(100)
	// numbers := []int{22, 2, 8, 60, 22, 29, 54, 63, 7, 10}
	fmt.Println("Unsorted Array :", numbers)
	fmt.Println("Merge Sort Algorithm :", mergeSort(numbers))
	// fmt.Println("Quick Sort Algorithm :", quickSort(numbers))
}

// Merge Sort Algorithm
func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	// fmt.Println("Splitting : ", numbers)
	mid := len(numbers) / 2
	left := mergeSort(numbers[:mid])
	// fmt.Println("Left : ", left)
	right := mergeSort(numbers[mid:])
	// fmt.Println("Right : ", right)
	return merge(left, right)
}

// Merge Function
func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// Quick Sort Algorithm
func quickSort(numbers []int) []int {
	fmt.Println("Splitting : ", numbers)
	if len(numbers) <= 1 {
		return numbers
	}
	pivot := numbers[0]
	var left []int
	var right []int

	for _, number := range numbers[1:] {
		fmt.Println("Number : ", number)
		if number < pivot {
			left = append(left, number)
			fmt.Println("Left : ", left)
		} else {
			right = append(right, number)
			fmt.Println("Right : ", right)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}
