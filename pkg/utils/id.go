package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// Generate uniq id
func IdGen() string {
	s1 := rand.NewSource(time.Now().Unix())
	r1 := rand.New(s1)
	name := strconv.Itoa(r1.Intn(100000))

	return name
}
