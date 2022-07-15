package main

// https://medium.com/@josueparra2892/state-pattern-in-go-290b13e4e387
///
type state interface {
	requestItem() error
	dispenseItem() error
}

type vendingMachine struct {
	selectAvailable state
	itemRequested   state

	currentState state
}

type selectItemMachineState struct {
	vendingMachine *vendingMachine
}

func main() {

}
