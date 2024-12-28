package main

import (
	"fmt"
	"math"
	"strings"
)

// Функция Карацубы для умножения двух чисел
func karatsuba(num1, num2 int) int {
	// Базовый случай: если одно из чисел однозначное, возвращаем их произведение
	if num1 < 10 || num2 < 10 {
		return num1 * num2
	}

	// Преобразуем числа в строки для вычисления длины
	len1 := int(math.Log10(float64(num1))) + 1
	len2 := int(math.Log10(float64(num2))) + 1

	// Определяем максимальную длину
	length := int(math.Max(float64(len1), float64(len2)))

	// Длина для деления числа пополам
	halfLength := length / 2

	// Разделяем числа на старшие и младшие разряды
	high1 := num1 / int(math.Pow(10, float64(halfLength)))
	low1 := num1 % int(math.Pow(10, float64(halfLength)))

	high2 := num2 / int(math.Pow(10, float64(halfLength)))
	low2 := num2 % int(math.Pow(10, float64(halfLength)))

	// Рекурсивные вызовы Карацубы
	zLow := karatsuba(low1, low2)
	zHigh := karatsuba(high1, high2)
	zCross := karatsuba(low1+high1, low2+high2)

	// Комбинируем результат по формуле Карацубы
	return zHigh*int(math.Pow(10, float64(2*halfLength))) +
		(zCross-zHigh-zLow)*int(math.Pow(10, float64(halfLength))) +
		zLow
}

func main() {
	x := 15
	y := 15
	result := karatsuba(x, y)

	// Вывод результата
	fmt.Printf("  %d\n", x)
	fmt.Printf("x %d\n", y)
	fmt.Println(strings.Repeat("-", int(math.Max(float64(len(fmt.Sprint(x))), float64(len(fmt.Sprint(y)))))+len(fmt.Sprint(result))))
	fmt.Printf("  %d\n", result)
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
