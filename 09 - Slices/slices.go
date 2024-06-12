/*
Arrays -> More fixed in size.
Slices -> Dynamically sized, flexible view into the elements, and much more common in the real world than arrays.
*/

package main

import ("fmt")

func main() {
	// Creating a slice:
		// sUsing the []datatype{values} format
		// Creating a slice from an array
		// Using the make() function

	// Using the []datatype{values} format
	mySlice := []int{1, 2, 3}

	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	mySecondSlice := []string{"Go", "Slices", "Are", "Really", "Powerful"}
	fmt.Println(mySecondSlice)
	fmt.Println(len(mySecondSlice))
	fmt.Println(cap(mySecondSlice))

	// Creating a slice from an array
	myArray := [6]int{10,11,12,13,14,15}
	sliceFromArray := myArray[2:4]

	fmt.Printf("sliceFromArray = %v\n", sliceFromArray)
	fmt.Printf("Length = %d\n", len(sliceFromArray))
	fmt.Printf("Capacity = %d\n", cap(sliceFromArray))

	// Creating a slice with the make() function
	// make([]type, length, capacity). If capacity is not defined, then it will be equal to the length.
	myThirdSlice := make([]int, 5, 10)
	fmt.Printf("myThirdSlice = %v\n", myThirdSlice)
	fmt.Printf("Length = %d\n", len(myThirdSlice))
	fmt.Printf("Capacity = %d\n", cap(myThirdSlice))

	// With omitted capacity -> means length=capacity
	myFourthSlice := make([]int, 6)
	fmt.Printf("myFourthSlice = %v\n", myFourthSlice)
	fmt.Printf("Length = %d\n", len(myFourthSlice))
	fmt.Printf("Capacity = %d\n", cap(myFourthSlice))
}