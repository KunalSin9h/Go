package main

import (
    "fmt"
    "time"
    "sync"
)

var wg sync.WaitGroup

const COLOR_OFF = "\033[0m" 
const RED = "\033[1;31m"
const GREEN = "\033[0;32m"
const  BOLD_GREEN = "\033[1;32m"

func Red(){
    defer wg.Done()
    for i := 1; i <= 3; i++ {
        fmt.Printf("%sRed%s\n", RED, COLOR_OFF)
        time.Sleep(1 * time.Second)
    }
}

func Green(){
    defer wg.Done()
    for i := 1; i <= 3; i++ {
        fmt.Printf("%sGreen%s\n", BOLD_GREEN, COLOR_OFF)
        time.Sleep(3 * time.Second)
    }
}

func main(){
    wg.Add(2) // WaitGroup Counter (it's like how many time we have to call
//                 wg.Done(), after which wg will not block
    go Red()
    go Green()
    wg.Wait() // Blocking line
}
