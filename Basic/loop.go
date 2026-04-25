package main

import "fmt"

func main() {
    // 기본 for (C언어 스타일)
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d번째 반복\n", i)
    }

    // while처럼 사용
    n := 0
    for n < 3 {
        fmt.Println("n =", n)
        n++
    }

    // 슬라이스 순회 (range)
    fruits := []string{"apple", "banana", "cherry"}
    for index, value := range fruits {
        fmt.Printf("[%d] %s\n", index, value)
    }
}