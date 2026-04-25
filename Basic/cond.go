package main

import "fmt"

func main() {
    score := 85

    if score >= 90 {
        fmt.Println("A등급")
    } else if score >= 80 {
        fmt.Println("B등급")
    } else if score >= 70 {
        fmt.Println("C등급")
    } else {
        fmt.Println("재수강 필요")
    }
}