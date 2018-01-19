package ratecounter

import (
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

/*
RateCounter counts the number of event recorded by a counter per second. It is not really goroutine safe but
it's good enough for this purpose.
*/
type RateCounter struct {
	ticker   *time.Ticker
	counter  *int64
	lastRead int64
}

/*
NewRateCounter is a factory method for RateCounter
*/
func NewRateCounter(name string, counter *int64) *RateCounter {

	rc := new(RateCounter)
	rc.counter = counter
	rc.ticker = time.NewTicker(1 * time.Second)

	printer := message.NewPrinter(language.English)

	go func() {
		for _ = range rc.ticker.C {
			rate := *counter - rc.lastRead
			rc.lastRead = *counter
			printer.Printf("%s - rate: %d events/second\n", name, rate)
		}
	}()

	return rc
}
