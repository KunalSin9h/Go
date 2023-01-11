package main

import "fmt"

func callback(a func(x int, y int) int ){
    b := a(1, 2)
    fmt.Println(b)
}

func main(){
    func(){
        fmt.Println("Function Literal")
    }()

    callback(func(a int, b int) int {
        return a - b
    })
}
