package bm13xx

import (
	"fmt"
	"io"
)

type Asic struct {
	regs     map[RegAddr]uint32
	coreRegs map[CoreRegID]uint16
}

func (a Asic) addr() byte {
	if chipAddress, exist := a.regs[ChipAddress]; exist {
		return byte(chipAddress & 0xff)
	}
	return 0
}

func (a Asic) coreNum() byte {
	if chipAddress, exist := a.regs[ChipAddress]; exist {
		return byte((chipAddress >> 8) & 0xff)
	}
	return 0
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

func (c *Chain) chipIndex(chipAddr byte) (int, error) {
	for i, a := range c.asics {
		if a.addr() == chipAddr {
			return i, nil
		}
	}
	return 0, fmt.Errorf("not found")
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
	if chipIndex >= len(c.asics) {
		return fmt.Errorf("chipIndex %d out of range", chipIndex)
	}
	regs := allRegisters
	for _, reg := range regs {
		c.ReadRegister(false, c.asics[chipIndex].addr(), reg)
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
	if chipIndex >= len(c.asics) {
		return fmt.Errorf("chipIndex %d out of range", chipIndex)
	}
	regs := []RegAddr{0x24, 0x30, 0x34, 0x88}
	for _, reg := range regs {
		c.ReadRegister(false, c.asics[chipIndex].addr(), reg)
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

func (c *Chain) ReadCoreRegister(chipAddr byte, coreID uint16, coreRegID CoreRegID) (uint16, error) {
	chipIndex, err := c.chipIndex(chipAddr)
	if err != nil {
		return 0, err
	}
	if coreID >= uint16(c.asics[chipIndex].coreNum()) {
		return 0, fmt.Errorf("coreID %d out of range", coreID)
	}
	coreRegCtrlVal := uint32(0x7e003000) // something in there must contain the coreID
	coreRegCtrlVal |= uint32(coreRegID) << 8
	err = c.WriteRegister(false, chipAddr, CoreRegisterControl, coreRegCtrlVal)
	if err != nil {
		return 0, err
	}
	coreRegVal, chip, reg, err := c.GetResponse()
	if err != nil {
		return 0, err
	}
	if chipAddr != chip {
		return 0, fmt.Errorf("bad chipAddr")
	}
	if byte(CoreRegisterValue) != reg {
		return 0, fmt.Errorf("bad regAddr")
	}
	if uint16(coreRegVal>>16) != coreID {
		return 0, fmt.Errorf("bad coreID")
	}
	return uint16(coreRegVal & 0xffff), nil
}

func (c *Chain) ReadAllCoreRegisters(chipAddr byte, coreID uint16) error {
	chipIndex, err := c.chipIndex(chipAddr)
	if err != nil {
		return err
	}
	if coreID >= uint16(c.asics[chipIndex].coreNum()) {
		return fmt.Errorf("coreID %d out of range", coreID)
	}
	regs := allCoreRegisters
	for _, reg := range regs {
		val, err := c.ReadCoreRegister(chipAddr, coreID, reg)
		if err != nil {
			fmt.Printf("ReadCoreRegister on reg %v error: %v\n", reg, err)
			continue
		}
		c.asics[chipIndex].coreRegs[reg] = val
	}
	return nil
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
	for _, id := range allCoreRegisters {
		if val, exist := c.asics[chipIndex].coreRegs[id]; exist {
			dumpCoreReg(id, val, debug)
		}
	}
	return nil
}
