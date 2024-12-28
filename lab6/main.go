package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Прямое дискретное преобразование Фурье (O(n^2))
func DFT(x []complex128) []complex128 {
	n := len(x)
	X := make([]complex128, n)
	for k := 0; k < n; k++ {
		for t := 0; t < n; t++ {
			angle := -2 * math.Pi * float64(k) * float64(t) / float64(n)
			X[k] += x[t] * cmplx.Exp(complex(0, angle))
		}
	}
	return X
}

// Обратное дискретное преобразование Фурье (O(n^2))
func IDFT(X []complex128) []complex128 {
	n := len(X)
	x := make([]complex128, n)
	for t := 0; t < n; t++ {
		for k := 0; k < n; k++ {
			angle := 2 * math.Pi * float64(k) * float64(t) / float64(n)
			x[t] += X[k] * cmplx.Exp(complex(0, angle))
		}
	}
	for i := range x {
		x[i] /= complex(float64(n), 0) // Деление на n
	}
	return x
}

// Дополнение массива нулями до длины n
func padArray(arr []complex128, n int) []complex128 {
	result := make([]complex128, n)
	copy(result, arr)
	return result
}

// Свертка через ДПФ
func convolutionDFT(signal, kernel []complex128) []complex128 {
	n := len(signal) + len(kernel) - 1

	// Дополнение массивов нулями
	signalPadded := padArray(signal, n)
	kernelPadded := padArray(kernel, n)

	// Прямое ДПФ
	signalDFT := DFT(signalPadded)
	kernelDFT := DFT(kernelPadded)

	// Перемножение в частотной области
	resultDFT := make([]complex128, n)
	for i := 0; i < n; i++ {
		resultDFT[i] = signalDFT[i] * kernelDFT[i]
	}

	// Обратное ДПФ
	result := IDFT(resultDFT)

	return result
}

func main() {
	// Пример входных данных
	signal := []complex128{1, 2, 3, 4, 5}
	kernel := []complex128{1, 0, -1}

	// Выполнение свертки
	result := convolutionDFT(signal, kernel)

	// Вывод результата
	fmt.Println("Сигнал:", signal)
	fmt.Println("Ядро:", kernel)
	fmt.Println("Результат свертки через ДПФ:", result)

	// Реальная часть результата
	realResult := make([]float64, len(result))
	for i, val := range result {
		realResult[i] = real(val)
	}
	fmt.Println("Результат свертки (реальная часть):", realResult)
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
