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
	employee := []Employee {
	{
		ID : 1,
		Name : "Vittesh",
		Email : "Vitti2058",
	},
	{
		ID : 2,
		Name : "dhanu",
		Email : "dhanu2058",
	},
	{
		ID : 1,
		Name : "guru",
		Email : "guru2058",
	},
}

// converting employee object into Json
	data, err := json.MarshalIndent(employee, "" , " ")

	if err != nil {
		log.Fatal(err)
	}
// write json data into employee.json

if err := os.WriteFile("employee.json", data, 0600);
err != nil {
	log.Fatal(err)
}
	fmt.Println("employee.json created successfully")

// convert [] byte into string and print
	fmt.Println(string(data))
}