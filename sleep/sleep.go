package main

import "github.com/pierods/chain"

func main() {

	forever := make(chan bool)

	preLoad := 10000

	pipeIn := make(chan chain.Payload, preLoad)
	pipeOut := make(chan chain.Payload, preLoad)
	pipeMiddle := make(chan chain.Payload, preLoad)

	chain.Load(pipeIn, preLoad)

	go chain.Sink(pipeOut)

	go chain.Sleep1Unit(pipeIn, pipeMiddle)

	go chain.Sleep10Units(pipeMiddle, pipeOut)
	go chain.Sleep10Units(pipeMiddle, pipeOut)
	go chain.Sleep10Units(pipeMiddle, pipeOut)
	go chain.Sleep10Units(pipeMiddle, pipeOut)
	go chain.Sleep10Units(pipeMiddle, pipeOut)
	go chain.Sleep10Units(pipeMiddle, pipeOut)
	go chain.Sleep10Units(pipeMiddle, pipeOut)
	go chain.Sleep10Units(pipeMiddle, pipeOut)
	go chain.Sleep10Units(pipeMiddle, pipeOut)
	go chain.Sleep10Units(pipeMiddle, pipeOut)

	<-forever
}
