package bm13xx

import "fmt"

type RegAddr byte

const (
	ChipAddress                   RegAddr = 0x00
	HashRate                      RegAddr = 0x04
	PLL0Parameter                 RegAddr = 0x08
	ChipNonceOffset               RegAddr = 0x0C
	HashCountingNumber            RegAddr = 0x10
	TicketMask                    RegAddr = 0x14
	MiscControl                   RegAddr = 0x18
	OrderedClockEnable            RegAddr = 0x20
	FastUARTConfiguration         RegAddr = 0x28
	UARTRelay                     RegAddr = 0x2C
	TicketMask2                   RegAddr = 0x38
	ExternalTemperatureSensorRead RegAddr = 0x44
	ErrorFlag                     RegAddr = 0x48
	NonceErrorCounter             RegAddr = 0x4C
	NonceOverflowCounter          RegAddr = 0x50
	AnalogMuxControl              RegAddr = 0x54
	IoDriverStrenghtConfiguration RegAddr = 0x58
	TimeOut                       RegAddr = 0x5C
	PLL1Parameter                 RegAddr = 0x60
	PLL2Parameter                 RegAddr = 0x64
	PLL3Parameter                 RegAddr = 0x68
	OrderedClockMonitor           RegAddr = 0x6C
	Pll0Divider                   RegAddr = 0x70
	Pll1Divider                   RegAddr = 0x74
	Pll2Divider                   RegAddr = 0x78
	Pll3Divider                   RegAddr = 0x7C
	ClockOrderControl0            RegAddr = 0x80
	ClockOrderControl1            RegAddr = 0x84
	ClockOrderStatus              RegAddr = 0x8C
	FrequencySweepControl1        RegAddr = 0x90
	GoldenNonceForSweepReturn     RegAddr = 0x94
	ReturnedGroupPatternStatus    RegAddr = 0x98
	NonceReturnedTimeout          RegAddr = 0x9C
	ReturnedSinglePatternStatus   RegAddr = 0xA0
)

var allRegisters []RegAddr = []RegAddr{ChipAddress, HashRate, PLL0Parameter, ChipNonceOffset, HashCountingNumber,
	TicketMask, MiscControl, OrderedClockEnable, FastUARTConfiguration, UARTRelay, TicketMask2, ExternalTemperatureSensorRead,
	ErrorFlag, NonceErrorCounter, NonceOverflowCounter, AnalogMuxControl, IoDriverStrenghtConfiguration,
	TimeOut, PLL1Parameter, PLL2Parameter, PLL3Parameter, OrderedClockMonitor, Pll0Divider, Pll1Divider,
	Pll2Divider, Pll3Divider, ClockOrderControl0, ClockOrderControl1, ClockOrderStatus, FrequencySweepControl1,
	GoldenNonceForSweepReturn, ReturnedGroupPatternStatus, NonceReturnedTimeout, ReturnedSinglePatternStatus}

