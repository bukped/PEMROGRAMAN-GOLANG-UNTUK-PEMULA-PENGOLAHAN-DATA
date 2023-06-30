package main
import "fmt"

// Definisikan struct "Person"
type Person struct {
	Name string
	Age  int
}

// Definisikan struct "Employee" dengan "Person" sebagai field
type Employee struct {
	Person
	Salary float64
}

func main() {
	// Buat instance struct "Employee"
	emp := Employee{
		Person: Person{
			Name: "John Doe",
			Age:  30,
		},
		Salary: 5000,
	}

	// Mengakses properti struct "Person" melalui struct "Employee"
	fmt.Println("Name:", emp.Name)
	fmt.Println("Age:", emp.Age)
	fmt.Println("Salary:", emp.Salary)
}
