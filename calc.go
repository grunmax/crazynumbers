package main

import (
	"math"
)

type Ijk struct {
	i int
	j int
	k int
}

func bingo(data Ijk) (bool, int64) {
	act := concat(data)
	exp := int64(data.i*data.i*data.i) + int64(data.j*data.j*data.j) + int64(data.k*data.k*data.k)
	return exp == act, act
}

func pow(a int, n int) int64 {
	if a == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	res := int64(a)
	for i := 1; i <= n-1; i++ {
		res = res * int64(a)
	}
	return res
}

func width(i int) int {
	if i < 10 {
		return 1
	}
	if i < 100 {
		return 2
	}
	if i < 1000 {
		return 3
	}
	return int(math.Log10(float64(i))) + 1
}

func concat(data Ijk) int64 {
	widthj := width(data.j)
	widthk := width(data.k)
	parti := int64(data.i) * pow(10, widthj+widthk)
	partj := int64(data.j) * pow(10, widthk)
	return parti + partj + int64(data.k)
}
