package main

import (
	"fmt"
)

/**
If Statement
C++ without ()
*/

func ifs() {

	var x int32

	fmt.Scan(&x)

	if x == 0 {
		fmt.Println("X is 0")
	} else if x > 0 {
		fmt.Println("X is positive")
	} else {
		fmt.Println("X is negative")
	}

	if y := 100; y >= 0 {
		fmt.Println("NaN")
	}

}

/**
Switch Statement
*/

func ss() {

	var x int
	fmt.Scan(&x)

	switch x {
	case 1:
		fmt.Println("You got first rank")
	case 2:
		fmt.Println("You got second rank")
	case 3:
		fmt.Println("You got third rank")
	case 4, 5:
		fmt.Println("You are in top 5")
	case 6, 7, 8, 9, 10:
		fmt.Println("You are in top 10")
	default:
		fmt.Println("You are not in merit list")
	}

	var y int
	fmt.Scan(&y)

	/**
	Without any candidate in switch, switch will just become nested if/else
	because it will just hope inside
	*/

	switch {
	case y == 0:
		fmt.Println("Y is 0")
	case y > 0:
		fmt.Println("Y is positive")
	default:
		fmt.Println("Y is negative")
	}

	var z int = 1

	// ❌ Something is not boiling properly
	switch {
	case z != 10:
		fmt.Println("Yes, Z is not equal to 10")
		fallthrough
	case z == 0:
		fmt.Println("Y is 0")
		fallthrough
	case z > 0:
		fmt.Println("Y is positive")
		fallthrough
	default:
		fmt.Println("Y is negative")
	}

}

func forLoop() {

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

}

func whileLoop() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func rangedBasedForLoop() {
	sentence := "This is Kunal Singh"

	for index, c := range sentence {
		if index%4 == 0 {
			fmt.Println(index, string(c))
		}
	}

	for _, b := range sentence {
		fmt.Println(string(b))
	}

}

func main() {
	/*	ifs()
		ss()
		forLoop()
		whileLoop()*/
	rangedBasedForLoop()
}