func regDump(regAddr RegAddr, regVal uint32, debug bool) {
	switch regAddr {
	case ChipAddress:
		fmt.Printf("Chip Address : 0x%08X\n", regVal)
		fmt.Printf("  CHIP_ID = 0x%04X\n", (regVal>>16)&0xffff)
		fmt.Printf("  CORE_NUM = 0x%02X\n", (regVal>>8)&0xff)
		fmt.Printf("  ADDR = 0x%02X\n", regVal&0xff)
	case HashRate:
		fmt.Printf("Hash Rate : 0x%08X\n", regVal)
		fmt.Printf("  LONG = %01X\n", (regVal>>31)&0x01)
		fmt.Printf("  HASHRATE = 0x%08X\n", regVal&0x7fffffff)
	case PLL0Parameter:
		fmt.Printf("PLL0 Parameter : 0x%08X\n", regVal)
		fmt.Printf("  LOCKED = %01X\n", (regVal>>31)&0x01)
		fmt.Printf("  PLLEN = %01X\n", (regVal>>30)&0x01)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>28)&0x03)
		}
		fmt.Printf("  FBDIV = 0x%03X\n", (regVal>>16)&0xfff)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>14)&0x03)
		}
		fmt.Printf("  REFDIV = 0x%02X\n", (regVal>>8)&0x3f)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>7)&0x01)
		}
		fmt.Printf("  POSTDIV1 = %01X\n", (regVal>>4)&0x07)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>3)&0x01)
		}
		fmt.Printf("  POSTDIV2 = %01X\n", regVal&0x07)
	case ChipNonceOffset:
		fmt.Printf("Chip Nonce Offset : 0x%08X\n", regVal)
		fmt.Printf("  CNOV = %01X\n", (regVal>>31)&0x01)
		if debug {
			fmt.Printf("  Reserved = 0x%04X\n", (regVal>>16)&0x7fff)
		}
		fmt.Printf("  CNO = %04X\n", regVal&0xffff)
	case HashCountingNumber:
		fmt.Printf("Hash Counting Number : 0x%08X\n", regVal)
	case TicketMask:
		fmt.Printf("Ticket Mask : 0x%08X\n", regVal)
		fmt.Printf("  TM3 = 0x%02X\n", (regVal>>24)&0xff)
		fmt.Printf("  TM2 = 0x%02X\n", (regVal>>16)&0xff)
		fmt.Printf("  TM1 = 0x%02X\n", (regVal>>8)&0xff)
		fmt.Printf("  TM0 = 0x%02X\n", regVal&0xff)
	case MiscControl:
		fmt.Printf("Misc Control : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>28)&0x0f)
			fmt.Printf("  BT8D_8_5 = 0x%01X\n", (regVal>>24)&0x0f)
			fmt.Printf("  Reserved = %01X\n", (regVal>>23)&0x01)
		}
		fmt.Printf("  CORE_SRST = %01X\n", (regVal>>22)&0x01)
		fmt.Printf("  SPAT_NOD = %01X\n", (regVal>>21)&0x01)
		fmt.Printf("  RVS_K0 = %01X\n", (regVal>>20)&0x01)
		fmt.Printf("  DSCLK_SEL = %01X\n", (regVal>>18)&0x03)
		fmt.Printf("  TOPCLK_SEL = %01X\n", (regVal>>17)&0x01)
		fmt.Printf("  BCLK_SEL = %01X\n", (regVal>>16)&0x01)
		fmt.Printf("  RET_ERR_NONCE = %01X\n", (regVal>>15)&0x01)
		fmt.Printf("  RFS = %01X\n", (regVal>>14)&0x01)
		fmt.Printf("  INV_CLKO = %01X\n", (regVal>>13)&0x01)
		if debug {
			fmt.Printf("  BT8D_4_0 = 0x%01X\n", (regVal>>8)&0x0f)
		}
		fmt.Printf("  BT8D = %d\n", (regVal>>8)&0x0f+32*(regVal>>24)&0x0f)
		fmt.Printf("  RET_WORK_ERR_FLAG = %01X\n", (regVal>>7)&0x01)
		fmt.Printf("  TFS = %01X\n", (regVal>>4)&0x07)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>2)&0x03)
		}
		fmt.Printf("  HASHRATE_TWS = %01X\n", regVal&0x03)
	case OrderedClockEnable:
		fmt.Printf("Ordered Clock Enable : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  Reserved = 0x%04X\n", (regVal>>16)&0xffff)
		}
		fmt.Printf("  CLKEN = 0x%04X\n", regVal&0xffff)
	case FastUARTConfiguration:
		fmt.Printf("Fast UART Configuration : 0x%08X\n", regVal)
		fmt.Printf("  DIV4_ODDSET = %01X\n", (regVal>>30)&0x03)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>28)&0x03)
		}
		fmt.Printf("  PLL3_DIV4 = 0x%01X\n", (regVal>>24)&0x0f)
		fmt.Printf("  USRC_ODDSET = %01X\n", (regVal>>22)&0x03)
		fmt.Printf("  USRC_DIV = 0x%02X\n", (regVal>>16)&0x3f)
		fmt.Printf("  ForceCoreEn = %01X\n", (regVal>>15)&0x01)
		fmt.Printf("  CLKO_SEL = %01X\n", (regVal>>14)&0x01)
		fmt.Printf("  CLKO_ODDSET = %01X\n", (regVal>>12)&0x03)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>8)&0x0f)
		}
		fmt.Printf("  CLKO_DIV = 0x%02X\n", regVal&0xff)
	case UARTRelay:
		fmt.Printf("UART Relay : 0x%08X\n", regVal)
		fmt.Printf("  GAP_CNT = 0x%04X\n", (regVal>>16)&0xffff)
		if debug {
			fmt.Printf("  Reserved = 0x%04X\n", (regVal>>2)&0x3fff)
		}
		fmt.Printf("  RO_RELAY_EN = %01X\n", (regVal>>1)&0x01)
		fmt.Printf("  CO_RELAY_EN = %01X\n", regVal&0x01)
	case TicketMask2:
		fmt.Printf("Ticket Mask2 : 0x%08X\n", regVal)
	case ExternalTemperatureSensorRead:
		fmt.Printf("External Temperature Sensor Read : 0x%08X\n", regVal)
		fmt.Printf("  LOCAL_TEMP_ADDR = 0x%02X\n", (regVal>>24)&0xff)
		fmt.Printf("  LOCAL_TEMP_DATA = 0x%02X\n", (regVal>>16)&0xff)
		fmt.Printf("  EXTERNAL_TEMP_ADDR = 0x%02X\n", (regVal>>8)&0xff)
		fmt.Printf("  EXTERNAL_TEMP_DATA = 0x%02X\n", regVal&0xff)
	case ErrorFlag:
		fmt.Printf("Error Flag : 0x%08X\n", regVal)
		fmt.Printf("  CMD_ERR_CNT = 0x%02X\n", regVal&0xff)
		fmt.Printf("  WORK_ERR_CNT = 0x%02X\n", (regVal>>8)&0xff)
		if debug {
			fmt.Printf("  Reserved = 0x%02X\n", (regVal>>16)&0xff)
		}
		fmt.Printf("  CORE_RESP_ERR = 0x%02X\n", (regVal>>24)&0xff)
	case NonceErrorCounter:
		fmt.Printf("Nonce Error Counter : 0x%08X\n", regVal)
	case NonceOverflowCounter:
		fmt.Printf("Nonce Overflow Counter : 0x%08X\n", regVal)
	case AnalogMuxControl:
		fmt.Printf("Analog Mux Control : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  Reserved = 0x%08X\n", regVal>>3)
		}
		fmt.Printf("  DIODE_VDD_MUX_SEL = %01X\n", regVal&0x07)
	case IoDriverStrenghtConfiguration:
		fmt.Printf("Io Driver Strenght Configuration : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>28)&0x0f)
		}
		fmt.Printf("  RF_DS = 0x%01X\n", (regVal>>24)&0x0f)
		fmt.Printf("  D3RS_DISA = %01X\n", (regVal>>23)&0x01)
		fmt.Printf("  D2RS_DISA = %01X\n", (regVal>>22)&0x01)
		fmt.Printf("  D1RS_DISA = %01X\n", (regVal>>21)&0x01)
		fmt.Printf("  D0RS_EN = %01X\n", (regVal>>20)&0x01)
		fmt.Printf("  R0_DS = 0x%01X\n", (regVal>>16)&0x0f)
		fmt.Printf("  CLKO_DS = 0x%01X\n", (regVal>>12)&0x0f)
		fmt.Printf("  NRSTO_DS = 0x%01X\n", (regVal>>8)&0x0f)
		fmt.Printf("  BO_DS = 0x%01X\n", (regVal>>4)&0x0f)
		fmt.Printf("  CO_DS = 0x%01X\n", regVal&0x0f)
	case TimeOut:
		fmt.Printf("Time Out : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  Reserved = 0x%04X\n", (regVal>>16)&0xffff)
		}
		fmt.Printf("  TMOUT = 0x%04X\n", regVal&0xffff)
	case PLL1Parameter:
		fmt.Printf("PLL1 Parameter : 0x%08X\n", regVal)
		fmt.Printf("  LOCKED = %01X\n", (regVal>>31)&0x01)
		fmt.Printf("  PLLEN = %01X\n", (regVal>>30)&0x01)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>28)&0x03)
		}
		fmt.Printf("  FBDIV = 0x%03X\n", (regVal>>16)&0xfff)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>14)&0x03)
		}
		fmt.Printf("  REFDIV = 0x%02X\n", (regVal>>8)&0x3f)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>7)&0x01)
		}
		fmt.Printf("  POSTDIV1 = %01X\n", (regVal>>4)&0x07)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>3)&0x01)
		}
		fmt.Printf("  POSTDIV2 = %01X\n", regVal&0x07)
	case PLL2Parameter:
		fmt.Printf("PLL2 Parameter : 0x%08X\n", regVal)
		fmt.Printf("  LOCKED = %01X\n", (regVal>>31)&0x01)
		fmt.Printf("  PLLEN = %01X\n", (regVal>>30)&0x01)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>28)&0x03)
		}
		fmt.Printf("  FBDIV = 0x%03X\n", (regVal>>16)&0xfff)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>14)&0x03)
		}
		fmt.Printf("  REFDIV = 0x%02X\n", (regVal>>8)&0x3f)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>7)&0x01)
		}
		fmt.Printf("  POSTDIV1 = %01X\n", (regVal>>4)&0x07)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>3)&0x01)
		}
		fmt.Printf("  POSTDIV2 = %01X\n", regVal&0x07)
	case PLL3Parameter:
		fmt.Printf("PLL3 Parameter : 0x%08X\n", regVal)
		fmt.Printf("  LOCKED = %01X\n", (regVal>>31)&0x01)
		fmt.Printf("  PLLEN = %01X\n", (regVal>>30)&0x01)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>28)&0x03)
		}
		fmt.Printf("  FBDIV = 0x%03X\n", (regVal>>16)&0xfff)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>14)&0x03)
		}
		fmt.Printf("  REFDIV = 0x%02X\n", (regVal>>8)&0x3f)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>7)&0x01)
		}
		fmt.Printf("  POSTDIV1 = %01X\n", (regVal>>4)&0x07)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>3)&0x01)
		}
		fmt.Printf("  POSTDIV2 = %01X\n", regVal&0x07)
	case OrderedClockMonitor:
		fmt.Printf("Ordered Clock Monitor : 0x%08X\n", regVal)
		fmt.Printf("  START = 0x%01X\n", (regVal>>31)&0x01)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>28)&0x07)
		}
		fmt.Printf("  CLK_SEL = 0x%01X\n", (regVal>>24)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%02X\n", (regVal>>16)&0xff)
		}
		fmt.Printf("  CLK_COUNT = 0x%04X\n", regVal&0xffff)
	case Pll0Divider:
		fmt.Printf("Pll0 Divider : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>28)&0x0f)
		}
		fmt.Printf("  PLL_DIV3 = 0x%01X\n", (regVal>>24)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>20)&0x0f)
		}
		fmt.Printf("  PLL_DIV2 = 0x%01X\n", (regVal>>16)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>12)&0x0f)
		}
		fmt.Printf("  PLL_DIV1 = 0x%01X\n", (regVal>>8)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>4)&0x0f)
		}
		fmt.Printf("  PLL_DIV0 = 0x%01X\n", regVal&0x0f)
	case Pll1Divider:
		fmt.Printf("Pll1 Divider : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>28)&0x0f)
		}
		fmt.Printf("  PLL_DIV3 = 0x%01X\n", (regVal>>24)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>20)&0x0f)
		}
		fmt.Printf("  PLL_DIV2 = 0x%01X\n", (regVal>>16)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>12)&0x0f)
		}
		fmt.Printf("  PLL_DIV1 = 0x%01X\n", (regVal>>8)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>4)&0x0f)
		}
		fmt.Printf("  PLL_DIV0 = 0x%01X\n", regVal&0x0f)
	case Pll2Divider:
		fmt.Printf("Pll2 Divider : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>28)&0x0f)
		}
		fmt.Printf("  PLL_DIV3 = 0x%01X\n", (regVal>>24)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>20)&0x0f)
		}
		fmt.Printf("  PLL_DIV2 = 0x%01X\n", (regVal>>16)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>12)&0x0f)
		}
		fmt.Printf("  PLL_DIV1 = 0x%01X\n", (regVal>>8)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>4)&0x0f)
		}
		fmt.Printf("  PLL_DIV0 = 0x%01X\n", regVal&0x0f)
	case Pll3Divider:
		fmt.Printf("Pll3 Divider : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>28)&0x0f)
		}
		fmt.Printf("  PLL_DIV3 = 0x%01X\n", (regVal>>24)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>20)&0x0f)
		}
		fmt.Printf("  PLL_DIV2 = 0x%01X\n", (regVal>>16)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>12)&0x0f)
		}
		fmt.Printf("  PLL_DIV1 = 0x%01X\n", (regVal>>8)&0x0f)
		if debug {
			fmt.Printf("  Reserved = 0x%01X\n", (regVal>>4)&0x0f)
		}
		fmt.Printf("  PLL_DIV0 = 0x%01X\n", regVal&0x0f)
	case ClockOrderControl0:
		fmt.Printf("Clock Order Control0 : 0x%08X\n", regVal)
		fmt.Printf("  CLK7_SEL = 0x%01X\n", (regVal>>28)&0x0f)
		fmt.Printf("  CLK6_SEL = 0x%01X\n", (regVal>>24)&0x0f)
		fmt.Printf("  CLK5_SEL = 0x%01X\n", (regVal>>20)&0x0f)
		fmt.Printf("  CLK4_SEL = 0x%01X\n", (regVal>>16)&0x0f)
		fmt.Printf("  CLK3_SEL = 0x%01X\n", (regVal>>12)&0x0f)
		fmt.Printf("  CLK2_SEL = 0x%01X\n", (regVal>>8)&0x0f)
		fmt.Printf("  CLK1_SEL = 0x%01X\n", (regVal>>4)&0x0f)
		fmt.Printf("  CLK0_SEL = 0x%01X\n", regVal&0x0f)
	case ClockOrderControl1:
		fmt.Printf("Clock Order Control1 : 0x%08X\n", regVal)
		fmt.Printf("  CLK15_SEL = 0x%01X\n", (regVal>>28)&0x0f)
		fmt.Printf("  CLK14_SEL = 0x%01X\n", (regVal>>24)&0x0f)
		fmt.Printf("  CLK13_SEL = 0x%01X\n", (regVal>>20)&0x0f)
		fmt.Printf("  CLK12_SEL = 0x%01X\n", (regVal>>16)&0x0f)
		fmt.Printf("  CLK11_SEL = 0x%01X\n", (regVal>>12)&0x0f)
		fmt.Printf("  CLK10_SEL = 0x%01X\n", (regVal>>8)&0x0f)
		fmt.Printf("  CLK9_SEL = 0x%01X\n", (regVal>>4)&0x0f)
		fmt.Printf("  CLK8_SEL = 0x%01X\n", regVal&0x0f)
	case ClockOrderStatus:
		fmt.Printf("Clock Order Status : 0x%08X\n", regVal)
	case FrequencySweepControl1:
		fmt.Printf("Frequency Sweep Control1 : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  Reserved = 0x%02X\n", regVal>>27)
		}
		fmt.Printf("  SWEEP_STATE = %01X\n", (regVal>>24)&0x07)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>21)&0x07)
		}
		fmt.Printf("  SWEEP_ST_ADDR = 0x%02X\n", (regVal>>21)&0x07)
		if debug {
			fmt.Printf("  Reserved = %01X\n", (regVal>>14)&0x03)
		}
		fmt.Printf("  ALL_CORE_CLK_SEL_CHANGE_ST = %01X\n", (regVal>>13)&0x01)
		fmt.Printf("  SWEEP_FAIL_LOCK_EN = %01X\n", (regVal>>12)&0x01)
		fmt.Printf("  SWEEP_RESET = %01X\n", (regVal>>11)&0x01)
		fmt.Printf("  CURR_PAT_ADDR = %01X\n", (regVal>>8)&0x07)
		fmt.Printf("  SWP_ONE_PAT_DONE = %01X\n", (regVal>>7)&0x01)
		fmt.Printf("  SWP_PAD_ADDR = %01X\n", (regVal>>4)&0x07)
		fmt.Printf("  SWP_DONE_ALL = %01X\n", (regVal>>3)&0x01)
		fmt.Printf("  SWP_ONGOING = %01X\n", (regVal>>2)&0x01)
		fmt.Printf("  SWP_TRIG = %01X\n", (regVal>>1)&0x01)
		fmt.Printf("  SWP_EN = %01X\n", regVal&0x01)
	case GoldenNonceForSweepReturn:
		fmt.Printf("Golden Nonce For Sweep Return : 0x%08X\n", regVal)
	case ReturnedGroupPatternStatus:
		fmt.Printf("Returned Group Pattern Status : 0x%08X\n", regVal)
	case NonceReturnedTimeout:
		fmt.Printf("Nonce Returned Timeout : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  Reserved = 0x%04X\n", (regVal>>16)&0xffff)
		}
		fmt.Printf("  SWEEP_TIMEOUT = 0x%04X\n", regVal&0xffff)
	case ReturnedSinglePatternStatus:
		fmt.Printf("Returned Single Pattern Status : 0x%08X\n", regVal)
	default:
		fmt.Printf("Unknown Register 0x%02X : 0x%08X\n", byte(regAddr), regVal)
	}
}
