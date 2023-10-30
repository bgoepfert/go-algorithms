package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func gcd(a float64, b float64) float64 {
	q := math.Floor(a / b)
	r := a - (b * q)

	if int64(a)%int64(b) == 0 {
		return b
	}

	return gcd(b, r)
}

func main() {
	args := os.Args[1:]

	a, errA := strconv.ParseFloat(args[0], 64)
	if errA != nil {
		panic(errA)
	}

	b, errB := strconv.ParseFloat(args[1], 64)
	if errB != nil {
		panic(errB)
	}

	gcd := gcd(a, b)

	fmt.Printf("gcd(%v,%v) = %g\n", args[0], args[1], gcd)
}
