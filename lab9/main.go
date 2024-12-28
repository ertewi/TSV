package main

import (
	"fmt"
	"strings"
)

func multiplyStrings(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	len1 := len(num1)
	len2 := len(num2)
	result := make([]int, len1+len2)

	// Вспомогательная функция для умножения цифр
	multiplyDigits := func(i, j int) int {
		return int(num1[i]-'0') * int(num2[j]-'0')
	}

	// Вспомогательная функция для добавления промежуточных результатов
	addToResult := func(pos1, pos2, product int) {
		total := product + result[pos2]
		result[pos2] = total % 10
		result[pos1] += total / 10
	}

	// Основной цикл умножения
	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			mul := multiplyDigits(i, j)
			addToResult(i+j, i+j+1, mul)
		}
	}

	// Преобразование массива результата в строку
	var sb strings.Builder
	leadingZero := true
	for _, v := range result {
		if v != 0 || !leadingZero {
			sb.WriteByte(byte(v) + '0')
			leadingZero = false
		}
	}

	if sb.Len() == 0 {
		return "0"
	}

	return sb.String()
}

func main() {
	num1 := "123"
	num2 := "456"
	result := multiplyStrings(num1, num2)

	fmt.Printf(" %s\n", num1)
	fmt.Printf("x %s\n", num2)
	fmt.Println("--")
	fmt.Println(result)
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
