package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var numbers []int64
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var tmp int64
		fmt.Fscan(reader, &tmp)
		numbers = append(numbers, tmp)
	}

	fmt.Println(numbers)

	for i := 0; i < n; i++ {
		var minindex = i

		for j := i + 1; j < n; j++ {
			if numbers[j] < numbers[minindex] {
				minindex = j
			}
		}
		var tmp int64 = numbers[minindex]

		fmt.Println(numbers)
		numbers[minindex] = numbers[i]
		numbers[i] = tmp
	}

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
