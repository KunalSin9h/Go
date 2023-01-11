package main

import "fmt"

// User Defined types
// like typedef in c++
/*
```cpp
typedef long long ll
```
*/
type LinesOfText [][]byte
type Matrix [3][3]int




func main(){

    // byte is uint8
    
    // :)
    var (
    	b = []uint8 {65, 66, 67, 68,  69}
    	c = []uint8  {97,  98, 99, 100, 101}
    	d = []uint8("This is kunal singh")
    )

    fmt.Println(string(b))
    fmt.Println(string(c))
    fmt.Println(string(d))

    var a = Matrix{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    fmt.Println(a[1][1])
    
    v := LinesOfText{
        []byte("First line"),
        []byte("2 line"),
        []byte("3 line"),
        []byte("4 line"),
        []byte("5 line"),
    }

    for _, val := range v {
        fmt.Println(string(val))
    }
    
    var matrix Matrix
    fmt.Println(matrix)
    /*
    for i := 0; i < 3; i++ {
        for  j := 0; j < 3; j++ {
            fmt.Scan(&matrix[i][j])
        }
    }

    fmt.Println(matrix)
    */
}

// Special function which will auto called
func init() {
    fmt.Print("hello form init\n")
}
