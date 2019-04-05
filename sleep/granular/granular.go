package main

import (
	"github.com/pierods/chain"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {

	runtime.SetMutexProfileFraction(20)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	forever := make(chan bool)

	preLoad := 10000

	pipeIn := make(chan chain.Payload, preLoad)
	pipeOut := make(chan chain.Payload, preLoad)
	pipeMiddle := make(chan chain.Payload, preLoad)

	chain.Load(pipeIn, preLoad)

	go chain.Sink(pipeOut)

	go chain.Sleep1SmallUnit(pipeIn, pipeMiddle)

	go chain.Sleep10SmallUnits(pipeMiddle, pipeOut)
	go chain.Sleep10SmallUnits(pipeMiddle, pipeOut)
	go chain.Sleep10SmallUnits(pipeMiddle, pipeOut)
	go chain.Sleep10SmallUnits(pipeMiddle, pipeOut)
	go chain.Sleep10SmallUnits(pipeMiddle, pipeOut)
	go chain.Sleep10SmallUnits(pipeMiddle, pipeOut)
	go chain.Sleep10SmallUnits(pipeMiddle, pipeOut)
	go chain.Sleep10SmallUnits(pipeMiddle, pipeOut)
	go chain.Sleep10SmallUnits(pipeMiddle, pipeOut)
	go chain.Sleep10SmallUnits(pipeMiddle, pipeOut)

	<-forever
}
