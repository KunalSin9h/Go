package main

import "fmt"

func main(){
    
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    /*
    for i := 0; i < 100; i++ {
        go worker(jobs, results)
    }*/

    // Each worker will accupy 1 core

    go worker(jobs, results)
    go worker(jobs, results)
    go worker(jobs, results)
    go worker(jobs, results)

    for i := 0; i < 100; i++ {
        jobs <- i
    }

    close(jobs)

    for i := 0; i < 100; i++ {
        fmt.Println(<-results)
    }
}

func worker(jobs <-chan int, results chan<- int) {
    for n :=  range jobs {
        results <- Fib(n)
    }
}

func Fib(n int) int {
    if n <= 1 {
        return n
    }
    return Fib(n - 1) + Fib(n - 2)
}

