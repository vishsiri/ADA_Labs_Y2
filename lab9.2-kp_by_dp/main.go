package main

import "fmt"

func printX(n int, x []int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", x[i])
	}
	fmt.Println()
}

func printZ(C int, z []int) {
	for i := 0; i <= C; i++ {
		fmt.Printf("%d ", z[i])
	}
	fmt.Println()
}

func main() {
	n := 7
	C := 9
	p := []int{6, 5, 8, 9, 6, 7, 3}
	w := []int{2, 3, 6, 7, 5, 9, 4}
	z := make([]int, C+1)
	fmt.Printf("n = %d C = %d\n", n, C)
	fmt.Printf("P: ")
	printX(n, p)
	fmt.Printf("W: ")
	printX(n, w)
	for d := 0; d <= C; d++ {
		fmt.Printf("%d ", d)
	}
	fmt.Println()
	printZ(C, z)
	for j := 0; j < n; j++ {
		for d := w[j]; d <= C; d++ {
			if z[d-w[j]]+p[j] > z[d] {
				z[d] = z[d-w[j]] + p[j]
			}
		}
		fmt.Printf("%d ", j)
		printZ(C, z)
	}
}
