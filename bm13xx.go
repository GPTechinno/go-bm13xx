package bm13xx

import (
	"fmt"
	"io"
	"time"
)

type Asic struct {
	Regs     map[RegAddr]uint32
	CoreRegs map[CoreRegID]uint16
}

func (a Asic) Addr() byte {
	if chipAddress, exist := a.Regs[ChipAddress]; exist {
		return byte(chipAddress & 0xff)
	}
	return 0
}

func (a Asic) CoreNum() byte {
	if chipAddress, exist := a.Regs[ChipAddress]; exist {
		return byte((chipAddress >> 8) & 0xff)
	}
	return 0
}

func (a Asic) PllFreq(pll int, clki uint32) (uint32, error) {
	pllParams := []RegAddr{PLL0Parameter, PLL1Parameter, PLL2Parameter, PLL3Parameter}
	if pll >= len(pllParams) || pll < 0 {
		return 0, fmt.Errorf("pll %d out of range", pll)
	}
	if pllParam, exist := a.Regs[pllParams[pll]]; exist {
		pllLocked := (pllParam >> 31) & 0x01
		pllEn := (pllParam >> 30) & 0x01
		if pllLocked == 0 || pllEn == 0 {
			return 0, nil
		}
		fbdiv := (pllParam >> 16) & 0xfff
		refdiv := (pllParam >> 8) & 0x3f
		postdiv1 := (pllParam >> 4) & 0x07
		postdiv2 := pllParam & 0x07
		divide := refdiv * postdiv1 * postdiv2
		if divide == 0 {
			return 0, fmt.Errorf("zero divider")
		}
		return uint32(clki * fbdiv / divide), nil
	}
	return 0, fmt.Errorf("PLL%dParameter not found", pll)
}

type Chain struct {
	port   io.ReadWriter
	is139x bool
	clk    uint32
	Asics  []Asic
}

func NewChain(port io.ReadWriter, is139x bool, clk uint32) *Chain {
	c := &Chain{port: port, is139x: is139x, clk: clk}
	return c
}

func (c *Chain) chipIndex(chipAddr byte) (int, error) {
	for i, a := range c.Asics {
		if a.Addr() == chipAddr {
			return i, nil
		}
	}
	return 0, fmt.Errorf("not found")
}

