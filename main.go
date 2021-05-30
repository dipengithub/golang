package main
//go mod init github.com/dipengithub/golang ==>command for package

import (
	"github.com/dipengithub/golang/helpers"
	"fmt"
	"log"
)

func main() {
	var y string
	var v int
    var myVar helpers.SomeType
	myVar.TypeName="dipenddcc"
	log.Println(myVar.TypeName)
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
