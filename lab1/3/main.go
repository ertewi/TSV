package main

import (
	"bufio"
	"fmt"
	"os"
)

func merge(left, right []int64) []int64 {
	result := []int64{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

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

	for k := 1; k < n; k *= 2 {
		for start := 0; start < n; start += 2 * k {
			mid := start + k
			end := start + 2*k
			if mid < n {
				if end > n {
					end = n
				}
				merged := merge(numbers[start:mid], numbers[mid:end])
				copy(numbers[start:start+len(merged)], merged)
				fmt.Println(numbers)
			}
		}
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
