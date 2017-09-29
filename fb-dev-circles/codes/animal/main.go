package main

// START Dog OMIT

type Dog struct {
	name string
}

// END Dog OMIT

// START Speak OMIT

func (d Dog) Speak() string { // HLSpeak
	return "Woof!"
}

// END Speak OMIT

// START Bark OMIT

func Bark(d Dog) {
	println(d.Speak())
}

// END Bark OMIT

// START Animal OMIT

type Animal interface {
	Speak() string // HLSpeak
}

// END Animal OMIT

// START SaySomething OMIT
func SaySomething(a Animal) {
	println(a.Speak())
}

// END SaySomething OMIT

func main() {
	dog := Dog{name: "Charlie"}
	SaySomething(dog)
}
