package main
import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile(
		"app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := fmt.Fprintln(f,"employee created"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Log message written successfully")
}