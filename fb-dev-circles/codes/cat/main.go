package main

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Animal interface {
	Speak() string
}

func SaySomething(a Animal) {
	println(a.Speak())
}

// START Cat OMIT

type Cat struct {
	name string
}

func (c Cat) Speak() string { // HLSpeak
	return "Miaow!"
}

// END Cat OMIT

// START TestSaySomething OMIT

func TestSaySomething() {
	cat := Cat{name: "Charlie beta"}
	SaySomething(cat) // HLSaySomething
}

// END TestSaySomething OMIT

func main() {
	TestSaySomething()
}
