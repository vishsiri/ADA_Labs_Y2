##THE GRAPH TRAVERSAL EXPLAIN


<hr>

ใน function "graphDFS"
```go
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

```

อธิบาย

parameter ที่จะรับค่าเข้ามาคำนวณ 
```go
(i int, n int, L []byte, v [][]byte, visit []int)
```

```go
"i int" คือ 
"n int" คือ
"L [byte]" คือ
"v [][]byte" คือข้อมูลของ Vertex แต่ละจุดจะสามารถแสดงให้เห็นว่าแต่ละจุดมีจุดไหนบ้างที่จะสามารถมีจุดต่อกันได้

"visit []int" คือ
```


<hr>