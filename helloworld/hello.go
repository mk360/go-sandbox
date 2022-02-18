package main

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
	var total uint8 = 0
	for {
		recursive(&total)
	}
}

func recursive(currentCount *uint8) {
	_ = Dummy{UserId: 224, Id: 3453, Title: "", Completed: true}
	recursive(currentCount)
}

func bToMb(bytes uint64) uint64 {
	return bytes / 1000 / 1000
}
