package main

import (
	"errors"
	"log"
)

func main()  {
	m,n:=testdividebyzero(12,0)
	log.Println(m,n)

}
func testdividebyzero(x,y float32) (float32,error)  {
	var result float32
	if y==0 {
		return result,errors.New("division by zero error")
     }
    result=x/y
    return result,nil
 }