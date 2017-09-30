package main

// START Engine Body OMIT

type Engine interface { // HLEngineBody
	Refill()
}

type Body interface { // HLEngineBody
	Load()
}

// END Engine Body OMIT

type ElectricEngine struct{}

func (e ElectricEngine) Refill() {
	println("Recharge Battery.")
}

type CarBody struct{}

func (c CarBody) Load() {
	println("Board Passengers.")
}

// START Vehicle OMIT

type Vehicle struct {
	Engine // HLEngineBody
	Body   // HLEngineBody
}

// END Vehicle OMIT

func main() {
	// START electricCar OMIT

	electricCar := Vehicle{Engine: ElectricEngine{}, Body: CarBody{}}
	electricCar.Refill()
	electricCar.Load()

	// END electricCar OMIT
}
