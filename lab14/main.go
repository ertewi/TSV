package main

import (
	"fmt"
)

// Rukzak решает задачу о рюкзаке с использованием динамического программирования
func Rukzak(M int, price map[string][]int) int {
	// n - количество предметов
	n := len(price["m"])

	// dp - массив для хранения максимальных стоимостей для каждой массы от 0 до M
	dp := make([]int, M+1)

	// Проходим по всем предметам
	for i := 0; i < n; i++ {
		// Обновляем массив dp в обратном порядке, чтобы избежать переиспользования обновленных значений
		for j := M; j >= price["m"][i]; j-- {
			dp[j] = max(dp[j], dp[j-price["m"][i]]+price["c"][i])
		}
	}

	// Возвращаем максимальную стоимость при данной грузоподъемности M
	return dp[M]
}

// Функция для нахождения максимального числа из двух
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Пример данных
	price := map[string][]int{
		"m": {3, 5, 8},
		"c": {8, 14, 23},
	}

	M := 15
	fmt.Printf("Максимальная стоимость набора товаров при грузоподъемности %d: %d\n", M, Rukzak(M, price))
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
