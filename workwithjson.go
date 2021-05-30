package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HasCar bool `json:"has_car"`

}


func main(){
	myjsonfile:=`[{
"first_name":"dipen",
"last_name":"thapa",
"has_car":true },
{ "first_name":"jain",
"last_name":"mallik",
"has_car":false }]`

var x []Person
err:=json.Unmarshal([]byte(myjsonfile),&x)
if err!=nil{
	log.Println("error occured",err)
}
log.Println("unmarshal: %V",x)

var myslice []Person
var p1 Person
p1.FirstName="dipendraaa"
p1.LastName="thapaaa"
p1.HasCar=false
var p2 Person
p2.FirstName="chandraaa"
p2.LastName="bahadur"
p2.HasCar=false

myslice=append(myslice,p1)
myslice=append(myslice,p2)
log.Println(myslice)
y,err:=json.MarshalIndent(myslice,"","         ")
if err!=nil {
	log.Println("error",err)
}
log.Println(string(y))
}
