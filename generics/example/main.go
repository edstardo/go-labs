package main

import "fmt"

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Sum[N Number](nums []N) N {
	var sum N
	for _, n := range nums {
		sum += n
	}

	return sum
}

func Average[N Number](nums []N) float64 {
	sum := float64(Sum(nums))

	return sum / float64(len(nums))
}

func main() {
	fmt.Println(Average([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(Average([]int32{1, 2, 3, 4, 5, 6}))
	fmt.Println(Average([]int64{1, 2, 3, 4, 5, 6}))
	fmt.Println(Average([]float32{1, 2, 3, 4, 5, 6}))
	fmt.Println(Average([]float64{1, 2, 3, 4, 5, 6}))
}
