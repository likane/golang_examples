package main

import (
	"fmt"
	"math/rand"
)

func main() {
	msgNum := randInt(1, 6)
	fmt.Println("Generated a", msgNum)
	printMessage(msgNum)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func printMessage(msgNum int) {
	if msgNum == 1 {
		fmt.Println("Printing this for a 1")
	} else if msgNum == 2 {
		fmt.Println("Printing this for a 2")
	} else if msgNum == 3 {
		fmt.Println("Printing this for a 3")
	} else if msgNum == 4 {
		fmt.Println("Printing this for a 4")
	} else if msgNum == 5 {
		fmt.Println("Printing this for a 5")
	} else {
		fmt.Println("Whoops")
	}
}
