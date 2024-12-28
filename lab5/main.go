package main

import (
	"fmt"
)

func main() {
	Array_A := [...]int{1, 2, 3}
	Array_B := [...]int{4, 5, 6}
	Na := len(Array_A)
	Nb := len(Array_B)
	Array_C := make([]int, Na+Nb)

	for i := 0; i < Na; i++ {
		for j := 0; j < Nb; j++ {
			Array_C[i+j+1] += Array_A[i] * Array_B[j]
		}
	}

	for i := len(Array_C) - 1; i > 0; i-- {
		if Array_C[i] >= 10 {
			Array_C[i-1] += Array_C[i] / 10
			Array_C[i] %= 10
		}
	}

	k := 0
	for k < len(Array_C) && Array_C[k] == 0 {
		k++
	}

	if k == len(Array_C) {
		fmt.Println("0")
		return
	}

	fmt.Println(Array_C[k:])
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
