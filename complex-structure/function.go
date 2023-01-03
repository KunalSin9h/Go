package main

import "fmt"

func getCoupleNames() (he string, she string) {
	he = "Kunal"
	she = "Girl"
	return
}

/*
Spreed Operator
*/
func seeArray(elem ...int) {
	fmt.Println(elem)
	for _, val := range elem {
		fmt.Print(val, " ")
	}
}

func sumOfThree(a, b, c int) (sum int) {
	sum = a + b + c
	return
}

func average(a, b, c float64) (avg float64) {
	avg = (a + b + c) / 3
	return
}

func main() {

	fmt.Println(sumOfThree(10, 20, 30))

	seeArray(1, 2, 3, 4, 5, 10, 2)

	fmt.Println()

	fmt.Println(average(10.11, 10.11, 10.11))

}
