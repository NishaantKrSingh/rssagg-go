package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("Hello World!")
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("Port not found in the environment!") //Immediately stops code and also outputs the proper date and time
	}
	fmt.Println("Port:", portString)
}