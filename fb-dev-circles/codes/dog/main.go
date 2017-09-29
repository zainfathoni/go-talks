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

func main() {
	dog := Dog{name: "Charlie"}
	Bark(dog)
}
