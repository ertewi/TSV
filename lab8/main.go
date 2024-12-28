package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Функция для вычисления FFT
func fft(x []complex128) []complex128 {
	N := len(x)
	if N <= 1 {
		return x
	}

	// Разделение на четные и нечетные элементы
	even := fft(evenElements(x))
	odd := fft(oddElements(x))

	// Вычисление T (множители для НШП)
	T := make([]complex128, N/2)
	for k := 0; k < N/2; k++ {
		T[k] = cmplx.Exp(complex(0, -2*math.Pi*float64(k)/float64(N))) * odd[k]
	}

	// Слияние четных и нечетных частей
	result := make([]complex128, N)
	for k := 0; k < N/2; k++ {
		result[k] = even[k] + T[k]
		result[k+N/2] = even[k] - T[k]
	}

	return result
}

// Функция для вычисления IFFT
func ifft(x []complex128) []complex128 {
	N := len(x)
	if N <= 1 {
		return x
	}

	// Разделение на четные и нечетные элементы
	even := ifft(evenElements(x))
	odd := ifft(oddElements(x))

	// Вычисление T (множители для НШП)
	T := make([]complex128, N/2)
	for k := 0; k < N/2; k++ {
		T[k] = cmplx.Exp(complex(0, 2*math.Pi*float64(k)/float64(N))) * odd[k]
	}

	// Слияние четных и нечетных частей
	result := make([]complex128, N)
	for k := 0; k < N/2; k++ {
		result[k] = (even[k] + T[k]) / 2
		result[k+N/2] = (even[k] - T[k]) / 2
	}

	return result
}

// Функция для извлечения четных элементов из среза
func evenElements(x []complex128) []complex128 {
	even := make([]complex128, len(x)/2)
	for i := 0; i < len(x)/2; i++ {
		even[i] = x[2*i]
	}
	return even
}

// Функция для извлечения нечетных элементов из среза
func oddElements(x []complex128) []complex128 {
	odd := make([]complex128, len(x)/2)
	for i := 0; i < len(x)/2; i++ {
		odd[i] = x[2*i+1]
	}
	return odd
}

// Функция для свертки двух последовательностей
func convolve(x, y []complex128) []complex128 {
	// Длина свертки
	N := len(x) + len(y) - 1

	// Паддинг до следующей степени двойки для эффективности
	N = 1 << (int(math.Log2(float64(N))) + 1)

	// Паддинг входных данных
	X := make([]complex128, N)
	Y := make([]complex128, N)

	copy(X, x)
	copy(Y, y)

	// Применение FFT к данным
	X_fft := fft(X)
	Y_fft := fft(Y)

	// Вычисление свертки в частотной области
	Z_fft := make([]complex128, N)
	for i := 0; i < N; i++ {
		Z_fft[i] = X_fft[i] * Y_fft[i]
	}

	// Применение IFFT для получения результата свертки
	Z := ifft(Z_fft)

	// Обрезка лишних элементов и возврат результата
	result := make([]complex128, len(x)+len(y)-1)
	copy(result, Z[:len(result)])

	return result
}

func main() {
	// Пример использования
	x := []complex128{1, 2, 3}
	y := []complex128{4, 5, 6, 7, 8}

	result := convolve(x, y)

	// Печать результата
	fmt.Println("Результат свертки:")
	for _, val := range result {
		fmt.Printf("%v\n", val)
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
