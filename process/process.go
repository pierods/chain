package main

import (
	"chain"
)

func main() {

	forever := make(chan bool)

	preLoad := 10000

	pipeIn := make(chan chain.Payload, preLoad)
	pipeOut := make(chan chain.Payload, preLoad)
	pipeMiddle := make(chan chain.Payload, preLoad)
	pipeMiddle2 := make(chan chain.Payload, preLoad)

	chain.Load(pipeIn, preLoad)

	go chain.Sink(pipeOut)

	go chain.Process1000Ms(pipeIn, pipeMiddle)

	go chain.Process1000Ms(pipeMiddle, pipeMiddle2)

	go chain.Process100Ms(pipeMiddle2, pipeOut)

	<-forever
}
