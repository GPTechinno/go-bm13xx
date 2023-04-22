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
	clk    uint32
	asics  []Asic
}

func NewChain(port io.ReadWriter, is139x bool, clk uint32) *Chain {
	c := &Chain{port: port, is139x: is139x, clk: clk}
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

func (c *Chain) SetBaudrate(baud uint32) error {
	if baud < 115200 || baud > 7000000 {
		return fmt.Errorf("bad baudrate, accepted range [115200:7000000]")
	}
	if len(c.asics) == 0 {
		return fmt.Errorf("no asic found")
	}
	baseClk := c.clk
	if baud > 3000000 {
		// LOCKED | PLLEN | FBDIV = 112 | REFDIV = 1 | POSTDIV1 = 1 | POSTDIV2 = 1
		// PLL3 = 25MHz * 112 = 2.8GHz
		c.WriteRegister(true, 0, PLL3Parameter, 0xc0700111)
		c.WriteRegister(true, 0, PLL3Parameter, 0xc0700111)
		// PLL3_DIV4 = 6 | CLKO_DIV = 15
		// uart baseClk is PLL3 / (DIV4 + 1) = 2.8GHz / (6 + 1) = 400MHz
		c.WriteRegister(true, 0, FastUARTConfiguration, 0x600000f)
		baseClk = 400000000
	}
	// TODO : calculate divider based on baseClk and baud
	divider := uint32(baseClk / baud)
	bt8d_4_0 := divider & 0x1f
	bt8d_8_5 := (divider >> 5) & 0x0f
	miscCtrl, exist := c.asics[0].regs[MiscControl]
	if !exist {
		// Use first asic in chain to read MiscControl register value
		c.ReadRegister(false, 0, MiscControl)
		val, _, reg, err := c.GetResponse()
		if err != nil {
			return err
		}
		c.asics[0].regs[RegAddr(reg)] = val
		if reg == byte(MiscControl) {
			miscCtrl = val
		} else {
			return fmt.Errorf("unexpected register")
		}
	}
	miscCtrl = miscCtrl&0xf0fee0ff | bt8d_4_0<<8 | bt8d_8_5<<24
	if baud > 3000000 {
		miscCtrl |= 1 << 16 // BCLK_SEL = 1
	}
	// Apply the new baudrate settings to all asics in chain
	c.WriteRegister(true, 0, MiscControl, miscCtrl)
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
