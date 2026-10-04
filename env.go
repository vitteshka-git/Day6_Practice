package main
import ("fmt"
"log"
"os"
"github.com/joho/godotenv"

)

func main() {

	// Load.env file

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}


	appName := os.Getenv("APP_NAME")
	port := os.Getenv("APP_PORT")
	dataFile := os.Getenv("DATA_FILE")

	// if name == "" {
	//name = "Employee Managenment"
	//}

	//if port == "" {
	//	port = "8080"
	//}

	fmt.Println("Application Name:" ,appName)
	fmt.Println("Application Port:" ,port)
	fmt.Println("Data File:" ,dataFile)
	
}