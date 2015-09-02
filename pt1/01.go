package main

import "fmt"

func main() {
	size := 0
	fmt.Scan(&size)
	data := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&data[i])
	}
	
	thisSum, maxSum := 0, 0
	for _, v := range data {
		thisSum += v
		if thisSum > maxSum {
			maxSum = thisSum
		} else if thisSum < 0 {
			thisSum = 0
		}
	}

	fmt.Println(maxSum)
}
