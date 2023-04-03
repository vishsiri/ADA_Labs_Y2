// Vish Siriwatana
// 64050229

package main

import "fmt"

func graphDFS(i int, n int, L []byte, v [][]byte, visit []int) {
	visit[i] = 1
	fmt.Printf("%c ", L[i])
	for j := 0; j < n; j++ {
		w := v[i][j]
		if w != 0 {
			I := int(w - 'A')
			if visit[I] == 0 {
				graphDFS(I, n, L, v, visit)
			}
		} else {
			break
		}
	}
}

func graphBFS(n int, L []byte, v [][]byte) {
	visit := make([]int, n)
	queue := make([]byte, n)
	front := -1
	rear := -1
	queue[0] = L[0]
	rear++
	visit[0] = 1
	for front != rear {
		node := queue[front+1]
		front++
		fmt.Printf("%c ", node)
		I := -1
		for j := 0; j < n; j++ {
			if L[j] == node {
				I = j
				break
			}
		}
		for k := 0; k < n; k++ {
			if v[I][k] != 0 {
				for i := 0; i < n; i++ {
					if L[i] == v[I][k] && visit[i] == 0 {
						queue[rear+1] = v[I][k]
						rear++
						visit[i] = 1
					}
				}
			} else {
				break
			}
		}
	}
}

func printAdjLists(n int, L []byte, v [][]byte) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d %c: ", i, L[i])
		for j := 0; j < n; j++ {
			if v[i][j] != 0 {
				fmt.Printf("%c ", v[i][j])
			} else {
				break
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	n := 8
	visit := make([]int, 100)
	L := []byte("ABCDEFGH")
	v := [][]byte{{'G', 'F', 'D', 'B', 0, 0, 0, 0},
		{'A', 'H', 'C', 0, 0, 0, 0, 0},
		{'B', 0, 0, 0, 0, 0, 0, 0},
		{'A', 'E', 0, 0, 0, 0, 0, 0},
		{'F', 'D', 0, 0, 0, 0, 0, 0},
		{'A', 'E', 0, 0, 0, 0, 0, 0},
		{'A', 0, 0, 0, 0, 0, 0, 0},
		{'B', 0, 0, 0, 0, 0, 0, 0}}
	fmt.Println("Adjacent Lists (n=", n, ")")
	printAdjLists(n, L, v)
	fmt.Println("\nDFS: ")
	graphDFS(0, n, L, v, visit)
	fmt.Println("\nBFS: ")
	graphBFS(n, L, v)
}

//Explain
//v
