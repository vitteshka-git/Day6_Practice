package main
import("fmt"
"reflect")

type Employee struct {
	ID int
	Name string
	Salary float64
}

func main() {
	employee := Employee {
		ID := 101,
		Name := "Vitti",
		Salary := 30000.00,
	}
	t := reflect.TypeOf(employee)

	fmt.Println("Type:" ,t)
	fmt.Println("Number of Fileds:" ,t.NumField())

	for i:=0; i < t.NumField() ; i++{
		field := t.field(i);

		fmt.Println("Field:",field.Name)
		fmt.Println("Type:",field.Type)
	}
}