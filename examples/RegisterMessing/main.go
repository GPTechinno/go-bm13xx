package main

import (
	"flag"
	"log"
	"time"

	"github.com/GPTechinno/go-bm13xx"
	"github.com/davecgh/go-spew/spew"
	"github.com/tarm/serial"
)

func main() {
	pCom := flag.String("c", "COM6", "COM Port")
	pBaud := flag.Int("b", 115200, "Baudrate")
	flag.Parse()
	// Open COM Port
	c := &serial.Config{
		Name:        *pCom,
		Baud:        *pBaud,
		Parity:      serial.ParityNone,
		ReadTimeout: 100 * time.Millisecond,
	}
	p, err := serial.OpenPort(c)
	if err != nil {
		log.Fatalln(err)
	}
	chain := bm13xx.NewChain(p, true)
	chain.ReadRegister(true, 0, bm13xx.ChipAddress)
	for {
		regVal, chipAddr, regAddr, err := chain.GetResponse()
		if err != nil {
			spew.Dump(err)
			break
		}
		spew.Dump(regVal, regAddr, chipAddr)
	}
}
