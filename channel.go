package main

import (
	"github.com/dipengithub/golang/helpers"
	"log"
)
const x=1000
func main()  {
	intChan:=make(chan int)
	defer close(intChan)
	go CalculateValue(intChan)
	go CalculateValue(intChan)
	num1,num2:=<-intChan,<-intChan
	log.Println(num1,num2)
}
func CalculateValue(intChan chan int) {
	rn := helpers.RandomNumber(x)
	intChan<-rn
}


