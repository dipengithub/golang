package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int{
	rand.Seed(time.Now().UnixNano())
	r:=rand.Intn(n)
	return r
}