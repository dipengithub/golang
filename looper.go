package main

import "log"

type Mytype struct {
	Name    string
	Country string
}

func main() {
	s := []string{"d", "i", "p", "e", "n"}
	m := make(map[string]string)
	m["name"] = "dipen"
	m["age"] = "12"

	t1 := Mytype{
		Name:    "dipendra",
		Country: "Nepal",
	}
	t2 := Mytype{
		Name:    "yankar",
		Country: "Pakistan",
	}

	var t []Mytype
	t = append(t, t1)
	t = append(t, t2)

	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
	for _, v := range s {
		log.Println(v)

	}
	for i, v := range m {
		log.Println(i, v)

	}
	for i, v := range t {
		log.Println(i, v, v.Name, v.Country)

	}
}
