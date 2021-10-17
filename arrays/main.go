package arrays

import "fmt"

func Describe() {
	fmt.Printf("Arrays\n\n")

	var x [10]int
	fmt.Printf("x => %v", x)
	fmt.Printf("\nlength => %v", len(x))

	fmt.Printf("\n")

	x[0] = 10
	x[1] = 5
	x[2] = 20

	fmt.Printf("\nx => %v", x)
	fmt.Printf("\nx[0] => %v", x[0])
	fmt.Printf("\nx[1] => %v", x[1])
	fmt.Printf("\nx[2] => %v", x[2])

	fmt.Printf("\n")

	z := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("\nz => %v", z)
}
