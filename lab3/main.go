package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Полубыстрое прямое преобразование Фурье
func halfFastFourierTransform(f []complex128) []complex128 {
	N := len(f)

	// Поиск оптимальных делителей для разбиения
	p1 := int(math.Sqrt(float64(N)))
	for N%p1 != 0 {
		p1--
	}
	p2 := N / p1

	// Проверка корректности деления
	if p1*p2 != N {
		panic("Ошибка: не удалось корректно разделить массив")
	}

	A := make([]complex128, N)

	// Предвычисление экспонент для ускорения
	exp1 := make([][]complex128, p1)
	for i := 0; i < p1; i++ {
		exp1[i] = make([]complex128, p1)
		for j := 0; j < p1; j++ {
			exp1[i][j] = cmplx.Exp(complex(0, -2*math.Pi*float64(i*j)/float64(p1)))
		}
	}

	exp2 := make([][]complex128, p2)
	for i := 0; i < p2; i++ {
		exp2[i] = make([]complex128, p2)
		for j := 0; j < p2; j++ {
			exp2[i][j] = cmplx.Exp(complex(0, -2*math.Pi*float64(i*j)/float64(p2)))
		}
	}

	// Применение двумерного FFT в разбиении на p1 и p2
	for k1 := 0; k1 < p1; k1++ {
		for k2 := 0; k2 < p2; k2++ {
			var summation complex128
			for j1 := 0; j1 < p1; j1++ {
				for j2 := 0; j2 < p2; j2++ {
					j := j1 + p1*j2
					exponent := exp1[k1][j1] * exp2[k2][j2]
					summation += f[j] * exponent
				}
			}
			A[k1+p1*k2] = summation
		}
	}

	return A
}

// Полубыстрое обратное преобразование Фурье
func inverseHalfFastFourierTransform(A []complex128) []complex128 {
	N := len(A)

	// Поиск оптимальных делителей для разбиения
	p1 := int(math.Sqrt(float64(N)))
	for N%p1 != 0 {
		p1--
	}
	p2 := N / p1

	// Проверка корректности деления
	if p1*p2 != N {
		panic("Ошибка: не удалось корректно разделить массив")
	}

	f := make([]complex128, N)

	// Предвычисление экспонент для обратного преобразования
	exp1 := make([][]complex128, p1)
	for i := 0; i < p1; i++ {
		exp1[i] = make([]complex128, p1)
		for j := 0; j < p1; j++ {
			exp1[i][j] = cmplx.Exp(complex(0, 2*math.Pi*float64(i*j)/float64(p1)))
		}
	}

	exp2 := make([][]complex128, p2)
	for i := 0; i < p2; i++ {
		exp2[i] = make([]complex128, p2)
		for j := 0; j < p2; j++ {
			exp2[i][j] = cmplx.Exp(complex(0, 2*math.Pi*float64(i*j)/float64(p2)))
		}
	}

	// Применение двумерного обратного FFT в разбиении на p1 и p2
	for j1 := 0; j1 < p1; j1++ {
		for j2 := 0; j2 < p2; j2++ {
			var summation complex128
			for k1 := 0; k1 < p1; k1++ {
				for k2 := 0; k2 < p2; k2++ {
					k := k1 + p1*k2
					exponent := exp1[j1][k1] * exp2[j2][k2]
					summation += A[k] * exponent
				}
			}
			f[j1+p1*j2] = summation / complex(float64(N), 0) // Нормировка
		}
	}

	return f
}

func main() {
	// Пример использования
	f := []complex128{1, 2, 3, 4}

	// Прямое полубыстрое преобразование Фурье
	result := halfFastFourierTransform(f)

	// Обратное полубыстрое преобразование Фурье
	inverseResult := inverseHalfFastFourierTransform(result)

	// Обычное преобразование Фурье для сравнения
	fftResult := make([]complex128, len(f))
	for i := range f {
		fftResult[i] = cmplx.Exp(complex(0, -2*math.Pi*float64(i)/float64(len(f))))
	}

	ifftResult := make([]complex128, len(fftResult))
	for i := range fftResult {
		ifftResult[i] = fftResult[i] / complex(float64(len(f)), 0)
	}

	// Печать результатов
	fmt.Println("Результат полубыстрого преобразования Фурье (ПШФ):")
	fmt.Println("Индекс | Значение")
	fmt.Println("-------------------")
	for i, val := range result {
		fmt.Printf("%6d | %.4f\n", i, val)
	}

	fmt.Println("\nРезультат обратного ПШФ:")
	fmt.Println("Индекс | Значение")
	fmt.Println("-------------------")
	for i, val := range inverseResult {
		fmt.Printf("%6d | %.4f\n", i, val)
	}

	fmt.Println("\nРезультат обычного FFT для сравнения:")
	fmt.Println("Индекс | Значение")
	fmt.Println("-------------------")
	for i, val := range fftResult {
		fmt.Printf("%6d | %.4f\n", i, val)
	}

	fmt.Println("\nРезультат обратного FFT для сравнения:")
	fmt.Println("Индекс | Значение")
	fmt.Println("-------------------")
	for i, val := range ifftResult {
		fmt.Printf("%6d | %.4f\n", i, val)
	}
}
