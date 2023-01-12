package main

import (
	"fmt"
	"time"
)

const COLOR_OFF = "\033[0m" 
const RED = "\033[1;31m"
const GREEN = "\033[0;32m"
const  BOLD_GREEN = "\033[1;32m"

func main(){
    /*
    var channel_G01 chan string = make(chan string, 2)
    channel_G01 <- "rV7zzV5829LA2SFY" // Send is block
    channel_G01 <- "d9f" // Send is block
    channel_G01 <- "d9e" // Send is block
    retrived := <- channel_G01 // So do receive
    fmt.Println(retrived)
    retrived = <- channel_G01 // So do receive
    fmt.Println(retrived)
    */

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for true {
            c1 <- "Red"
            fmt.Printf("%sRed%s in every 0.5 sec\n", RED, COLOR_OFF)
            time.Sleep(time.Millisecond * 500)
        }
    } ()

    go func() {
        for true {
            c2 <- "Green"
            fmt.Printf("%sGreen%s in every 5 sec\n", BOLD_GREEN, COLOR_OFF)
            time.Sleep(time.Second * 5)
        }
    } ()
    
    for true {
        select {
        case msg1 := <- c1:
            fmt.Println(msg1)
        case msg2 := <- c2:
            fmt.Println(msg2)
        }
    }

}
