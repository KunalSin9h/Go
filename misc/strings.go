package main

import (
    "fmt"
    "strings"
)

func main(){
    
    url := "picture/dsomething/39fh"
    newUrl := strings.Join(strings.Split(url, "/"), " ")
    fmt.Println(strings.ToUpper(newUrl))
}
