package main

import (
	"chain"
)

func main() {

	forever := make(chan bool)

	preLoad := 100000

	pipeIn := make(chan chain.Payload, preLoad)
	pipeOut := make(chan chain.Payload, preLoad)
	pipeMiddle := make(chan chain.Payload, preLoad)
	//pipeMiddle2 := make(chan chain.Payload, preLoad)

	chain.Load(pipeIn, preLoad)

	go chain.Sink(pipeOut)

	go chain.Process1Ms(pipeIn, pipeMiddle)

	//go chain.Process01Unit(pipeMiddle, pipeMiddle2)
	//go chain.Process01UnitRnd(pipeMiddle, pipeMiddle2)

	go chain.Process10MsRnd(pipeMiddle, pipeOut)
	go chain.Process10MsRnd(pipeMiddle, pipeOut)
	go chain.Process10MsRnd(pipeMiddle, pipeOut)
	go chain.Process10MsRnd(pipeMiddle, pipeOut)

	<-forever
}
