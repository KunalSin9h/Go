package main

import "fmt"

type Any interface{}

func main() {

	anything := map[Any]Any{
		1:   "2",
		"2": 1,
	}

	fmt.Println(anything)
}
