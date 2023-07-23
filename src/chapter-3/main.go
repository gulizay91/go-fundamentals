package main

import (
	"fmt"
	"strconv"
)

func concat(s1, s2 string) string {
	return s1 + s2
}

func add(i1 int, i2 int) int {
	return i1 + i2
}

func printer(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

func main() {
	printer("Lane,", " happy birthday!")
	printer("Your age is now ", strconv.Itoa(add(15, 15)))

	/* results:
	Lane, happy birthday!
	Your age is now 30
	*/
}
