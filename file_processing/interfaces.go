package interfaces

import "fmt"

type Session interface {
	ChargeStart()
	ChargeEnd()
}

type EvChargeSession struct {
	Id    string
	Watts int
}

type HybridChargeSession struct {
	Id    string
	Watts int
	Hp    int
}

func (p EvChargeSession) ChargeStart() { // this implements the ChargeStart() method
	fmt.Println("Charging Session initiated for ev", p.Id)
}

func (p EvChargeSession) ChargeEnd() {
	fmt.Println("Charging Session ended for ev", p.Id)
}

func (p HybridChargeSession) ChargeStart() { // this implements the ChargeStart() method
	fmt.Println("Charging Session initiated for plug-in hybrid", p.Id)
}

func (p HybridChargeSession) ChargeEnd() { // this implements the ChargeEnd() method
	fmt.Println("Charging Session ended for plug-in hybrid", p.Id)
}

func ChargeInitiate(cs Session) { // this takes type of our interface
	cs.ChargeStart()
}

func ChargeTerminate(cs Session) { // this takes type of our interface
	cs.ChargeEnd()
}

func main() {
	cs := EvChargeSession{Id: "1111", Watts: 420}
	ChargeInitiate(cs)

	csh := HybridChargeSession{"5555", 120, 120}
	ChargeInitiate(csh)
	ChargeTerminate(csh)
	ChargeTerminate(cs)
}
