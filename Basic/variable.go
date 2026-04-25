package main

import "fmt"

func main() {
    // 방법 1: var 키워드 (타입 명시)
    var name string = "Alice"
    var age int = 25

    // 방법 2: := 단축 선언 (타입 자동 추론, 함수 내부에서만 사용 가능)
    city := "Seoul"
    score := 98.5

    fmt.Println(name, age, city, score)

    // 변수 값 변경
    age = 26
    fmt.Println("변경된 나이:", age)
}