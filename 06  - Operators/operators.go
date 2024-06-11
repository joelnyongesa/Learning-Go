// Operators: Operating symbols in Go that instruct the compiler to perform specific operations on operands.
package main

import ("fmt")

/*
+ -> Addition
- -> Subtraction
* -> Multiplication
/ -> Division
% -> Modulus operator
*/

func main() {
	const a = 10
	const b = 3
	// Arithmetic operators
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)


	// Comparison operators
	/*
		Used to evaluate whether the given values of two operands are equal, or whether one valu e is greater than or less than the other.
		Results: true or false.
		These operators include:
			== (Equal operator)
			!= (Not Equal Operator)
			> (Greater than operator)
			< (Less than operator)
			>= (Greater than or equal to operator)
			<= (Less than or equal to operator)
	*/

	fmt.Println("Is a == b?", a == b)
	fmt.Println("Is a != b?", a != b)
	fmt.Println("Is a < b?", a < b)
	fmt.Println("Is a <= b?", a <= b)
	fmt.Println("Is a > b?", a > b)
	fmt.Println("Is a >= b?", a >= b)


	// Assignment Operators -- Used to assign values to a variable.
	/*
		= (Simple assignment operator)	
		+= (Addition assignment operator)
		-= (Substraction assignment operator)
		*= (Multiplication assignment operator)
		/= (Division assignment operator)
		%= (Modulus assignment operator)
	*/

	fmt.Printf("The value of a is %v, and the value of b is %v\n", a, b)
	fmt.Println("*******************")
	c := a;
	fmt.Println("The value of C is: ", c)

	c += b
	fmt.Println("The value of c += b is: ", c)

	c -= b
	fmt.Println("The value of c -= b is: ", c)

	c *= b
	fmt.Println("The value of c *= b is: ", c)

	c /= b
	fmt.Println("The value of c/=b is: ", c)

	c %= b
	fmt.Println("The value of c %= b is: ", c)
}