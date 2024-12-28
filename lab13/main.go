package main

import (
	"fmt"
	"math"
)

// Graph представляет граф с количеством вершин и списком рёбер
type Graph struct {
	numVertices int
	edges       []Edge
}

// Edge представляет ребро в графе с начальной вершиной, конечной вершиной и весом
type Edge struct {
	start, end, weight int
}

// NewGraph создаёт новый граф с указанным количеством вершин
func NewGraph(numVertices int) *Graph {
	return &Graph{
		numVertices: numVertices,
		edges:       []Edge{},
	}
}

// AddEdge добавляет ребро в граф
func (g *Graph) AddEdge(start, end, weight int) {
	g.edges = append(g.edges, Edge{start, end, weight})
}

// BellmanFord выполняет алгоритм Беллмана-Форда для поиска кратчайших путей
func (g *Graph) BellmanFord(source int) ([]int, error) {
	// Инициализация расстояний: все бесконечны, кроме исходной вершины
	distances := make([]int, g.numVertices)
	for i := range distances {
		distances[i] = math.MaxInt
	}
	distances[source] = 0

	// Расслабляем все рёбра (V-1) раз
	for i := 0; i < g.numVertices-1; i++ {
		for _, edge := range g.edges {
			if distances[edge.start] != math.MaxInt && distances[edge.start]+edge.weight < distances[edge.end] {
				distances[edge.end] = distances[edge.start] + edge.weight
			}
		}
	}

	// Проверка на наличие циклов с отрицательным весом
	for _, edge := range g.edges {
		if distances[edge.start] != math.MaxInt && distances[edge.start]+edge.weight < distances[edge.end] {
			return nil, fmt.Errorf("граф содержит цикл с отрицательным весом")
		}
	}

	// Возвращаем итоговые кратчайшие расстояния
	return distances, nil
}

// PrintDistances выводит кратчайшие расстояния от исходной вершины
func (g *Graph) PrintDistances(source int, distances []int) {
	fmt.Printf("Кратчайшие расстояния от вершины %d:\n", source)
	for vertex, distance := range distances {
		if distance == math.MaxInt {
			fmt.Printf("Вершина %d: недостижима\n", vertex)
		} else {
			fmt.Printf("Вершина %d: %d\n", vertex, distance)
		}
	}
}

func main() {
	// Создаём граф с 4 вершинами
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 3, 4)
	g.AddEdge(1, 2, 2)
	g.AddEdge(3, 2, 3)
	g.AddEdge(2, 3, -1)

	// Запуск алгоритма Беллмана-Форда из вершины 0
	result, err := g.BellmanFord(0)
	if err != nil {
		fmt.Println(err) // Если был найден цикл с отрицательным весом
	} else {
		g.PrintDistances(0, result) // Если успешный расчёт, выводим расстояния
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
