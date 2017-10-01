package main

type Dog struct {
	name string // HLname
}

func (d Dog) Speak() string {
	return "Woof!, I am " + d.name
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
	dog := Dog{name: "Scooby"}
	SaySomething(dog) // HLSaySomething
}

// END main OMIT
