package chain

import (
	"chain/primes"
	"math/rand"
	"time"

	"chain/ratecounter"
)

type Payload struct{}

/*
Sleep1Unit will select from its input channel, sleep 100Msec (1 unit of time) and write to its output channel
*/
func Sleep1Unit(pipeIn, pipeOut chan Payload) {
	payLoad := struct{}{}
	for {
		select {
		case _ = <-pipeIn:
			time.Sleep(time.Millisecond * 100)
			pipeOut <- payLoad
		}
	}
}

/*
Sleep3Units will select from its input channel, sleep 300Msec (3 units of time) and write to its output channel
*/
func Sleep3Units(pipeIn, pipeOut chan Payload) {
	payLoad := struct{}{}
	for {
		select {
		case _ = <-pipeIn:
			time.Sleep(time.Millisecond * 100 * 3)
			pipeOut <- payLoad
		}
	}
}

/*
Sleep10Units will select from its input channel, sleep 1000Msec (10 units of time) and write to its output channel
*/
func Sleep10Units(pipeIn, pipeOut chan Payload) {
	payLoad := struct{}{}
	for {
		select {
		case _ = <-pipeIn:
			time.Sleep(time.Millisecond * 100 * 10)
			pipeOut <- payLoad
		}
	}
}

/*
Load preloads a channel with size values
*/
func Load(pipe chan Payload, size int) {
	payLoad := struct{}{}

	for i := 0; i < size; i++ {
		pipe <- payLoad
	}
}

/*
Sink acts as a sink for messages and counts the message rate
*/
func Sink(pipe chan Payload) {

	var counter int64
	ratecounter.NewRateCounter("sink", &counter)

	for {
		select {
		case _ = <-pipe:
			counter++
		}
	}
}

/*
Process1MsRnd will read a value from its input channel, wait a random time lesser than 1Ms, keep the CPU busy for 1Ms and then write to its output channel
*/
func Process1MsRnd(pipeIn, pipeOut chan Payload) {
	payLoad := struct{}{}
	number := uint64(2e3) // 0.001secs

	for {
		select {
		case _ = <-pipeIn:
			wait := rand.Intn(1000)
			time.Sleep(time.Microsecond * time.Duration(wait))
			primes.AllPrimes(number)
			pipeOut <- payLoad
		}
	}
}

/*
Process1Ms will read a value from its input channel, keep the CPU busy for 1Ms and then write to its output channel
*/
func Process1Ms(pipeIn, pipeOut chan Payload) {
	payLoad := struct{}{}
	number := uint64(2e3)

	for {
		select {
		case _ = <-pipeIn:
			primes.AllPrimes(number)
			pipeOut <- payLoad
		}
	}
}

/*
Process10Ms will read a value from its input channel, keep the CPU busy for 10Ms and then write to its output channel
*/
func Process10Ms(pipeIn, pipeOut chan Payload) {
	payLoad := struct{}{}
	number := uint64(6e3)

	for {
		select {
		case _ = <-pipeIn:
			primes.AllPrimes(number)
			pipeOut <- payLoad
		}
	}
}

/*
Process10MsRnd will read a value from its input channel, wait a random time lesser than 10Ms, keep the CPU busy for 10Ms and then write to its output channel
*/
func Process10MsRnd(pipeIn, pipeOut chan Payload) {
	payLoad := struct{}{}
	number := uint64(6e3)

	for {
		select {
		case _ = <-pipeIn:
			wait := rand.Intn(10)
			time.Sleep(time.Millisecond * time.Duration(wait))
			primes.AllPrimes(number)
			pipeOut <- payLoad
		}
	}
}

/*
Process100Ms will read a value from its input channel, keep the CPU busy for 100Ms and then write to its output channel
*/
func Process100Ms(pipeIn, pipeOut chan Payload) {
	payLoad := struct{}{}
	number := uint64(2e4)

	for {
		select {
		case _ = <-pipeIn:
			primes.AllPrimes(number)
			pipeOut <- payLoad
		}
	}
}

/*
Process100MsRnd will read a value from its input channel, wait a random time lesser than 100Ms, keep the CPU busy for 100Ms and then write to its output channel
*/
func Process100MsRnd(pipeIn, pipeOut chan Payload) {
	payLoad := struct{}{}
	number := uint64(2e4)

	for {
		select {
		case _ = <-pipeIn:
			wait := rand.Intn(100)
			time.Sleep(time.Millisecond * time.Duration(wait))
			primes.AllPrimes(number)
			pipeOut <- payLoad
		}
	}
}

/*
Process1000Ms will read a value from its input channel, keep the CPU busy for 1000Ms and then write to its output channel
*/
func Process1000Ms(pipeIn, pipeOut chan Payload) {
	payLoad := struct{}{}
	number := uint64(7e4)

	for {
		select {
		case _ = <-pipeIn:
			for i := 0; i < 10; i++ {
				primes.AllPrimes(number)
			}
			pipeOut <- payLoad
		}
	}
}

/*
Process1000MsRnd will read a value from its input channel, wait a random time lesser than 1000Ms, keep the CPU busy for 1000Ms and then write to its output channel
*/
func Process1000MsRnd(pipeIn, pipeOut chan Payload) {
	payLoad := struct{}{}
	number := uint64(7e4)

	for {
		select {
		case _ = <-pipeIn:
			wait := rand.Intn(1000)
			time.Sleep(time.Millisecond * time.Duration(wait))

			for i := 0; i < 10; i++ {
				primes.AllPrimes(number)
			}
			pipeOut <- payLoad
		}
	}
}
