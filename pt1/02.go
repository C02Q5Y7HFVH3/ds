package main

import "fmt"

func main() {
	size := 0
	fmt.Scan(&size)
	data := make([]int, size)
	isNegative := true
	for i := 0; i < size; i++ {
		fmt.Scan(&data[i])
		if data[i] >= 0 {
			isNegative = false
		}
	}

	thisSum, maxSum := 0, 0
	thisSub := make([]int, 0)
	maxSub := make([]int, 1)
	for _, v := range data {
		thisSum += v
		thisSub = append(thisSub, v)
		if thisSum > maxSum {
			maxSum = thisSum
			maxSub = append([]int{}, thisSub...)
		} else if thisSum < 0 {
			thisSum = 0
			thisSub = []int{}
		}
	}

	if isNegative {
		fmt.Println(0, data[0], data[size-1])
	} else {
		fmt.Println(maxSum, maxSub[0],
			maxSub[len(maxSub)-1])
	}
}
