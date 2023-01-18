package main

import "fmt"

func main(){
    
    c := make(chan int)

    // Built in function to close the channel
    // Should be done with senders side
    close(c)

    // Complex Numbers
    complx := complex(4.2, 80.089)

    r := real(complx)
    i := imag(complx)

    fmt.Println(complx)
    fmt.Println(r, i)

    // copy
   //The copy built-in function copies elements from a source slice 
   // into a destination slice. 
   a := []int{1, 2, 3}
   var b []int = []int{1, 3, 4, 5, 6, 7, 8}

   copy(b, a)

   fmt.Println(b)

   print("Hello\n")
}
