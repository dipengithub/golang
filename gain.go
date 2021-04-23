package main

import "log"

func main() {
	var d string
	d = "dipen"

	log.Println(&d)
	log.Println(d)
	dipen(&d)
	log.Println(d)

}
func dipen(s *string) {
	g := "thapa"
	log.Println(s)
	log.Println(*s)
	log.Println(&s)
	log.Println(s)

	*s = g

}
