package main

import (
	"flag"
	"log"
	"time"

	"github.com/GPTechinno/go-bm13xx"
	"github.com/pkg/term"
)

func main() {
	pCom := flag.String("c", "/dev/serial/by-id/usb-FTDI_TTL232RG-VREG1V8_FT62FVAA-if00-port0", "COM Port")
	pBaud := flag.Int("b", 115200, "Baudrate")
	flag.Parse()
	p, err := term.Open(*pCom)
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Close()
	p.SetReadTimeout(100 * time.Millisecond)
	p.SetFlowControl(term.HARDWARE)
	// p.SetRaw()
	// try a Reset
	err = p.SetRTS(true)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(100 * time.Millisecond)
	p.SetRTS(false)
	time.Sleep(time.Second)
	p.SetRTS(true)
	time.Sleep(100 * time.Millisecond)
	p.SetSpeed(*pBaud)
	chain := bm13xx.NewChain(p, true, 25000000)
	baud, err := chain.Init(8)
	if err != nil {
		log.Fatalln(err)
	}
	time.Sleep(time.Second)
	p.SetSpeed(baud)
	time.Sleep(200 * time.Millisecond)

	// Chip Core messing
	// chain.ReadRegister(true, 0, bm13xx.CoreRegisterControl)
	// time.Sleep(200 * time.Millisecond)
	// chain.ReadRegister(false, 0, bm13xx.CoreRegisterControl)
	// time.Sleep(200 * time.Millisecond)
	// chain.WriteRegister(false, 0, bm13xx.CoreRegisterControl, 0x80008501)
	// time.Sleep(time.Millisecond)
	// chain.WriteRegister(false, 0, bm13xx.CoreRegisterControl, 0x6FF)
	// time.Sleep(2 * time.Millisecond)
	// chain.WriteRegister(false, 0, bm13xx.CoreRegisterControl, 0x80008500)

	// Mining Testing
	chain.WriteRegister(true, 0, bm13xx.Pll0Divider, 0x0F0F0F00)
	time.Sleep(10 * time.Millisecond)
	chain.WriteRegister(true, 0, bm13xx.Pll0Divider, 0x0F0F0F00)
	time.Sleep(10 * time.Millisecond)
	chain.WriteRegister(true, 0, bm13xx.PLL0Parameter, 0x40A00225) // 25MHz * 0x0A0 / 0x02 / 2 / 5 = 200MHz
	time.Sleep(10 * time.Millisecond)
	chain.WriteRegister(true, 0, bm13xx.PLL0Parameter, 0x40A00225)
	time.Sleep(10 * time.Millisecond)
	chain.ReadRegister(true, 0, bm13xx.PLL0Parameter)
	chain.GetResponse()
	time.Sleep(30 * time.Millisecond)
	// example job from Skot bm1397_protocol.md
	// chain.SendJob(4, 0, 0x170689A3, 0x640AA82F, 0xE8C51C6F, []bm13xx.Midstate{
	// 	{0x99, 0xCE, 0x15, 0xA2, 0xDE, 0xDF, 0x61, 0x49, 0x77, 0x90, 0x69, 0x84, 0x1C, 0x8A, 0x0C, 0xE0, 0xA0, 0x30, 0xE4, 0x0A, 0xCB, 0xA6, 0xA5, 0xFB, 0x33, 0x19, 0x30, 0x5D, 0xE2, 0x46, 0xEF, 0x18},
	// 	{0x1D, 0x63, 0x63, 0xEB, 0xE7, 0xBA, 0xD9, 0x8D, 0x4B, 0xCA, 0xFE, 0x9C, 0x4F, 0x45, 0xF6, 0x45, 0xFA, 0x71, 0xA0, 0x1E, 0x8C, 0xB8, 0x2D, 0x68, 0xDC, 0x6C, 0xB8, 0x4E, 0x25, 0x39, 0x8C, 0x50},
	// 	{0xFA, 0x7E, 0x2E, 0xC6, 0xC8, 0x08, 0x61, 0xB9, 0xA5, 0x89, 0x90, 0x71, 0xC4, 0x75, 0x56, 0xE4, 0x78, 0x85, 0x35, 0x22, 0x65, 0x51, 0xEA, 0x68, 0xEB, 0xF8, 0x96, 0xB0, 0xCA, 0x40, 0x77, 0xD4},
	// 	{0x0C, 0xAF, 0x1B, 0xD4, 0x47, 0x37, 0x85, 0xBB, 0x39, 0x6A, 0x22, 0xC3, 0x9C, 0x23, 0x56, 0xE7, 0xCE, 0xB6, 0x57, 0x4C, 0x1F, 0xA3, 0xA9, 0x9A, 0xD3, 0xC1, 0xA0, 0x17, 0x79, 0x1F, 0xBC, 0x38},
	// })
	chain.SendJob(48, 0, 0x17079E15, 0x638E3275, 0x995F3ED7, []bm13xx.Midstate{
		{0x03, 0x53, 0x4B, 0x27, 0xC1, 0xBD, 0xF5, 0x47, 0x07, 0xCA, 0xD9, 0x13, 0xB9, 0x69, 0x07, 0x01, 0x57, 0xC7, 0xFC, 0xDB, 0x48, 0xE3, 0xE0, 0xAB, 0x48, 0x7C, 0xE3, 0xA7, 0xDD, 0xFA, 0x2F, 0xA0},
		{0xED, 0x30, 0x1B, 0x59, 0x82, 0x15, 0xAB, 0xAC, 0x77, 0xEF, 0xEC, 0xD4, 0xF8, 0x3D, 0x95, 0x62, 0x1A, 0x5F, 0x4D, 0xCB, 0xB4, 0x18, 0x01, 0x88, 0xF3, 0x43, 0x30, 0xE0, 0xC9, 0xE2, 0xFD, 0x50},
		{0x55, 0x98, 0x13, 0xE0, 0x1E, 0x9C, 0x88, 0x28, 0x4E, 0xD3, 0x3E, 0xCE, 0x92, 0xA4, 0x82, 0xBA, 0x37, 0xA0, 0x47, 0x8A, 0x42, 0x87, 0xBE, 0xCD, 0xC5, 0xE1, 0x0A, 0x48, 0xBA, 0x4E, 0x41, 0x08},
		{0xB9, 0x2A, 0xA7, 0x52, 0x8F, 0xBE, 0x92, 0x59, 0x98, 0x2E, 0x59, 0x5B, 0xBB, 0x10, 0x28, 0xC7, 0x67, 0x89, 0x81, 0x49, 0x06, 0x51, 0xDA, 0x52, 0xA4, 0x59, 0xE6, 0x5C, 0x05, 0x77, 0xF5, 0xC4},
	})
	chain.ReadRegister(false, 0, bm13xx.HashRate)
	chain.GetResponse()
	chain.ReadRegister(false, 0, bm13xx.HashCountingNumber)
	chain.GetResponse()
	chain.ReadRegister(false, 0, bm13xx.HashRate)
	chain.GetResponse()
	chain.ReadRegister(false, 0, bm13xx.HashCountingNumber)
	chain.GetResponse()

	// Register Dump
	// chain.ReadAllRegisters(0)
	// chain.ReadAllCoreRegisters(0, 0)
	// chain.ReadUnknownRegisters(0)
	// fmt.Println("INITIAL REGISTERS MAP")
	// chain.DumpChipRegiters(0, false)
	time.Sleep(time.Second)

	// test Resgiter Writability
	// fmt.Println("INVERT ALL REGISTERS VALUES")
	// fmt.Println("-------------------------------------")
	// if len(chain.Asics) >= 1 {
	// 	for regAddr, regVal := range chain.Asics[0].Regs {
	// 		if regAddr == bm13xx.ClockOrderControl0 ||
	// 			regAddr == bm13xx.HashCountingNumber ||
	// 			regAddr == bm13xx.PLL0Parameter ||
	// 			regAddr == bm13xx.PLL1Parameter ||
	// 			regAddr == bm13xx.PLL2Parameter ||
	// 			regAddr == bm13xx.PLL3Parameter ||
	// 			regAddr == bm13xx.Pll0Divider ||
	// 			regAddr == bm13xx.Pll1Divider ||
	// 			regAddr == bm13xx.Pll2Divider ||
	// 			regAddr == bm13xx.Pll3Divider ||
	// 			regAddr == bm13xx.OrderedClockEnable ||
	// 			regAddr == bm13xx.AnalogMuxControl ||
	// 			regAddr == bm13xx.ChipNonceOffset ||
	// 			regAddr == bm13xx.CoreRegisterControl ||
	// 			regAddr == bm13xx.CoreRegisterValue ||
	// 			regAddr == bm13xx.OrderedClockMonitor ||
	// 			regAddr == bm13xx.ExternalTemperatureSensorRead ||
	// 			regAddr == bm13xx.ChipAddress ||
	// 			regAddr == bm13xx.NonceOverflowCounter ||
	// 			regAddr == bm13xx.ErrorFlag ||
	// 			regAddr == bm13xx.NonceErrorCounter ||
	// 			regAddr == bm13xx.TicketMask2 ||
	// 			regAddr == bm13xx.TicketMask ||
	// 			regAddr == bm13xx.GoldenNonceForSweepReturn ||
	// 			regAddr == bm13xx.SomeTempRelated ||
	// 			regAddr == bm13xx.ClockOrderStatus ||
	// 			regAddr == bm13xx.HashRate ||
	// 			regAddr == bm13xx.MiscControl ||
	// 			regAddr == bm13xx.ReturnedGroupPatternStatus ||
	// 			regAddr == bm13xx.ReturnedSinglePatternStatus ||
	// 			regAddr == bm13xx.IoDriverStrenghtConfiguration ||
	// 			regAddr == bm13xx.TimeOut ||
	// 			regAddr == bm13xx.FastUARTConfiguration ||
	// 			regAddr == bm13xx.NonceReturnedTimeout ||
	// 			regAddr == bm13xx.UARTRelay ||
	// 			regAddr == bm13xx.FrequencySweepControl1 ||
	// 			regAddr == bm13xx.ClockOrderControl1 {
	// 			continue
	// 		}
	// 		bm13xx.DumpAsicReg(regAddr, regVal, true)
	// 		inverter := uint32(0xffffffff)
	// 		if regAddr == bm13xx.MiscControl {
	// 			inverter = 0xffbfffff // avoid CORE_SRST
	// 		}
	// 		err := chain.WriteRegister(false, 0, regAddr, regVal^inverter)
	// 		fmt.Printf("WriteRegister %02X: %08X err: %v\n", regAddr, regVal^inverter, err)
	// 		err = chain.ReadRegister(true, 0, regAddr)
	// 		if err == nil {
	// 			regValue, chipAddr, regAddress, err := chain.GetResponse()
	// 			fmt.Printf("ReadRegister %02X %02X: %08X err: %v\n", chipAddr, regAddress, regValue, err)
	// 			bm13xx.DumpAsicReg(regAddr, regValue, true)
	// 		}
	// 		fmt.Println("-------------------------------------")
	// 	}
	// }

	// Test Chip Address
	// fmt.Println("CHIP ADDRESS")
	// fmt.Println("-------------------------------------")
	// if len(chain.Asics) >= 1 {
	// 	chain.SetChipAddr(0)
	// 	chain.SetChipAddr(4)
	// 	chain.ReadRegister(true, 0, bm13xx.ChipAddress)
	// 	regValue, chipAddr, regAddr, err := chain.GetResponse()
	// 	fmt.Printf("ReadRegister %02X %02X: %08X err: %v\n", chipAddr, regAddr, regValue, err)
	// 	bm13xx.DumpAsicReg(bm13xx.RegAddr(regAddr), regValue, true)
	// 	chain.ReadRegister(false, 4, bm13xx.ChipAddress)
	// 	regValue, chipAddr, regAddr, err = chain.GetResponse()
	// 	fmt.Printf("ReadRegister %02X %02X: %08X err: %v\n", chipAddr, regAddr, regValue, err)
	// 	bm13xx.DumpAsicReg(bm13xx.RegAddr(regAddr), regValue, true)
	// }
}
