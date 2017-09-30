package main

// START Engine Body OMIT

type Engine interface { // HLEngineBody
	Refill()
}

type Body interface { // HLEngineBody
	Load()
}

// END Engine Body OMIT

type PetrolEngine struct{}

func (p PetrolEngine) Refill() {
	println("Refuel Tank.")
}

type TruckBody struct{}

func (t TruckBody) Load() {
	println("Mount Loads.")
}

// START Vehicle OMIT

type Vehicle struct {
	Engine // HLEngineBody
	Body   // HLEngineBody
}

// END Vehicle OMIT

func main() {
	// START petrolTruck OMIT

	petrolTruck := Vehicle{Engine: PetrolEngine{}, Body: TruckBody{}}
	petrolTruck.Refill()
	petrolTruck.Load()

	// END petrolTruck OMIT
}
