package main

import (
	"fmt"
)

/*
Generic Function
func name[T interface{}](a, b T) T {
    // code...
}
*/

func main(){
    var x any
    x = 10
    fmt.Println(x)
    x = "kunal"
    fmt.Println(x)
    x = 10.20
    fmt.Println(x)
    
    fmt.Println(Min(1, 2))
    fmt.Println(Min(1.1, 2.2))
    fmt.Printf("%c\n", Min('d', 'z'))

    var r rune
    fmt.Printf("%T", r) // int32

}
    

// Type Set for Generic
type numbers interface {
    int | float64 | rune
}

//             > Type Constrained
func Min[T numbers](a, b T) T {
    if a < b {
        return a
    }
    return b
} 
