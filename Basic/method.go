package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// 메서드: (receiver) 를 func 키워드 뒤에 선언
func (p Person) Introduce() string {
    return fmt.Sprintf("저는 %s이고 %d살입니다.", p.Name, p.Age)
}

func (p Person) IsAdult() bool {
    return p.Age >= 18
}

func main() {
    p := Person{Name: "Alice", Age: 20}

    fmt.Println(p.Introduce())

    if p.IsAdult() {
        fmt.Println(p.Name, "은 성인입니다.")
    }
}