func (c *Chain) Init(increment byte) error {
	if increment == 0 {
		return fmt.Errorf("increment must be greater than 0")
	}
	if len(c.Asics) > 0 {
		return fmt.Errorf("already enumerated")
	}
	// Enumerate the chips
	c.ReadRegister(true, 0, ChipAddress)
	for {
		regVal, chipAddr, regAddr, err := c.GetResponse()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if regAddr != byte(ChipAddress) {
			return fmt.Errorf("bad regAddr")
		}
		if chipAddr != 0x00 {
			return fmt.Errorf("bad chipAddr")
		}
		a := Asic{}
		a.Regs = make(map[RegAddr]uint32)
		a.Regs[ChipAddress] = regVal
		a.CoreRegs = make(map[CoreRegID]uint16)
		c.Asics = append(c.Asics, a)
	}
	// ChainInactive 3 times
	for i := 0; i < 3; i++ {
		c.Inactive()
		time.Sleep(30 * time.Millisecond)
	}
	// Gives new ChipAddresses
	if len(c.Asics)-1*int(increment) > 255 {
		return fmt.Errorf("too many chips or too big increment")
	}
	newChipAddr := byte(0)
	for i := range c.Asics {
		err := c.SetChipAddr(newChipAddr)
		if err != nil {
			return err
		}
		c.Asics[i].Regs[ChipAddress] += uint32(newChipAddr)
		newChipAddr += increment
		time.Sleep(30 * time.Millisecond)
	}
	time.Sleep(120 * time.Millisecond)
	// Init Ordered Clock
	c.WriteRegister(true, 0, ClockOrderControl0, 0)
	c.WriteRegister(true, 0, ClockOrderControl1, 0)
	time.Sleep(100 * time.Millisecond)
	c.WriteRegister(true, 0, OrderedClockEnable, 0)
	time.Sleep(100 * time.Millisecond)
	c.WriteRegister(true, 0, OrderedClockEnable, 0xFF)
	time.Sleep(10 * time.Millisecond)
	c.WriteRegister(true, 0, CoreRegisterControl, 0x800080B4)
	time.Sleep(5 * time.Millisecond)
	c.WriteRegister(true, 0, TicketMask, 0xFC)
	time.Sleep(10 * time.Millisecond)
	// c.WriteRegister(true, 0, MiscControl, 0x1A01)
	c.WriteRegister(true, 0, MiscControl, 0x2001)
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (c *Chain) ReadAllRegisters(chipIndex int) error {
	if chipIndex >= len(c.Asics) {
		return fmt.Errorf("chipIndex %d out of range", chipIndex)
	}
	regs := allRegisters
	for _, reg := range regs {
		c.ReadRegister(false, c.Asics[chipIndex].Addr(), reg)
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
		c.Asics[chipIndex].Regs[reg] = regVal
	}
	return nil
}

func (c *Chain) ReadUnknownRegisters(chipIndex int) error {
	if chipIndex >= len(c.Asics) {
		return fmt.Errorf("chipIndex %d out of range", chipIndex)
	}
	regs := []RegAddr{0x24, 0x30, 0x34, 0x88}
	for _, reg := range regs {
		c.ReadRegister(false, c.Asics[chipIndex].Addr(), reg)
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
		c.Asics[chipIndex].Regs[reg] = regVal
	}
	return nil
}

func (c *Chain) ReadCoreRegister(chipAddr byte, coreID uint16, coreRegID CoreRegID) (uint16, error) {
	chipIndex, err := c.chipIndex(chipAddr)
	if err != nil {
		return 0, err
	}
	if coreID >= uint16(c.Asics[chipIndex].CoreNum()) {
		return 0, fmt.Errorf("coreID %d out of range", coreID)
	}
	// coreRegCtrlVal := uint32(0x7e003000)
	coreRegCtrlVal := uint32(0x000000ff)
	coreRegCtrlVal |= uint32(coreRegID) << 8
	coreRegCtrlVal |= uint32(coreID) << 16
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
	if coreID >= uint16(c.Asics[chipIndex].CoreNum()) {
		return fmt.Errorf("coreID %d out of range", coreID)
	}
	regs := allCoreRegisters
	for _, reg := range regs {
		val, err := c.ReadCoreRegister(chipAddr, coreID, reg)
		if err != nil {
			fmt.Printf("ReadCoreRegister on reg %v error: %v\n", reg, err)
			continue
		}
		c.Asics[chipIndex].CoreRegs[reg] = val
	}
	return nil
}

func (c *Chain) SetBaudrate(baud uint32) error {
	if baud < 115200 || baud > 7000000 {
		return fmt.Errorf("bad baudrate, accepted range [115200:7000000]")
	}
	if len(c.Asics) == 0 {
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
	miscCtrl, exist := c.Asics[0].Regs[MiscControl]
	if !exist {
		// Use first asic in chain to read MiscControl register value
		c.ReadRegister(false, 0, MiscControl)
		val, _, reg, err := c.GetResponse()
		if err != nil {
			return err
		}
		c.Asics[0].Regs[RegAddr(reg)] = val
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
	// Apply the new baudrate settings to all Asics in chain
	c.WriteRegister(true, 0, MiscControl, miscCtrl)
	return nil
}

func (c *Chain) DumpChipRegiters(chipIndex int, debug bool) error {
	if chipIndex >= len(c.Asics) || chipIndex < 0 {
		return fmt.Errorf("bad chipIndex")
	}
	for _, addr := range allRegisters {
		if val, exist := c.Asics[chipIndex].Regs[addr]; exist {
			DumpAsicReg(addr, val, debug)
		}
	}
	for _, id := range allCoreRegisters {
		if val, exist := c.Asics[chipIndex].CoreRegs[id]; exist {
			DumpCoreReg(id, val, debug)
		}
	}
	return nil
}
