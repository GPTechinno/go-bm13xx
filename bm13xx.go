package bm13xx

import (
	"fmt"
	"io"
)

type Asic struct {
	regs map[RegAddr]uint32
}

type Chain struct {
	port   io.ReadWriter
	is139x bool
	asics  []Asic
}

func NewChain(port io.ReadWriter, is139x bool) *Chain {
	c := &Chain{port: port, is139x: is139x}
	return c
}

func (c *Chain) Enumerate() error {
	c.ReadRegister(true, 0, ChipAddress)
	for {
		regVal, chipAddr, regAddr, err := c.GetResponse()
		if err != nil {
			return err
		}
		if regAddr != byte(ChipAddress) {
			return fmt.Errorf("bad regAddr")
		}
		if chipAddr != 0x00 {
			return fmt.Errorf("bad chipAddr")
		}
		a := Asic{}
		a.regs = make(map[RegAddr]uint32)
		a.regs[ChipAddress] = regVal
		c.asics = append(c.asics, a)
	}
}

func (c *Chain) DumpChipRegiters(chipIndex int, debug bool) error {
	if chipIndex >= len(c.asics) || chipIndex < 0 {
		return fmt.Errorf("bad chipIndex")
	}
	for addr, val := range c.asics[chipIndex].regs {
		regDump(addr, val, debug)
	}
	return nil
}
