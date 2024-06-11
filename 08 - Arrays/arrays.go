package main

import ("fmt")

func main(){
	var arr1 = [3] int{1, 2, 3}
	arr2 := [5] int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// Declaring variables with inferred length
	var arr3 = [...] int{10, 20, 30}
	arr4 := [...]int{40, 50, 60, 70, 80}

	fmt.Println(arr3)
	fmt.Println(arr4)

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}

	fmt.Println(cars)

	// Accessing Elements of an Array ->Possible by referring to the index number - ZERO INDEXED
	prices := [3]int{10, 30, 80}
	fmt.Println(prices[0])
	fmt.Println(prices[1])
	fmt.Println(prices[2])

	// What if we want to change the elements of an array? Refer to the index number
	prices[1] = 100
	fmt.Println(prices)

	// Array initialization
	// If an array/one of its elements has not been initialize, it is assigned the default value of its type: int=0 and str=""
	arr5 := [5]int{} //Not initialized.
	arr6 := [5]int{1, 2} //partly initialized
	arr7 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	// Initializing only specific elements of an array
	arr8 := [5]int{1:10, 2:50} // the index 1 is assigned 10, and index 2 assigned 50
	fmt.Println(arr8)

	// Finding the length of an array
	employees := [...]string{"Joel", "James", "John", "Jacob"}
	days := [5]int{1, 2, 3, 4, 5}

	fmt.Println(len(employees))
	fmt.Println(len(days))
}