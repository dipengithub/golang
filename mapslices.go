package main

// slice means array and map means key value store nimmutable very fast
import "log"

type Dipentype struct {
	Name   string
	Height float64
}

func main() {
	log.Println("h")
	m := make(map[string]string)
	mn := make(map[float32]int)
	mno := make(map[string]Dipentype)
	x := Dipentype{Name: "dipen", Height: 12.2}
	m["name"] = "dipendra thapa"
	m["spatialinfo"] = "latang"
	log.Println(m["name"], m["spatialinfo"])
	mn[1.1] = 1
	mn[5.5] = 6
	log.Println(mn, mn)
	mno["info"] = x
	log.Println(mno["info"].Height)

	// part of slice
	s := []string{"dip", "en", "dra", "th", "ap", "a"}
	var sa []string
	sa = append(sa, "sa")
	sa = append(sa, "ga")
	sa = append(sa, "sa")
	sa = append(sa, "sa")
	sa = append(sa, "sas")
	sa = append(sa, "sac")
	sa = append(sa, "sab")
	sa = append(sa, "san")
	sa = append(sa, "sah")
	sa = append(sa, "saj")
	sa = append(sa, "sau")

	log.Println(slices(s))
	log.Println(slices(sa))

}
func slices(s []string) (string, []string) {
	return s[2], s[1:4]
}
