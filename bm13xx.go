package bm13xx

import (
	"fmt"
	"io"
)

type Asic struct {
	regs     map[RegAddr]uint32
	coreRegs map[CoreRegID]uint16
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
		a.coreRegs = make(map[CoreRegID]uint16)
		c.asics = append(c.asics, a)
	}
}

func (c *Chain) ReadAllRegisters(chipIndex int) error {
	regs := allRegisters
	for _, reg := range regs {
		c.ReadRegister(false, 0, reg)
		regVal, chipAddr, regAddr, err := c.GetResponse()
		if err != nil {
			fmt.Printf("GetResponse error: %v\n", err)
			return err
		}
		if regAddr != byte(reg) {
			return fmt.Errorf("bad regAddr")
		}
		if chipAddr != 0x00 {
			return fmt.Errorf("bad chipAddr")
		}
		c.asics[chipIndex].regs[reg] = regVal
	}
	return nil
}

func (c *Chain) ReadUnknownRegisters(chipIndex int) error {
	regs := []RegAddr{0x24, 0x30, 0x34, 0x88}
	for _, reg := range regs {
		c.ReadRegister(false, 0, reg)
		regVal, chipAddr, regAddr, err := c.GetResponse()
		if err != nil {
			fmt.Printf("GetResponse error: %v\n", err)
			continue
		}
		if regAddr != byte(reg) {
			fmt.Printf("bad regAddr: %v\n", regAddr)
			continue
		}
		if chipAddr != 0x00 {
			fmt.Printf("bad chipAddr: %v\n", chipAddr)
			continue
		}
		c.asics[chipIndex].regs[reg] = regVal
	}
	return nil
}

func (c *Chain) ReadCoreRegister() {

}

func (c *Chain) DumpChipRegiters(chipIndex int, debug bool) error {
	if chipIndex >= len(c.asics) || chipIndex < 0 {
		return fmt.Errorf("bad chipIndex")
	}
	for _, addr := range allRegisters {
		if val, exist := c.asics[chipIndex].regs[addr]; exist {
			dumpAsicReg(addr, val, debug)
		}
	}
	return nil
}
