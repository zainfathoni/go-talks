package main

// START Dog OMIT

type Dog struct {
	name string // HLname
}

// END Dog OMIT

// START Speak OMIT

func (d Dog) Speak() string { // HLSpeak
	return "Woof!, I am " + d.name
}

// END Speak OMIT

// START main OMIT

func main() {
	dog := Dog{name: "Scooby"}
	println(dog.Speak()) // HLSpeak
}

// END main OMIT
