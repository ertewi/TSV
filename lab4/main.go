package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// FFT выполняет быстрое преобразование Фурье
func FFT(input []complex128) []complex128 {
	N := len(input)

	// Базовый случай: если длина массива равна 1, возвращаем его
	if N <= 1 {
		return input
	}

	// Проверка: N должно быть степенью двойки
	if N&(N-1) != 0 {
		panic("Длина входного массива должна быть степенью двойки")
	}

	// Разделяем сигнал на чётные и нечётные элементы
	even := make([]complex128, N/2)
	odd := make([]complex128, N/2)
	for i := 0; i < N/2; i++ {
		even[i] = input[2*i]
		odd[i] = input[2*i+1]
	}

	// Рекурсивные вызовы для чётных и нечётных частей
	fftEven := FFT(even)
	fftOdd := FFT(odd)

	// Объединяем результаты
	output := make([]complex128, N)
	for k := 0; k < N/2; k++ {
		// Вычисляем комплексный множитель (так называемый "бабочка")
		t := cmplx.Exp(complex(0, -2*math.Pi*float64(k)/float64(N))) * fftOdd[k]
		output[k] = fftEven[k] + t
		output[k+N/2] = fftEven[k] - t
	}

	return output
}

func main() {
	// Пример входного сигнала
	input := []complex128{
		1 + 0i, 1 + 0i, 1 + 0i, 0 + 0i,
		0 + 0i, 0 + 0i, 0 + 0i, 0 + 0i,
	}

	// Выполнение БПФ
	output := FFT(input)

	// Вывод результата
	fmt.Println("Результат БПФ:")
	for k, value := range output {
		fmt.Printf("X[%d] = %.3f + %.3fi\n", k, real(value), imag(value))
	}
}
