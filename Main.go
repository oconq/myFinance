package myFinance

import (
	"fmt"
	"os"
	"awesomeProject/config"
	"log"
	"net/http"
)


var text_to_display = "Hello world %s!";
var user = os.Getenv("USER")

/* TEST */
func main() {
	fmt.Printf("Tentative de connection...")
	config.DatabaseInit()
	fmt.Println("OK")

	fmt.Printf("Initialisation des routes...")
	router := InitializeRouter()
	fmt.Println("OK")

	fmt.Println("Ecoute...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
