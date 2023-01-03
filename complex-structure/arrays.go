package main

import "fmt"

func main() {
	var arr [5]float64

	arr[0] = 1.1
	arr[1] = 1.2
	arr[2] = 1.3
	arr[3] = 1.4
	arr[4] = 1.5

	sgpa := [8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	ratings := [...]float32{4.5, 4.7, 3.9, 4.3}

	for _, v := range sgpa {
		fmt.Print(v, " ")
	}

	fmt.Println()

	fmt.Println(sgpa)
	fmt.Println(ratings)

	fmt.Println(arr)
}
