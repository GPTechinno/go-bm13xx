package main

import (
	"flag"
	"log"
	"time"

	"github.com/GPTechinno/go-bm13xx"
	"github.com/tarm/serial"
)

func main() {
	pCom := flag.String("c", "/dev/serial/by-id/usb-FTDI_TTL232RG-VREG1V8_FT62FVAA-if00-port0", "COM Port")
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
	chain.Enumerate()
	// chain.Inactive()
	time.Sleep(time.Second)
	chain.ReadAllRegisters(0)
	chain.DumpChipRegiters(0, false)
}
