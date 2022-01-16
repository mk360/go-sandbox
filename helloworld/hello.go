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
	var total int = 0
	defer func() {
		fmt.Println(total)
	}()
	recursive(&total)
}

func recursive(currentCount *int) {
	(*currentCount)++
	recursive(currentCount)
}
