package main

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

// START Animal OMIT

type Animal interface {
	Speak() string // HLSpeak
}

// END Animal OMIT

// START SaySomething OMIT

func SaySomething(a Animal) { // HLSaySomething
	println(a.Speak())
}

// END SaySomething OMIT

// START main OMIT
func main() {
	dog := Dog{name: "Charlie"}
	SaySomething(dog) // HLSaySomething
}

// END main OMIT
