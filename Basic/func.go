package main

import "fmt"

// 기본 함수
func greet(name string) string {
    return "안녕하세요, " + name + "님!"
}

// 다중 반환값 (Go의 특징)
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("0으로 나눌 수 없습니다")
    }
    return a / b, nil
}

func main() {
    msg := greet("Alice")
    fmt.Println(msg)

    result, err := divide(10, 3)
    if err != nil {
        fmt.Println("에러:", err)
    } else {
        fmt.Printf("결과: %.2f\n", result)
    }

    // 에러 케이스
    _, err = divide(5, 0)
    if err != nil {
        fmt.Println("에러:", err)
    }
}