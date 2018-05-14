package main

import (
	"fmt"
	"math"
)

func main() {
	scores := []int{55, 49, 100, 150, 0}
	sum := 0

	fmt.Println(scores)

	for _, x := range scores {
		sum += x
	}
	fmt.Println(sum)

	average := float64(sum) / float64(len(scores))
	fmt.Println(average)

	/* 分散を求める */
	temp := 0.0
	for _, x := range scores {
		temp += math.Pow(float64(x)-float64(average), 2)
	}
	variance := temp / float64(len(scores))
	fmt.Println(variance)

	/* 標準偏差を求める */
	sd := math.Sqrt(variance)

	fmt.Println(sd)
}
