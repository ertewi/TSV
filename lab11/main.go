package main

import (
	"fmt"
	"sort"
)

// Представление ребра графа
type Edge struct {
	u, v   int
	weight int
}

// Найти корень элемента с применением сжатия пути
func find(parent []int, u int) int {
	if parent[u] != u {
		parent[u] = find(parent, parent[u])
	}
	return parent[u]
}

// Объединить два поддерева
func union(parent, rank []int, u, v int) {
	rootU := find(parent, u)
	rootV := find(parent, v)

	if rootU != rootV {
		if rank[rootU] > rank[rootV] {
			parent[rootV] = rootU
		} else if rank[rootU] < rank[rootV] {
			parent[rootU] = rootV
		} else {
			parent[rootV] = rootU
			rank[rootU]++
		}
	}
}

// Реализация алгоритма Краскала
func kruskal(n int, edges []Edge) []Edge {
	// Сортируем ребра по весу
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	parent := make([]int, n)
	rank := make([]int, n)

	// Инициализация структуры данных
	for i := range parent {
		parent[i] = i
	}

	var mst []Edge
	for _, edge := range edges {
		if find(parent, edge.u) != find(parent, edge.v) {
			union(parent, rank, edge.u, edge.v)
			mst = append(mst, edge)
		}
	}

	return mst
}

func main() {
	edges := []Edge{
		{0, 1, 1},
		{0, 2, 2},
		{1, 2, 1},
	}
	n := 3
	mst := kruskal(n, edges)
	fmt.Println("MST:")
	for _, edge := range mst {
		fmt.Printf("(%d, %d, %d)\n", edge.u, edge.v, edge.weight)
	}

	edges2 := []Edge{
		{0, 1, 4},
		{0, 2, 1},
		{1, 2, 2},
		{1, 3, 5},
		{2, 3, 3},
	}
	n2 := 4
	mst2 := kruskal(n2, edges2)
	fmt.Println("\nMST 2:")
	for _, edge := range mst2 {
		fmt.Printf("(%d, %d, %d)\n", edge.u, edge.v, edge.weight)
	}
}

// --- read from stdin ---
// reader := bufio.NewReader(os.Stdin)
// fmt.Fscan(reader, &a, &b)
// fmt.Println(a + b)

// --- read from file ---
// file, _ := os.Open("input.txt")
// defer file.Close()
// reader := bufio.NewReader(file)
// fmt.Fscan(reader, &a, &b)
// fmt.Println(a + b)

// --- write in file ---
// file2, _ := os.Create("output.txt")
// s := fmt.Sprintf("%f", inputdata)
// file2.WriteString(s)

// --- split string on symbol ---
// strings.Split(string, symbol)

// --- string to int64 ---
// a, _ = strconv.ParseInt(string, 10, 64)

// --- string to float64 ---
// a, _ = strconv.ParseFloat(string, 64)
