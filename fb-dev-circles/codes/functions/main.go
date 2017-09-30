package main

// START Functions OMIT

func Hello(aFunc func(words string) string) {
	println(aFunc("Hello"))
}

func Goodbye(aFunc func(words string) string) {
	println(aFunc("Goodbye"))
}

func main() {
	world := func(words string) string {
		return words + ", world!"
	}
	Hello(world)
	Goodbye(world)
}

// END Functions OMIT
