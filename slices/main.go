package slices

import "fmt"

func Describe() {
	fmt.Printf("Slices\n\n")

	slice := make([]int, 2, 5)
	fmt.Printf("slice => %v", slice)
	fmt.Printf("\nslice len => %v", len(slice))
	fmt.Printf("\nslice capacity (default) => %v", cap(slice))

	slice = append(slice, 10, 20, 30, 40, 50)
	fmt.Printf("\n\nslice (after append) => %v", slice)
	fmt.Printf("\nslice len => %v", len(slice))
	fmt.Printf("\nslice capacity (2x after append)=> %v", cap(slice))

	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Printf("\n\nslice (after second append) => %v", slice)
	fmt.Printf("\nslice len => %v", len(slice))
	fmt.Printf("\nslice capacity (2x after second append)=> %v", cap(slice))

	fmt.Printf("\n\nHint: Slice capacity is doubled after append that exceeds current capacity\n")

	slice = make([]int, 2, 5)
	fmt.Printf("\nslice => %v\n", slice)
	fmt.Printf("\nfor i := 0; i < 10; i++ => append(slice, i)\n")

	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		fmt.Printf("\nLen: %v, Cap: %v", len(slice), cap(slice))
	}

	fmt.Printf("\n\n\nSlice pointers\n")

	slice = make([]int, 2, 5)
	pointedSlice := slice

	fmt.Printf("\nslice => %v", slice)
	fmt.Printf("\npointed slice => %v", pointedSlice)

	slice[0] = 10

	fmt.Printf("\n\nslice (after setting x[0] to 10) => %v", slice)
	fmt.Printf("\npointed slice (same memory address) => %v", pointedSlice)

	slice = append(slice, 20, 30, 40, 50, 60)
	slice[0] = 0

	fmt.Printf("\n\nslice (after append that exceeds capacity) => %v", slice)
	fmt.Printf("\nold pointed slice (different memory address because go had created a new array to slice) => %v", pointedSlice)

	fmt.Printf("\n\nShorthand slice declarator")
	stringSlice := []string{
		"Hello",
		"World",
		"Good",
		"Morning",
		"!",
	}
	fmt.Printf(`

stringSlice := []string{
  "Hello",
  "World",
  "Good",
  "Morning",
  "!",
}
	`)
	fmt.Printf("\nstringSlice = %v", stringSlice)

	fmt.Printf("\n\nstringSlice[0] => %v", stringSlice[0])
	fmt.Printf("\nstringSlice[:2] (until 2) => %v", stringSlice[:2])
	fmt.Printf("\nstringSlice[2:4] (starting from 2, until 4 - 2 not included) => %v", stringSlice[2:4])
	fmt.Printf("\nstringSlice[2:] (starting from 2, to infinite - 2 not included) => %v", stringSlice[2:])
}
