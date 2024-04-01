package main

import "fmt"

type largeType [813223]int

type ChargingSession struct {
	Id    string
	Watts int
}

func processLargeStruct(s largeType) largeType {
	fmt.Println(len(s))
	return s
}

func processLargeStructPtr(ps *largeType) *largeType {
	fmt.Println(len(ps))
	return ps
}

func main() {
	var n [2]string
	n[0] = "Brian"
	n[1] = "Enochson"
	fmt.Println(n[0], n[1])
	fmt.Println(n)

	n2 := [2]string{"Brian", "Enochson"}
	fmt.Println(n2[0], n2[1])
	fmt.Println(n2)

	nums := [6]int{6, 9, 3, 41, 33, 55}
	fmt.Println(nums)

	nums2 := [6]int{1: 20, 3: 40, 5: 43} // only 1, 3, and 5
	fmt.Println(nums2)

	nums3 := [...]int{10, 20, 30, 40, 50, 60} // size is inferred, not provided
	fmt.Println(nums3)
	fmt.Println("Number of elements", len(nums3))

	// nums := [6]int{6, 9, 3, 41, 33, 55}
	fmt.Println(nums)

	for i := range nums {
		fmt.Print(i, " ")
	}
	fmt.Println("")

	for i, item := range nums {
		fmt.Print(i, ": ", item, " ")
	}

	fmt.Println("")
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}

	// multidim arrays
	multi := [4][2]int{{1, 10}, {2, 20}, {3, 30}, {4, 40}}
	fmt.Println(multi)

	multi[0][1] = 100
	multi[1][1] = 200
	multi[2][1] = 300
	multi[3][1] = 400
	fmt.Println(multi)

	// pointers to arr
	var s largeType
	s = processLargeStruct(s)

	processLargeStructPtr(&s)

	// slices
	// standard way to initialize a slice
	sliceVar := []int{1, 2, 3, 4, 5}
	fmt.Println("Actual Size is:", len(sliceVar))
	fmt.Println("Capacity is:", cap(sliceVar))

	// using make and the length
	sliceVar2 := make([]int, 5) // identical to []int{0, 0, 0, 0, 0}
	fmt.Println("Via make([]int, 5): ", sliceVar2)
	fmt.Println("Actual Size is:", len(sliceVar2))
	fmt.Println("Capacity is:", cap(sliceVar2))

	// using make and the length and capacity
	sliceVar3 := make([]int, 5, 10)
	fmt.Println("Via make([]int, 5, 10): ", sliceVar3)
	fmt.Println("Actual Size is:", len(sliceVar3))
	fmt.Println("Capacity is:", cap(sliceVar3))

	fmt.Println(len(sliceVar), cap(sliceVar))
	fmt.Println(sliceVar == nil)

	// make slices from arr
	carMakes := [4]string{"Tesla", "Ford", "Volkswagen", "Mercedes"}

	fmt.Println(carMakes)

	germanMakes := carMakes[2:4] // since 4 is the end, can also write
	// carMakes[2:] in this case

	germanMakes = append(germanMakes, "BMW")

	fmt.Println(germanMakes)

	// maps
	chargeSessions := map[string]int{
		"11111": 420, "22222": 395, "33333": 456,
	}
	key := "11111"
	fmt.Println("Value for", key, "is", chargeSessions[key])
	key = "33333"
	fmt.Println("Value for", key, "is", chargeSessions[key])

	cs := make(map[string]int) // Declare and initialize a map with 'make'
	cs["11111"] = 420
	fmt.Println(cs)

	// Declare an empty map with a literal.
	nilMap := map[int]string{} // this is identical to using make with {}
	fmt.Println(nilMap)

	// Declare and populate a map with a literal.
	// Keys separated from values with a colon,
	// and commas separate key-value pairs.
	m := map[string]int{

		"33333": 343, "44444": 410,
	}
	fmt.Println(m)

	var m2 map[string]int
	if m2 == nil {
		fmt.Println("Map is nil, initializing")
		m2 = map[string]int{}
	}

	m2["22222"] = 389 // Assignment succeeds after the map has been initialized.
	fmt.Println(m2)   // without the initialization the assign would cause a panic

	cs1 := map[string]ChargingSession{"11111": ChargingSession{"11111", 420},
		"22222": ChargingSession{"22222", 405},
		"33333": ChargingSession{"33333", 390}}

	fmt.Println(cs1)

	delete(cs1, "22222")
	fmt.Println(cs)
}
