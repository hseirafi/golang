package main

import "fmt"

func main() {
	a, b := 2, 3
	c := (map[bool]int{true: a, false: a - 1})[a > b]
	fmt.Println(c)
}
