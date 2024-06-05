package main;
import ("fmt");

func main() {
	// We can also declare multiple variables in Go. If the type is specified, then we can only declare one variable type per line
	var a, b, c, d int = 1, 2, 3, 4;

	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);
	fmt.Println(d);


	// If the type is not specified, then we can declare different types of variables in the same line
	var name, age = "Joel Nyongesa", 25;
	fmt.Println(name);
	fmt.Println(age);


	// We can also group multiple variable declarations together for readability
	var (
		employeeID int = 1;
		employeeName string = "Employee 1";
		isMarried bool = true;
	);
	fmt.Println(employeeID);
	fmt.Println(employeeName);
	fmt.Println(isMarried);
}