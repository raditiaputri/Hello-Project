package main

import "fmt"

func main() {

	var n, a, r, i, f int
	fmt.Scanln(&n, &a, &r)
	fmt.Print(0)

	for i = 1; i <= n; i = i + 1 {
		f = a * i * r
		fmt.Print(" + ", f)
	}
}
