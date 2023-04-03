// 64050229
// Vish Siriwatana
// Lab5
package main

import (
	"fmt"
)

func main() {

	//INPUT DATA FROM USER
	var items []int
	var weights []int
	var capacity int
	var item int
	var weight int
	var n int
	fmt.Print("Enter number of items : ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Enter item value : ")
		fmt.Scanln(&item)
		items = append(items, item)
		fmt.Print("Enter item weight : ")
		fmt.Scanln(&weight)
		weights = append(weights, weight)
	}
	fmt.Print("Enter capacity : ")
	fmt.Scanln(&capacity)

	// items := []int{16, 41, 51, 22, 90}
	// weights := []int{8, 13, 12, 15, 20}
	// capacity := 45

	fmt.Println("Items :", items)
	fmt.Println("Weights :", weights)
	fmt.Println("Capacity :", capacity)

	fmt.Println("Max Value :", knapSack(items, weights, capacity))

}

func knapSack(items, weights []int, capacity int) int {
	// P/w ratio
	var ratio []float64
	for i := 0; i < len(items); i++ {
		ratio = append(ratio, float64(items[i])/float64(weights[i]))
	}
	fmt.Println("Ratio :", ratio)

	// Sort ratio
	for i := 0; i < len(ratio); i++ {
		for j := i + 1; j < len(ratio); j++ {
			if ratio[i] < ratio[j] {
				ratio[i], ratio[j] = ratio[j], ratio[i]
				items[i], items[j] = items[j], items[i]
				weights[i], weights[j] = weights[j], weights[i]
			}
		}
	}
	fmt.Println("Sorted Ratio :", ratio)
	fmt.Println("Sorted Items :", items)
	fmt.Println("Sorted Weights :", weights)

	// Max Value
	var maxValue int
	for i := 0; i < len(items); i++ {
		// Skip if weight is 0
		if weights[i] == 0 {
			continue
		}
		if capacity >= weights[i] {
			capacity -= weights[i]
			maxValue += items[i]
		} else {
			maxValue += items[i] * capacity / weights[i]
			break
		}
	}
	return maxValue
}
