package main
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Employee struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {

// converting employee object into Json
	data, err := os.ReadFile("employee.json")

	if err != nil {
		log.Fatal(err)
	}
// create an empty slice of employee

var employee []Employee

// convert json data into employee objects
if err := json.Unmarshal(data, &employee);
err != nil {
	log.Fatal(err)
}

	fmt.Println(employee)
}