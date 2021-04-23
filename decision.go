package main

import (
	"fmt"	"log"
)


func main() {
	var x bool
	y:="cat"

	if x {
		fmt.Println("heeel")

		// == for comparison
	} else if !x {
		fmt.Println("geeel")

	} else {

		fmt.Println("feeel")
	}

	switch y{
	case "cat":
		log.Println("it is",y)
	case "dog:
		log.Println("it is" ,y)
	case "rat":
		log.Println("it is",y)
		
	default:
		log.Println("it is","defauylt")

		
	}
}
