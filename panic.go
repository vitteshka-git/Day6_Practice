package main
import "fmt"

func main() {

	findEmployee(3)

	// fmt.Println("Program start")
	// panic("Something went wrong")
	// fmt.Println("Program finished")



}

func findEmployee(id int) {

	if id <= 0 {
		panic("Invalid employee ID")
	}
	fmt.Println("Employee ID:", id)
}
