package main

import (
	"fmt"
)

func main() {
	var x, biggest int
	fmt.Print("Input many array : ")
	fmt.Scan(&x)

	y := make([]int, x)
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

	for i := 0; i < biggest; i++ {
		for j := 0; j < len(y); j++ {
			if i >= biggest-y[j] {
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
