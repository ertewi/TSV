package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int = 2
	var numbers []int64
	// fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var tmp int64
		fmt.Fscan(reader, &tmp)
		numbers = append(numbers, tmp)
	}

	result := numbers[1]
	speed := int64(1500)
	for i := int64(0); i < numbers[0]; i++ {
		result = result * numbers[1]
		time := result / speed / 60
		fmt.Println(i+2, result, time, "m")
	}

	fmt.Println(numbers)

	fmt.Println(numbers)
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
