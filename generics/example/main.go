package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Average[N Number](nums []N) float64 {
	sum := 0.0
	for n := range nums {
		sum += float64(n)
	}
	return sum / float64(len(nums))
}

func main() {
	fmt.Println(Average([]int{2, 1, 4, 5, 3, 7}))
	fmt.Println(Average([]float64{2, 1, 4, 5, 3, 7}))
}
