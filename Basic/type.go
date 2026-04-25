package main

import "fmt"

func main() {
    var i int     = 42          // 정수
    var f float64 = 3.14        // 실수
    var b bool    = true        // 불리언
    var s string  = "Gopher"   // 문자열

    fmt.Printf("int:     %d\n", i)
    fmt.Printf("float64: %.2f\n", f)
    fmt.Printf("bool:    %t\n", b)
    fmt.Printf("string:  %s\n", s)
}