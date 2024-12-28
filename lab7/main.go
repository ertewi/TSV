package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Half Fast Fourier Transform
func halfFastFourierTransform(f []complex128) []complex128 {
	N := len(f)
	p1 := int(math.Sqrt(float64(N)))
	p2 := N / p1

	if p1*p2 != N {
		panic("N должно быть произведением p1 и p2")
	}

	// Первый шаг БПФ
	A1 := make([][]complex128, p1)
	for i := range A1 {
		A1[i] = make([]complex128, p2)
	}

	for k2 := 0; k2 < p2; k2++ {
		for k1 := 0; k1 < p1; k1++ {
			var sum complex128
			for j1 := 0; j1 < p1; j1++ {
				angle := -2 * math.Pi * float64(k1*j1) / float64(p1)
				sum += f[j1+p1*k2] * cmplx.Exp(complex(0, angle))
			}
			A1[k1][k2] = sum
		}
	}

	// Второй шаг БПФ
	A2 := make([][]complex128, p1)
	for i := range A2 {
		A2[i] = make([]complex128, p2)
	}

	for k1 := 0; k1 < p1; k1++ {
		for k2 := 0; k2 < p2; k2++ {
			var sum complex128
			for j2 := 0; j2 < p2; j2++ {
				angle := -2 * math.Pi * float64(k2*j2) / float64(p2)
				sum += A1[k1][j2] * cmplx.Exp(complex(0, angle))
			}
			A2[k1][k2] = sum
		}
	}

	// Объединение результатов
	A := make([]complex128, N)
	for k1 := 0; k1 < p1; k1++ {
		for k2 := 0; k2 < p2; k2++ {
			A[k1+p1*k2] = A2[k1][k2]
		}
	}

	return A
}

// Inverse Half Fast Fourier Transform
func inverseHalfFastFourierTransform(F []complex128) []complex128 {
	N := len(F)
	p1 := int(math.Sqrt(float64(N)))
	p2 := N / p1

	if p1*p2 != N {
		panic("N должно быть произведением p1 и p2")
	}

	FReshaped := make([][]complex128, p1)
	for i := range FReshaped {
		FReshaped[i] = make([]complex128, p2)
	}

	for k1 := 0; k1 < p1; k1++ {
		for k2 := 0; k2 < p2; k2++ {
			FReshaped[k1][k2] = F[k1+p1*k2]
		}
	}

	// Первый шаг обратного БПФ
	AInv1 := make([][]complex128, p1)
	for i := range AInv1 {
		AInv1[i] = make([]complex128, p2)
	}

	for j2 := 0; j2 < p2; j2++ {
		for j1 := 0; j1 < p1; j1++ {
			var sum complex128
			for k2 := 0; k2 < p2; k2++ {
				angle := 2 * math.Pi * float64(k2*j2) / float64(p2)
				sum += FReshaped[j1][k2] * cmplx.Exp(complex(0, angle))
			}
			AInv1[j1][j2] = sum
		}
	}

	// Второй шаг обратного БПФ
	AInv2 := make([][]complex128, p1)
	for i := range AInv2 {
		AInv2[i] = make([]complex128, p2)
	}

	for j1 := 0; j1 < p1; j1++ {
		for j2 := 0; j2 < p2; j2++ {
			var sum complex128
			for k1 := 0; k1 < p1; k1++ {
				angle := 2 * math.Pi * float64(k1*j1) / float64(p1)
				sum += AInv1[k1][j2] * cmplx.Exp(complex(0, angle))
			}
			AInv2[j1][j2] = sum
		}
	}

	// Объединение результатов и нормализация
	AInv := make([]complex128, N)
	for j1 := 0; j1 < p1; j1++ {
		for j2 := 0; j2 < p2; j2++ {
			AInv[j1+p1*j2] = AInv2[j1][j2] / complex(float64(N), 0)
		}
	}

	return AInv
}

// Свертка через БПФ
func halfFastFourierConvolution(f, g []complex128) []complex128 {
	N := len(f) + len(g) - 1

	// Дополнение массивов
	fPadded := make([]complex128, N)
	gPadded := make([]complex128, N)
	copy(fPadded, f)
	copy(gPadded, g)

	F := halfFastFourierTransform(fPadded)
	G := halfFastFourierTransform(gPadded)

	FG := make([]complex128, N)
	for i := 0; i < N; i++ {
		FG[i] = F[i] * G[i]
	}

	return inverseHalfFastFourierTransform(FG)
}

func main() {
	// Пример данных
	f := []complex128{1, 2, 3, 4}
	g := []complex128{1, 2, 3}

	// Вызов функции свертки
	result := halfFastFourierConvolution(f, g)

	// Красивый вывод результата
	fmt.Println("Результат свёртки:")
	for i, val := range result {
		fmt.Printf("Элемент %d: %.4f + %.4fj\n", i, real(val), imag(val))
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
