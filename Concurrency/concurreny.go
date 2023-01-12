package main

import (
	"fmt"
	"time"
)

const (
	COLOR_OFF = "\033[0m" 
	RED = "\033[1;31m"
	GREEN = "\033[0;32m"
	BOLD_GREEN = "\033[1;32m"
)

func main(){
    var c chan int = make(chan int)
    go count(c)

    fmt.Println("Wating for first byte...")
    fmt.Println(<- c)
    fmt.Println("One Already Get")
    
    /*
    for  {
        rec, open := <- c
        if !open {
            break
        }
        fmt.Println(rec)
    }*/
    for res := range c {
        fmt.Println(res)
    }
}

func count(c chan int) {
    for i := 1; i <= 5; i++ {
        if i == 1 {
            time.Sleep(5 * time.Second)
        }
        c <- i
        time.Sleep(500 * time.Millisecond)
    }
    close(c)
}
