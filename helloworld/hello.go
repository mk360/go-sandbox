package main

import "fmt"

type S struct {
	X int
}

type Dummy struct {
	UserId    float32 `json:"userId"`
	Id        int64   `json:"A"`
	Title     string  `json:"title"`
	Completed bool    `json:"completed"`
}

func main() {
	var total int8 = 0
	for {
		recursive(&total)
		if total == 0 {
			fmt.Println("Looped back at zero")
		}
	}
}

func recursive(currentCount *int8) {
	(*currentCount)++
}
