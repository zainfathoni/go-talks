package main

func main() {
	println("Opening file A...")
	defer println("Closing file A...") // Closing here

	println("Do anything with file A")

	panic("Error happens sometimes")

	// println("Closing file A...") // Instead of here
}
