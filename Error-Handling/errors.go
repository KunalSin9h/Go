package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func isGreaterThenTen(num int32) error {
	if num < 10 {
		return errors.New("something went wrong")
	}
	return nil
}

func openConfigFile() error {
	file, err := os.Open("config.znip.js")

	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}

func main() {

	if err := isGreaterThenTen(9); err != nil {
		//fmt.Println(fmt.Errorf("%d is not greater then 10", num))
		//panic(err)
		log.Println("Wow Build in logging module")
		//log.Fatalf("%d is not greater then 10", num)
	} else {
		fmt.Printf("%d is greater then 10", 9)
	}

	if err := openConfigFile(); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

}
