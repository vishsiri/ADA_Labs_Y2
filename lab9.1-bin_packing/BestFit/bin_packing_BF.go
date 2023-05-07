package main

import (
	"fmt"
)

func BFBinPacking(n int, BinS int, s []int, Bin []int) int {
	bestB := 0
	m := 1
	used := make([]int, 1000)
	remainB := make([]int, 1000)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if used[j]+s[i] <= BinS {
				remainB[j] = BinS - used[j] - s[i]
				if j > m {
					m = j
				}
			} else {
				remainB[j] = -1
				if j == m {
					m += 1
				}
			}
		}
		min := 9999
		for j := 1; j <= m; j++ {
			r := remainB[j]
			if r >= 0 && min > r {
				min = r
				bestB = j
			}
		}
		b := bestB
		Bin[i] = b
		used[b] += s[i]
	}
	return m
}

func printObjectAndBin(n int, s []int, Bin []int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("s%d = %d in Bin%d\n", i, s[i], Bin[i])
	}
}

func printBinj(n int, B int, s []int, Bin []int) {
	fmt.Printf("\nBin%d: ", B)
	for i := 1; i <= n; i++ {
		if Bin[i] == B {
			fmt.Printf("%d(%d), ", i, s[i])
		}
	}
}

func PackAndPrint(n int, BinSize int, s []int, Bin []int) {
	fmt.Println()
	for i := 1; i <= n; i++ {
		fmt.Printf("%d, ", s[i])
	}
	fmt.Println()
	m := BFBinPacking(n, BinSize, s, Bin)
	printObjectAndBin(n, s, Bin)
	for j := 1; j <= m; j++ {
		printBinj(n, j, s, Bin)
	}
}

func sortObjects(n int, x []int) {
	for loop := 1; loop < n; loop++ {
		for i := 1; i <= n-loop; i++ {
			if x[i] < x[i+1] {
				temp := x[i]
				x[i] = x[i+1]
				x[i+1] = temp
			}
		}
	}
}

func main() {
	n := 9
	s := []int{0, 2, 2, 7, 8, 3, 6, 3, 2, 6}
	BinSize := 10
	Bin := make([]int, 1000)
	PackAndPrint(n, BinSize, s, Bin)
	sortObjects(n, s)
	PackAndPrint(n, BinSize, s, Bin)
}
