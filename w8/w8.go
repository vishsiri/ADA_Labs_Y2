package main

import "fmt"

func Pow2n(n int) int {
	return 1 << n
}

func solutionX(X int, b []int, w []int) int {
	var sum int
	for i := 0; i < 4; i++ {
		b[i] = X % 2
		X = X / 2
		sum += b[i] * w[i]
	}
	return sum
}

func main() {
	var N, ix, sum int
	var x [100]int
	n, M, w := 4, 31, [4]int{11, 13, 24, 7}
	N = Pow2n(n)
	for ix = 1; ix < N; ix++ {
		for i := 0; i < n; i++ {
			x[i] = 0
		}
		sum = solutionX(ix, x[:], w[:])
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", x[i])
		}
		fmt.Printf("sum = %d ", sum)
		if sum == M {
			fmt.Printf("*\n")
		} else {
			fmt.Printf("\n")
		}
	}
}
