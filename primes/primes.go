package primes

import (
	"math"
)

/*
AllPrimes keeps the CPU busy for a time exponentially proportional to value
*/
func AllPrimes(value uint64) {
	for j := uint64(0); j < value; j++ {
		result := IsPrime(j)
		result = !result
	}

}

/*
IsPrime isa naive implementation of IsPrime, chosen for its ability to keep the CPU busy
*/
func IsPrime(value uint64) bool {
	for i := uint64(2); i <= uint64(math.Floor(float64(value)/float64(2))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
