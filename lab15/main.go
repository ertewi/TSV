package main

import (
	"fmt"
	"math"
)

// Функция для вычисления минимального числа операций для умножения матриц
func calculateMinOperations(dimensions []int) ([][]int, [][]int) {
	numMatrices := len(dimensions) - 1 // количество матриц
	minOperations := make([][]int, numMatrices)
	splitPoints := make([][]int, numMatrices)

	// Инициализация таблиц
	for i := range minOperations {
		minOperations[i] = make([]int, numMatrices)
		splitPoints[i] = make([]int, numMatrices)
	}

	// l - длина цепочки матриц, начиная с 2
	for chainLength := 2; chainLength <= numMatrices; chainLength++ {
		for start := 0; start <= numMatrices-chainLength; start++ {
			end := start + chainLength - 1
			minOperations[start][end] = math.MaxInt32 // инициализируем большим значением

			// k - точка разбиения цепочки матриц на две подзадачи
			for split := start; split < end; split++ {
				operations := minOperations[start][split] +
					minOperations[split+1][end] +
					dimensions[start]*dimensions[split+1]*dimensions[end+1]

				// Если текущее количество операций меньше предыдущего, обновляем минимальное значение
				if operations < minOperations[start][end] {
					minOperations[start][end] = operations
					splitPoints[start][end] = split
				}
			}
		}
	}

	return minOperations, splitPoints
}

// Функция для восстановления оптимального порядка умножения
func getOptimalParenthesization(splitPoints [][]int, start, end int) string {
	if start == end {
		return fmt.Sprintf("A%d", start+1) // Одиночная матрица
	} else {
		left := getOptimalParenthesization(splitPoints, start, splitPoints[start][end])
		right := getOptimalParenthesization(splitPoints, splitPoints[start][end]+1, end)
		return fmt.Sprintf("(%s x %s)", left, right) // Возвращаем строку с расстановкой скобок
	}
}

func main() {
	// Пример использования
	dimensions := []int{10, 20, 50, 1, 100} // Размерности матриц
	minOperations, splitPoints := calculateMinOperations(dimensions)

	optimalOrder := getOptimalParenthesization(splitPoints, 0, len(dimensions)-2)
	minOperationsCount := minOperations[0][len(dimensions)-2]

	fmt.Printf("Минимальное количество операций: %d\n", minOperationsCount)
	fmt.Printf("Оптимальный порядок умножения: %s\n", optimalOrder)
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
