package main

import (
	"log"
	"time"
)

type Dipen struct {
	First string
}

type User struct {
	FirstName string
	LastName  string
	Location  string
	Phone     string
	BirthDate time.Time
}

// var dipen string
var v = "dipen"

func main() {
	log.Println(v)
	// variable shadwoing capital indicate public and small private
	log.Println(dipen(v))
	log.Println(v)

	var x Dipen
	x.First = "hahaha"
	log.Println(x.newstruct())

	user := User{
		FirstName: "dipen",
		LastName:  "thapa",
		Location:  "lamachour",
		Phone:     "1222",
	}
	log.Println(user.BirthDate)

}
func dipen(v string) string {
	v = "dipndra"
	return v

}

return u.First
func (u *Dipen) newstruct() string {
	// function using receiver
}
