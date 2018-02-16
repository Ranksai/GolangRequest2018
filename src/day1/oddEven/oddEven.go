package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		// use if
		if i%2 == 0 {
			fmt.Printf("%v - even\n", i)
		} else {
			fmt.Printf("%v - odd\n", i)
		}
		/*
			// use switch
			switch isOdd := i % 2; isOdd {
			case 0:
				fmt.Printf("%v - even\n", i)
			case 1:
				fmt.Printf("%v - odd\n", i)
			}
		*/
	}
}
