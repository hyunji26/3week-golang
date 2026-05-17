package main

import "fmt"

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Println("array:", numbers)
	fmt.Println("first number:", numbers[0])

	names := []string{"Alice", "Bob", "Carol"}
	names = append(names, "Dave")

	fmt.Println("slice:", names)
	fmt.Println("slice length:", len(names))

	scores := map[string]int{
		"Alice": 90,
		"Bob":   75,
		"Hyunji": 100,
	}
	scores["Carol"] = 88

	aliceScore, ok := scores["Alice"]
	if ok {
		fmt.Println("Alice score:", aliceScore)
	}

	for name, score := range scores {
		fmt.Printf("%s: %d점\n", name, score)
	}

	total := 0
	for _, score := range scores {
   		 total += score
	}
	avg := float64(total) / float64(len(scores))
	fmt.Printf("평균 점수: %.2f\n", avg)
}
