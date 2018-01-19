package primes

import (
	"testing"
)

func TestIsPrime(t *testing.T) {

	if !IsPrime(uint64(3)) {
		t.Fatal("Should recognize a prime number")
	}

	if !IsPrime(uint64(5)) {
		t.Fatal("Should recognize a prime number")
	}

	if IsPrime(uint64(4)) {
		t.Fatal("Should recognize a prime number")
	}
}

func BenchmarkPrimes(b *testing.B) {

	//number := uint64(2e4) // 100ms
	number := uint64(2e3)

	for i := 0; i < b.N; i++ {
		AllPrimes(number)
	}
}
