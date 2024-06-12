package main

import ("fmt")

func main() {
	// Accessing, changing, appending and copying slices

	// Accessing elements of a slice.
	// Simply refer to the index number, ZERO indexed.
	
	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Changing the elements of a slice
	// Also, simply refer to the index of the element we want to change
	prices[2] = 50

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Appending elements to a slice
	// This can be achieved by using the append() function

	fmt.Println("Appending to the Prices Slice")
	fmt.Println("Original prices array:", prices)

	prices = append(prices, 100,400)
	
	fmt.Println("Updated Prices:", prices)
	fmt.Printf("Length: %d\n", len(prices))
	fmt.Printf("Capacity: %d\n", cap(prices))

	// Appending a slice to another slice
	// Syntax: slice = append(slice1, slice2...) -> The dots are necessary
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	slice3 := append(slice1, slice2...)

	fmt.Printf("slice3 = %v\n", slice3)
	fmt.Printf("Length = %d\n", len(slice3))
	fmt.Printf("Capacity = %d\n", cap(slice3))

	// Changing the length of a slice. 
	// Unlike arrays, this is very much possible in slices
	arr1 := [6]int{9, 10, 11, 12, 13, 14}
	mySlice1 := arr1[1:5] //here, we are slicing the array

	fmt.Printf("mySlice1 = %v\n", mySlice1)
	fmt.Printf("length = %d\n", len(mySlice1))
	fmt.Printf("capacity = %d\n", cap(mySlice1))

	// changing the length by re-slicing the array
	mySlice1 = arr1[1:3]
	
	fmt.Printf("mySlice1 = %v\n", mySlice1)
	fmt.Printf("length = %d\n", len(mySlice1))
	fmt.Printf("capacity = %d\n", cap(mySlice1))

	// changing the length by appending items
	mySlice1 = append(mySlice1, 20, 21, 22, 23)
	fmt.Printf("mySlice1 = %v\n", mySlice1)
	fmt.Printf("length = %d\n", len(mySlice1))
	fmt.Printf("capacity = %d\n", cap(mySlice1))
}