package main

import ("fmt"
		"time"
	"runtime")

	func getValue() int {
		return 30
	   }

	   type ChargeSession struct {
		id    string
		watts int
		vin   string
	   }
	   
	   func (cs *ChargeSession) Charge(f int) {
		cs.watts = cs.watts + f
	   
	   }

func main() {
 fmt.Println("Review Go Data Types")

 var x int = 45 / 9. // add type with var
 var y = 7 - x.      // use var, let type be inferred
 z := 32 / 8.        // short version, preferred if can use
 // z := 7.          // this is not be allowed, cannot use := unless new variable
 z = 7               // can use = on a var created with :=
 fmt.Println("Value of x and y and z is", x, "and", y, "and", z)

 // type inference
 var a bool = false           // Boolean
 var b int = 8                // Integer
 var c float64 = 3.145        // Float
 var d string = "Go Is Cool!" // String
 fmt.Println("Initialized with var =")
 fmt.Println("This is a type", reflect.TypeOf(a), "with value", a)
 fmt.Println("This is a type", reflect.TypeOf(b), "with value", b)
 fmt.Println("This is a type", reflect.TypeOf(c), "with value", c)
 fmt.Println("This is a type", reflect.TypeOf(d), "with value", d)

 a2 := false         // Boolean Inferred
 b2 := 8             // Integer Inferred
 c2 := 3.145         // Float Inferred
 d2 := "Go Is Cool!" // String Inferred
 fmt.Println("Initialized with := ")
 fmt.Println("This is a type", reflect.TypeOf(a2), "with value", a2)
 fmt.Println("This is a type", reflect.TypeOf(b2), "with value", b2)
 fmt.Println("This is a type", reflect.TypeOf(c2), "with value", c2)
 fmt.Println("This is a type", reflect.TypeOf(d2), "with value", d2)

 // const and scope
 const name = "Brian"
 // name = "Tim" // can't do this, it is a const
 fmt.Println("Hey", name, "the value of Pi is", Pi)

 count := 0

 for count < 10 { // this is the inner scope
  num := 10 + rand.Intn(10) // num is only visible within these brads
  fmt.Println(num)
  count++                      // count is in scope here as part of outer scope
 }
 // num = 1 // num is no longer in scope here
 fmt.Println("Generated", count, "random numbers") // count is still in scope

 // loops
  // standard 3 part for loop
  for i := 1; i <= 10; i++ {
	fmt.Println(i)
   }

    // create an array to iterate over
 numbers := [5]int{4, 8, 12, 16, 20}

 // for loop with range
 for element := range numbers {
  fmt.Println("Element is", element)
  fmt.Println("Value at that index is", numbers[element])
 }

 numbers := [5]int{4, 8, 12, 16, 20}

 for _, item := range numbers {
  fmt.Println("Value is", item)
 }

 // flow control
 t := time.Now()
 switch {
 case t.Hour() < 12:
  fmt.Println("Hello, have a wonderful morning!")
 case t.Hour() < 17:
  fmt.Println("Have a great afternoon.")
 default:
  fmt.Println("Wish you a very nice evening.")
 }

 switch os := runtime.GOOS; os {
case "darwin":
 fmt.Println("Running on Mac OS with", runtime.GOARCH)
case "windows":
 fmt.Println("Running on Windows with", runtime.GOARCH)
case "linux":
 fmt.Println("Running on Linux with", runtime.GOOS)
default:
 // others
 fmt.Printf("Running on %s with %s\n", os, runtime.GOOS)
}

limit := 20
if x := getValue(); x >= limit {
 fmt.Println("Value", x, "is greater equal", limit)
} else {
 fmt.Println("Value", x, "is less", limit)
}

	// defer
	defer fmt.Println("Deferred print")
	cs := ChargeSession{"1111", 410, "UNKNOWN"}
	cs.Charge(10)
	fmt.Println(cs.watts)
}