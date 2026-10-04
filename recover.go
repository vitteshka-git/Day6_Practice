package main 
import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:" ,r)
		}
	} ()
	fmt.Println("Program start")
	panic("Something went wrong")
	fmt.Println("Program finished")
}