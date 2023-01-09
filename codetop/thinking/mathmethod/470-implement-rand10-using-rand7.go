package mathmethod

import (
	"math/rand"
)

func rand7() int {
	return rand.Intn(7) + 1
}

func rand10() int {
	rand5, rand2 := 8, 8
	for ; rand5 > 5; rand5 = rand7() {
	}
	for ; rand2 > 6; rand2 = rand7() {
	}

	rand2 = rand2 & 1

	return rand5 + rand2*5
}

func rand11() int {
	rand11 := 12
	for rand11 > 11 {
		rand2 := 8
		for ; rand2 > 6; rand2 = rand7() {

		}
		rand2 = rand2 & 1
		rand11 = rand7() + rand2*7
	}
	return rand11
}
