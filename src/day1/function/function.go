package main

func isEven(number int) bool {
	return number%2 == 0
}

func swap(x int, y int) (int, int) {
	return y, x
}

func swap2(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}
