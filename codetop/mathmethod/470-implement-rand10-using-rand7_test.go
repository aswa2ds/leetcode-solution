package mathmethod

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_rand10(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	m := make(map[int]int)

	for i := 0; i < 11000; i++ {
		m[rand11()]++
	}

	fmt.Println(m)
}
