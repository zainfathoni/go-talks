package main

import (
	"strconv"
)

// START PrintBool OMIT

func PrintBool(s string) {
	if b, err := strconv.ParseBool(s); err == nil {
		println(b)
	} else {
		println(err.Error())
	}
}

func main() {
	PrintBool("true")
	PrintBool("ture")
}

// END PrintBool OMIT
