package main

import "fmt"

func ff(a, b, m uint64) uint64 {
	if b == 1 {
		return a % m
	}

	val := ff(a, b / 2, m)
	val = val * val % m

	if b % 2 == 0 {
		return val
	}

	return val * a % m
}

func main() {
	var a, b, m uint64

	fmt.Scanln(&a, &b, &m)

	fmt.Println(ff(a, b, m))
}
