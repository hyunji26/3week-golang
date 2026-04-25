package main

import "fmt"

// 구조체 정의
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // 구조체 생성 방법 1: 필드 이름 지정
    p1 := Person{
        Name: "Alice",
        Age:  25,
        City: "Seoul",
    }

    // 구조체 생성 방법 2: 순서대로 값 지정
    p2 := Person{"Bob", 30, "Busan"}

    fmt.Println(p1)
    fmt.Println(p2)

    // 필드 접근
    fmt.Println("이름:", p1.Name)
    fmt.Println("나이:", p1.Age)

    // 필드 수정
    p1.Age = 26
    fmt.Println("변경된 나이:", p1.Age)
}