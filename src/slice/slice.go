package main

import "fmt"

func main() {
	ns := []int{19, 86, 1, 12}
	var sum int
	for _, v := range ns {
		sum += v
	}
	fmt.Println(sum)
}
