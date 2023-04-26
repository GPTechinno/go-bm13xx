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
	err = chain.Init(8)
	if err != nil {
		log.Fatalln(err)
	}
	time.Sleep(time.Second)
	p.SetSpeed(3000000) // 3.125 MBps on the line
	time.Sleep(200 * time.Millisecond)
	chain.ReadRegister(true, 0, bm13xx.CoreRegisterControl)
	time.Sleep(200 * time.Millisecond)
	chain.ReadRegister(false, 0, bm13xx.CoreRegisterControl)
	time.Sleep(200 * time.Millisecond)
	chain.WriteRegister(false, 0, bm13xx.CoreRegisterControl, 0x80008501)
	time.Sleep(time.Millisecond)
	chain.WriteRegister(false, 0, bm13xx.CoreRegisterControl, 0x6FF)
	time.Sleep(2 * time.Millisecond)
	chain.WriteRegister(false, 0, bm13xx.CoreRegisterControl, 0x80008500)
	// chain.ReadAllRegisters(0)
	// chain.ReadAllCoreRegisters(0, 0)
	// chain.ReadUnknownRegisters(0)
	// fmt.Println("INITIAL REGISTERS MAP")
	// chain.DumpChipRegiters(0, false)
	time.Sleep(time.Second)
	// // test Resgiter Writability
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
