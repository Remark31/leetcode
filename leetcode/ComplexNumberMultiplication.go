package leetcode

import (
	"strconv"
	"strings"
)

func stringTocomplexNumber(a string) (int, int) {
	al := strings.Split(a, "+")
	ar, err := strconv.Atoi(al[0])
	if err != nil {
		panic("string int is error")
	}

	ai, err := strconv.Atoi(al[1][:len(al[1])-1])
	if err != nil {
		panic("string complex is error")
	}

	return ar, ai
}

func complexNumberToString(ar, ai int) string {
	ars := strconv.Itoa(ar)
	ais := strconv.Itoa(ai)

	return ars + "+" + ais + "i"
}

func complexNumberMultiply(a string, b string) string {
	ar, ai := stringTocomplexNumber(a)
	br, bi := stringTocomplexNumber(b)

	ansr := ar*br - (ai * bi)
	ansi := ar*bi + br*ai

	return complexNumberToString(ansr, ansi)
}
