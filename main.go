package main

import (
	"fmt"
	"log"
)

func main() {
	var y string
	var v int
	v = 12

	y, _ = dipen("world")

	fmt.Println(dipen("helowww"))

	log.Println(dipen("helowww"))
	log.Println(v, y)
	log.Println(dipen("helowww"))

}

func dipen(dip string) (string, int) {
	return dip, 11
}
