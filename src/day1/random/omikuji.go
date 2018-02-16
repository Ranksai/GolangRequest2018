package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandom0to5() int {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6)
	return s
}
func main() {
	switch dice := getRandom0to5(); dice {
	case 0:
		fmt.Println("大吉")
	case 1:
		fallthrough
	case 2:
		fmt.Println("中吉")
	case 3:
		fallthrough
	case 4:
		fmt.Println("吉")
	case 5:
		fmt.Println("凶")
	}
}
