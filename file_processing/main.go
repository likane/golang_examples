package main

import "fmt"

type ChargeSessionDetails struct {
	id    string
	watts int
	level int // level 2, level3 etc
}

type ChargingSession struct {
	Id      string
	Watts   int
	Details ChargeSessionDetails
}

func newChargeSessionLevel2(id string, watts int) *ChargingSession {
	cs := ChargingSession{Details: ChargeSessionDetails{id: id, watts: watts, level: 2}}
	return &cs
}

func (c ChargingSession) NotifyEndSession() {
	fmt.Println("sending notification to id", c.Id, "end of session", c.Watts)
}

func (c *ChargingSession) IncrementChargeAmount() {
	c.Watts += 10
}

func NewChargeSession(id string, watts int) *ChargingSession {
	cs := ChargingSession{Id: id, Watts: watts}
	return &cs
}

func main() {
	cs := ChargingSession{Id: "11111", Watts: 420}
	cs2 := ChargingSession{"22222", 390}
	fmt.Println(cs, "and", cs2)
	cs.IncrementChargeAmount()
	cs.IncrementChargeAmount()
	cs2.IncrementChargeAmount()
	cs.NotifyEndSession()
	cs2.NotifyEndSession()

	cs3 := NewChargeSession("4444", 500)
	fmt.Println(cs3)
}
