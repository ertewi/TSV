package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// DFT выполняет дискретное преобразование Фурье
func DFT(input []complex128) []complex128 {
	N := len(input)                 // Длина входного сигнала
	output := make([]complex128, N) // Инициализация результата

	for k := 0; k < N; k++ { // Перебор частотных компонент
		var sum complex128
		for n := 0; n < N; n++ { // Перебор временных отсчётов
			angle := -2 * math.Pi * float64(k*n) / float64(N)
			sum += input[n] * cmplx.Exp(complex(0, angle))
		}
		output[k] = sum
	}

	return output
}

func main() {
	// Пример входного сигнала
	input := []complex128{
		1 + 0i, 1 + 0i, 1 + 0i, 0 + 0i,
		0 + 0i, 0 + 0i, 0 + 0i, 0 + 0i,
	}

	// Выполнение ДПФ
	output := DFT(input)

	// Вывод результата
	fmt.Println("Результат ДПФ:")
	for k, value := range output {
		fmt.Printf("X[%d] = %.3f + %.3fi\n", k, real(value), imag(value))
	}
}
