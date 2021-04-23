package main

import "log"

type Vehicle interface {
	Says() string
	Havewhills() int
}
type Bike struct {
	Owner          string
	Bikename       string
	Combustionrate int
}
type Car struct {
	Owner          string
	Carname        string
	Combustionrate int
}

func main() {
	b := Bike{
		Owner:          "dipen",
		Bikename:       "pulsor",
		Combustionrate: 11,
	}
	c := Car{
		Owner:          "dipenra",
		Carname:        "Ford,Ferrari,Tesla",
		Combustionrate: 10,
	}
	log.Println(b.Says(), b.Havewhills(), b)
	log.Println(c.Says(), c.Havewhills(), c)
	Printinfo(b)
	Printinfo(c)

}
func (b Bike) Says() string {

	return "bike sound"
}
func (b Bike) Havewhills() int {

	return 2
}

func (c Car) Says() string {

	return "car sound"
}
func (c Car) Havewhills() int {

	return 4
}

func Printinfo(a Vehicle) {
	log.Println("Vehicle says", a.Says(), "Have whills", a.Havewhills(), a)
}
