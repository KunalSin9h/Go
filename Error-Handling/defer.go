package main

import "fmt"

func doThings() {
	defer fmt.Println("D")
	defer fmt.Println("E")
	fmt.Println("C")
	defer fmt.Println("D")
	defer fmt.Println("E")
}

func recoverFromPanic() {
	// Recover() is literally recovering from panic
	// it is stoping panic to terminate the program
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	fmt.Println("inRecover")
}

func shoutPanic() {
	defer recoverFromPanic()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		if i == 2 {
			panic("PANIC!")
			fmt.Println("After Panic") // X
		}
	}
}

func shouldItRun() {
	fmt.Println("This is going to run")
}

func main() {

	shoutPanic()
	fmt.Println("Here")
	shouldItRun()
}
