package maps

import "fmt"

func Describe() {
	fmt.Printf("Maps\n")

	myMap := make(map[string]int)

	myMap["a"] = 10
	myMap["b"] = 20
	myMap["c"] = 30

	fmt.Printf("\nmymap => %v", myMap)

	delete(myMap, "c")

	fmt.Printf("\nmymap (after deleting c) => %v", myMap)
	fmt.Printf("\nmymap['c'] (deleted) => %v", myMap["c"])

	_, cExists := myMap["c"]
	_, bExists := myMap["b"]

	fmt.Printf("\n\nc (deleted) exists? %v", cExists)
	fmt.Printf("\nb exists? %v", bExists)

	fmt.Printf("\n\n\nShorthand syntax")

	newMap := map[string]int{
		"x": 5,
		"y": 10,
	}

	fmt.Printf(`

newMap := map[string]int{
  "x": 5,
  "y": 10,
}
	`)

	fmt.Printf("\nnewMap => %v\n", newMap)

	if zValue, zExists := newMap["z"]; zExists {
		fmt.Printf("\nnewMap['z'] => %v", zValue)
	} else {
		fmt.Printf("\nz doesn't exists")
	}
}
