package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Определение структуры для приоритетной очереди
type Item struct {
	node     int
	distance int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Алгоритм Дейкстры
func dijkstra(graph map[int]map[int]int, start int) map[int]int {
	// Инициализация расстояний
	distances := make(map[int]int)
	for node := range graph {
		distances[node] = math.MaxInt64
	}
	distances[start] = 0

	// Инициализация приоритетной очереди
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, distance: 0})

	// Основной цикл
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		currentNode := current.node
		currentDistance := current.distance

		// Если расстояние уже оптимальное, пропускаем
		if currentDistance > distances[currentNode] {
			continue
		}

		// Обновляем соседей
		for neighbor, cost := range graph[currentNode] {
			newDistance := currentDistance + cost
			if newDistance < distances[neighbor] {
				distances[neighbor] = newDistance
				heap.Push(pq, &Item{node: neighbor, distance: newDistance})
			}
		}
	}

	return distances
}

func main() {
	// Пример графа
	graph := map[int]map[int]int{
		0: {1: 4, 2: 1},
		1: {3: 1},
		2: {1: 2, 3: 5},
		3: {},
	}

	start := 0
	distances := dijkstra(graph, start)

	// Вывод результатов
	fmt.Println("Кратчайшие расстояния от узла", start)
	for node, distance := range distances {
		fmt.Printf("До узла %d: %d\n", node, distance)
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
