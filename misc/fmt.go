package main

import "fmt"

/*
Go Standard Library

    Strings
*/

//const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
var sample  string

func init(){
    sample = "\xe2\x8c\x98"
}

func main(){
    /*
    Any Thing (value)
    %v

    Boolean
    %t

    Binary , Decimal, Ocatal,  HexaDecimal
    %b       %d       %o (%O)     %x (%X)

    Character
    %c
    'a' | 97
    */
    fmt.Printf("%b\n", 10)
    fmt.Printf("%d\n", 10)
    fmt.Printf("%O\n", 10)
    fmt.Printf("%X\n", 130489)

    /*

    Syntafif Notation
    %e 2.330000e+01
    %E 2.330000E+01

    Float 
    %f (%F)

    When Small Number then %f and when large then %e
    %g -> small e
    %G -> for large E

    */

    fmt.Printf("%g\n", 23.3)
    // x.yf x-> width y->precision
    fmt.Printf("%.10f\n", 10000000000000000000000000.0)
    fmt.Print(1E5, "\n")


    /*
    String
    %s

    Pointers
    %p
    */
    
    var ptr *int

    ptr = new(int)

    *ptr = 10000000

    fmt.Println("Address of new int:", ptr)
    fmt.Println("Address of pointer it self:", &ptr)

    fmt.Printf("%p\n", &ptr)

}
