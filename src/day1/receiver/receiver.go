package main

type Int int

func (n *Int) Inc() {
	*n++
}

func main() {
	var n Int
	println(n)
	n.Inc()
	println(n)
}
