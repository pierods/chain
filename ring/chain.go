package ring

import (
	"github.com/Workiva/go-datastructures/queue"
	"github.com/pierods/chain/ratecounter"
	"time"
)

type Payload struct{}

/*
Sleep1SmallUnit will select from its input channel, sleep 1msec (1 small unit of time, 1000ops/sec theoretical throughput) and write to its output channel
*/
func Sleep1SmallUnit(ringIn, ringOut *queue.RingBuffer) {

	for {
		_, err := ringIn.Get()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond)
		err = ringOut.Put(Payload{})
		if err != nil {
			panic(err)
		}
	}
}

func Sleep10SmallUnits(ringIn, ringOut *queue.RingBuffer) {
	for {
		_, err := ringIn.Get()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond * 10)
		err = ringOut.Put(Payload{})
		if err != nil {
			panic(err)
		}
	}
}

func Load(ring *queue.RingBuffer, size int) {
	payLoad := Payload{}

	for i := 0; i < size; i++ {
		err := ring.Put(payLoad)
		if err != nil {
			panic(err)
		}
	}
}

func Sink(sink *queue.RingBuffer) {

	var counter int64
	ratecounter.NewRateCounter("sink", &counter)

	for {
		_, err := sink.Get()
		if err != nil {
			panic(err)
		}
		counter++
	}
}
