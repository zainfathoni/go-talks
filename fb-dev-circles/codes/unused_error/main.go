package main

import (
	"strconv"
)

func main() {
	b, err := strconv.ParseBool("true")
	println(b)
}
