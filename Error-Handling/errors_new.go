package main

import (
    "fmt"
    "errors"
)

func giveError() error  {
    return errors.New("This is Error message")
}

func main(){
    
    new_error := giveError()

    // Both give same message
    fmt.Println(new_error)
    fmt.Println(new_error.Error())

}
