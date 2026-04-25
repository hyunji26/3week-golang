package main

import "fmt"

func main() {
    // 슬라이스 (동적 배열)
    names := []string{"Alice", "Bob", "Carol"}
    names = append(names, "Dave")   // 요소 추가

    fmt.Println("전체:", names)
    fmt.Println("첫 번째:", names[0])
    fmt.Println("길이:", len(names))

    // 맵 (key-value 저장소)
    scores := map[string]int{
        "Alice": 90,
        "Bob":   75,
    }
    scores["Carol"] = 88   // 추가

    fmt.Println("Alice 점수:", scores["Alice"])

    // 맵 순회
    for name, score := range scores {
        fmt.Printf("%s: %d점\n", name, score)
    }
}