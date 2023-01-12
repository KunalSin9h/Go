package main

import "fmt"

func sort(list []int, c chan int){
    for i := 0; i < len(list); i++ {

        least_num := list[i]
        least_num_ind := i
        for j := i + 1; j < len(list); j++ {
            if list[j] < least_num {
                least_num = list[j]
                least_num_ind = j
            }
        }
        if least_num_ind != i {
            temp := list[i]
            list[i] = least_num
            list[least_num_ind] = temp
        }
        c <- list[i]
    }
}

func main(){
    A := []int{4, 5, 1, 3, 2}
    c := make(chan int, len(A))
    
    go sort(A, c)

    func(){
        for i := 0; i < len(A); i++ {
            fmt.Println(<- c)
        }
    }()

    fmt.Scanln()
}
