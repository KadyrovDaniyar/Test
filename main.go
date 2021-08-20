package main

import "fmt"

func main() {
	type a [3]int

	var b a
	b[0] = 2
	b[1] = 3
	b[2] = 4

	func(*a) {
		for _, i := range &b {
			fmt.Println(i)
		}
	}(&b)
}
