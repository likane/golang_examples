package main

import "fmt"

type ChargingSession struct {
	Id    string
	Watts int
}

func ModifySession(cs *ChargingSession) {
	cs.Watts += 10
}

func main() {
	cs := ChargingSession{"11111", 420}
	s := &cs
	fmt.Println("ID:", (*s).Id)    // one way too dereference
	fmt.Println("Watts:", s.Watts) // other more common way

	ModifySession(s)
	fmt.Println("After modify, Watts:", s.Watts)

	t := new(ChargingSession) // allocate a ChargingSession struct using new and assign to pointer t
	t.Id = "22222"
	t.Watts = 389
	fmt.Println("ID:", t.Id)
	fmt.Println("Watts:", t.Watts)
}
