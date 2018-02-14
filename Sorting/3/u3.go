package main

import (
	"fmt"
)

var (
	y       []int
	biggest int
)

func main() {
	var x int
	fmt.Print("Input many array : ")
	fmt.Scan(&x)

	y = make([]int, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&y[i])
		if y[i] >= 10 {
			fmt.Println("Input under 10 Please")
			fmt.Scan(&y[i])
		}
	}

	for _, v := range y {
		if v > biggest {
			biggest = v
		}
	}

	visual(biggest)

	sorting(y)
}

func sorting(array []int) {
	for i := 0; i < len(array)-1; i++ {
		max := i
		for j := i + 1; j <= len(array)-1; j++ {
			if array[j] > array[max] {
				max = j
			}
		}
		swap(array, i, max)
		visual(biggest)
	}
}

func swap(array []int, i, j int) {
	tmp := array[j]
	array[j] = array[i]
	array[i] = tmp
}

func visual(big int) {
	for i := 0; i < big; i++ {
		for j := 0; j < len(y); j++ {
			if i >= big-y[j] {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}

	for _, v := range y {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()
}
