package main

import "fmt"

/**
Slices are like `Vector` in C++
We can achieve it my using make function
*/

func sliced() {
	fruits := [6]string{"apple", "banana", "mango", "lichi", "green", "peach"}

	slic := fruits[0:3]

	fmt.Println(slic)
	fmt.Println(len(slic))
	fmt.Println(cap(slic))

	slic = append(slic, "A", "B", "C", "D", "E", "F", "G", "H", "i", "j")

	fmt.Println(slic)
	fmt.Println(len(slic))
	fmt.Println(cap(slic))

}

func cp() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, 0, n)

	//              len  cap
	a := make([]int8, 3, 4)
	b := make([]int8, 3, 4)

	fmt.Println(&a[0])
	fmt.Println(&b[0])

	for n > 0 {

		var x int
		fmt.Scan(&x)

		arr = append(arr, x)

		n--
	}

	fmt.Println(arr)
}

func main() {

	var myArray [5]int32
	var mySlice = make([]int32, 10, 100)

	myArray[0] = 1
	mySlice[0] = 1

	fmt.Println(myArray)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	sliced()
	cp()
}
