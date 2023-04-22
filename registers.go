package bm13xx

import "fmt"

type CoreRegID byte

const (
	ClockDelayCtrl     CoreRegID = 0
	ProcessMonitorCtrl CoreRegID = 1
	ProcessMonitorData CoreRegID = 2
	CoreError          CoreRegID = 3
	CoreEnable         CoreRegID = 4
	HashClockCtrl      CoreRegID = 5
	HashClockCounter   CoreRegID = 6
	SweepClockCtrl     CoreRegID = 7
)

var lastCoreRegID CoreRegID
var allCoreRegisters []CoreRegID = []CoreRegID{ClockDelayCtrl, ProcessMonitorCtrl, ProcessMonitorData,
	CoreError, CoreEnable, HashClockCtrl, HashClockCounter, SweepClockCtrl}

func dumpCoreReg(regID CoreRegID, regVal uint16, debug bool) {
	switch regID {
	case ClockDelayCtrl:
		fmt.Printf("  Clock Delay Control : 0x%04X\n", regVal)
		if debug {
			fmt.Printf("    BIT[15:8] Reserved = %02X\n", (regVal>>8)&0xff)
		}
		fmt.Printf("    BIT[7:6] CCDLY_SEL = %01X\n", (regVal>>6)&0x03)
		fmt.Printf("    BIT[5:4] PWTH_SEL = %01X\n", (regVal>>4)&0x03)
		fmt.Printf("    BIT[3] HASH_CLKEN = %01X\n", (regVal>>3)&0x01)
		fmt.Printf("    BIT[2] MMEN = %01X\n", (regVal>>2)&0x01)
		if debug {
			fmt.Printf("    BIT[1] Reserved = %01X\n", (regVal>>1)&0x01)
		}
		fmt.Printf("    BIT[0] SWPF_MODE = %d\n", regVal&0x01)
	case ProcessMonitorCtrl:
		fmt.Printf("  Process Monitor Control : 0x%04X\n", regVal)
		if debug {
			fmt.Printf("    BIT[15:3] Reserved = %04X\n", (regVal>>3)&0x1fff)
		}
		fmt.Printf("    BIT[2] PM_START = %d\n", (regVal>>2)&0x01)
		fmt.Printf("    BIT[1:0] PM_SEL = %d\n", regVal&0x03)
	case ProcessMonitorData:
		fmt.Printf("  Process Monitor Data : 0x%04X\n", regVal)
		fmt.Printf("    BIT[15:0] FREQ_CNT = 0x%04X\n", regVal)
	case CoreError:
		fmt.Printf("  Core Error : 0x%04X\n", regVal)
		if debug {
			fmt.Printf("    BIT[15:5] Reserved = %03X\n", (regVal>>5)&0x7ff)
		}
		fmt.Printf("    BIT[4] INI_NONCE_ERR = %d\n", (regVal>>4)&0x01)
		fmt.Printf("    BIT[3:0] CMD_ERR_CNT = %01X\n", regVal&0x0f)
	case CoreEnable:
		fmt.Printf("  Core Enable : 0x%04X\n", regVal)
		if debug {
			fmt.Printf("    BIT[15:8] Reserved = %02X\n", (regVal>>8)&0xff)
		}
		fmt.Printf("    BIT[7:0] CORE_EN_I = %02X\n", regVal&0xff)
	case HashClockCtrl:
		fmt.Printf("  Process Monitor Control : 0x%04X\n", regVal)
		if debug {
			fmt.Printf("    BIT[15:8] Reserved = %02X\n", (regVal>>3)&0x1fff)
		}
		fmt.Printf("    BIT[7:0] CLOCK_CTRL = %02X\n", regVal&0xff)
	case HashClockCounter:
		fmt.Printf("  Process Monitor Control : 0x%04X\n", regVal)
		if debug {
			fmt.Printf("    BIT[15:8] Reserved = %02X\n", (regVal>>3)&0x1fff)
		}
		fmt.Printf("    BIT[7:0] CLOCK_CNT = %02X\n", regVal&0xff)
	case SweepClockCtrl:
		fmt.Printf("  Process Monitor Control : 0x%04X\n", regVal)
		if debug {
			fmt.Printf("    BIT[15:8] Reserved = %02X\n", (regVal>>3)&0x1fff)
		}
		fmt.Printf("    BIT[7] SWPF_MODE = %d\n", (regVal>>7)&0x01)
		if debug {
			fmt.Printf("    BIT[6:4] Reserved = %01X\n", (regVal>>4)&0x0f)
		}
		fmt.Printf("    BIT[3:0] CLK_SEL = %01X\n", regVal&0x0f)
	default:
		fmt.Printf("  Unknown Core Register %d : 0x%04X\n", byte(regID), regVal)
	}
}

type RegAddr byte

const (
	ChipAddress                   RegAddr = 0x00
	HashRate                      RegAddr = 0x04
	PLL0Parameter                 RegAddr = 0x08
	ChipNonceOffset               RegAddr = 0x0C
	HashCountingNumber            RegAddr = 0x10
	TicketMask                    RegAddr = 0x14
	MiscControl                   RegAddr = 0x18
	SomeTempRelated               RegAddr = 0x1C
	OrderedClockEnable            RegAddr = 0x20
	FastUARTConfiguration         RegAddr = 0x28
	UARTRelay                     RegAddr = 0x2C
	TicketMask2                   RegAddr = 0x38
	CoreRegisterControl           RegAddr = 0x3C
	CoreRegisterValue             RegAddr = 0x40
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
	TicketMask, MiscControl, SomeTempRelated, OrderedClockEnable, FastUARTConfiguration, UARTRelay, TicketMask2,
	CoreRegisterControl, CoreRegisterValue, ExternalTemperatureSensorRead,
	ErrorFlag, NonceErrorCounter, NonceOverflowCounter, AnalogMuxControl, IoDriverStrenghtConfiguration,
	TimeOut, PLL1Parameter, PLL2Parameter, PLL3Parameter, OrderedClockMonitor, Pll0Divider, Pll1Divider,
	Pll2Divider, Pll3Divider, ClockOrderControl0, ClockOrderControl1, ClockOrderStatus, FrequencySweepControl1,
	GoldenNonceForSweepReturn, ReturnedGroupPatternStatus, NonceReturnedTimeout, ReturnedSinglePatternStatus}

func dumpAsicReg(regAddr RegAddr, regVal uint32, debug bool) {
	switch regAddr {
	case ChipAddress:
		fmt.Printf("Chip Address : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31:16] CHIP_ID = 0x%04X\n", (regVal>>16)&0xffff)
		fmt.Printf("  BIT[15:8]  CORE_NUM = 0x%02X\n", (regVal>>8)&0xff)
		fmt.Printf("  BIT[7:0]   ADDR = 0x%02X\n", regVal&0xff)
	case HashRate:
		fmt.Printf("Hash Rate : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31]   LONG = %01X\n", (regVal>>31)&0x01)
		fmt.Printf("  BIT[30:0] HASHRATE = 0x%08X\n", regVal&0x7fffffff)
	case PLL0Parameter:
		fmt.Printf("PLL0 Parameter : 0x%08X\n", regVal)
		dumpPLLParam(regVal, debug)
	case ChipNonceOffset:
		fmt.Printf("Chip Nonce Offset : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31] CNOV = %01X\n", (regVal>>31)&0x01)
		if debug {
			fmt.Printf("  BIT[30:16] Reserved = 0x%04X\n", (regVal>>16)&0x7fff)
		}
		fmt.Printf("  BIT[15:0] CNO = %04X\n", regVal&0xffff)
	case HashCountingNumber:
		fmt.Printf("Hash Counting Number : 0x%08X\n", regVal)
	case TicketMask:
		fmt.Printf("Ticket Mask : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31:24] TM3 = 0x%02X\n", (regVal>>24)&0xff)
		fmt.Printf("  BIT[23:16] TM2 = 0x%02X\n", (regVal>>16)&0xff)
		fmt.Printf("  BIT[15:8]  TM1 = 0x%02X\n", (regVal>>8)&0xff)
		fmt.Printf("  BIT[7:0]   TM0 = 0x%02X\n", regVal&0xff)
	case CoreRegisterControl:
		fmt.Printf("Core Register Control : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31]    WR_RD#_MSB? = %d\n", (regVal>>31)&0x01)
		fmt.Printf("  BIT[30:16] Always0x7e00? = %04X\n", (regVal>>16)&0x7fff)
		fmt.Printf("  BIT[15]    WR_RD#_LSB? = %d\n", (regVal>>15)&0x01)
		fmt.Printf("  BIT[14:12] Always3? = %d\n", (regVal>>12)&0x07)
		fmt.Printf("  BIT[11:8]  CORE_REG_ID = %d\n", (regVal>>8)&0x0f)
		fmt.Printf("  BIT[7:0]   CORE_REG_VAL = 0x%02X\n", regVal&0xff)
		dumpCoreReg(CoreRegID((regVal>>8)&0x0f), uint16(regVal&0xff), debug)
		if (regVal>>15)&0x01 == 0 { // Read Core Register
			lastCoreRegID = CoreRegID((regVal >> 8) & 0x0f)
		}
	case CoreRegisterValue:
		fmt.Printf("Core Register Value : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31:16] CORE_ID = 0x%04X\n", (regVal>>16)&0xffff)
		fmt.Printf("  BIT[15:0]  CORE_REG_VAL = 0x%04X\n", regVal&0xffff)
		dumpCoreReg(lastCoreRegID, uint16(regVal&0xffff), debug)
	case MiscControl:
		fmt.Printf("Misc Control : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  BIT[31:23] Reserved = 0x%01X\n", (regVal>>28)&0x0f)
			fmt.Printf("  BIT[27:24] BT8D_8_5 = 0x%01X\n", (regVal>>24)&0x0f)
			fmt.Printf("  BIT[23]    Reserved = %01X\n", (regVal>>23)&0x01)
		}
		fmt.Printf("  BIT[22]    CORE_SRST = %01X\n", (regVal>>22)&0x01)
		fmt.Printf("  BIT[21]    SPAT_NOD = %01X\n", (regVal>>21)&0x01)
		fmt.Printf("  BIT[20]    RVS_K0 = %01X\n", (regVal>>20)&0x01)
		fmt.Printf("  BIT[19:18] DSCLK_SEL = %01X\n", (regVal>>18)&0x03)
		fmt.Printf("  BIT[17]    TOPCLK_SEL = %01X\n", (regVal>>17)&0x01)
		fmt.Printf("  BIT[16]    BCLK_SEL = %d (=1 if baud>3_000_000)\n", (regVal>>16)&0x01)
		fmt.Printf("  BIT[15]    RET_ERR_NONCE = %01X\n", (regVal>>15)&0x01)
		fmt.Printf("  BIT[14]    RFS = %01X\n", (regVal>>14)&0x01)
		fmt.Printf("  BIT[13]    INV_CLKO = %01X\n", (regVal>>13)&0x01)
		if debug {
			fmt.Printf("  BIT[12:8]  BT8D_4_0 = 0x%01X\n", (regVal>>8)&0x0f)
		}
		fmt.Printf("  calculated BT8D = %d (chip_divider)\n", (regVal>>8)&0x0f+32*(regVal>>24)&0x0f)
		fmt.Printf("  BIT[7]     RET_WORK_ERR_FLAG = %01X\n", (regVal>>7)&0x01)
		fmt.Printf("  BIT[6:4]   TFS = %01X\n", (regVal>>4)&0x07)
		if debug {
			fmt.Printf("  BIT[3:2]   Reserved = %01X\n", (regVal>>2)&0x03)
		}
		fmt.Printf("  BIT[1:0]   HASHRATE_TWS = %01X\n", regVal&0x03)
	case SomeTempRelated:
		fmt.Printf("Some Temperature Related : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31]    SOMETHING = %01X\n", (regVal>>31)&0x01)
		fmt.Printf("  BIT[30:25] SOMETHING = %01X\n", (regVal>>25)&0x01)
		fmt.Printf("  BIT[24]    SOMETHING = %01X\n", (regVal>>24)&0x01)
		fmt.Printf("  BIT[23:17] SOMETHING = %02X\n", (regVal>>17)&0x3f)
		fmt.Printf("  BIT[16]    SOMETHING = %01X\n", (regVal>>16)&0x01)
		fmt.Printf("  BIT[15:8]  REG = %02X\n", (regVal>>8)&0xff)
		fmt.Printf("  BIT[7:0]   TEMP_SENSOR_TYPE = %02X\n", regVal&0xff)
	case OrderedClockEnable:
		fmt.Printf("Ordered Clock Enable : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  BIT[31:16] Reserved = 0x%04X\n", (regVal>>16)&0xffff)
		}
		fmt.Printf("  BIT[15:0]  CLKEN = 0x%04X\n", regVal&0xffff)
	case FastUARTConfiguration:
		fmt.Printf("Fast UART Configuration : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31:30] DIV4_ODDSET = %01X\n", (regVal>>30)&0x03)
		if debug {
			fmt.Printf("  BIT[29:28] Reserved = %01X\n", (regVal>>28)&0x03)
		}
		fmt.Printf("  BIT[27:24] PLL3_DIV4 = 0x%01X\n", (regVal>>24)&0x0f)
		fmt.Printf("  BIT[23:22] USRC_ODDSET = %01X\n", (regVal>>22)&0x03)
		fmt.Printf("  BIT[21:16] USRC_DIV = 0x%02X\n", (regVal>>16)&0x3f)
		fmt.Printf("  BIT[15]    ForceCoreEn = %01X\n", (regVal>>15)&0x01)
		fmt.Printf("  BIT[14]    CLKO_SEL = %01X\n", (regVal>>14)&0x01)
		fmt.Printf("  BIT[13:12] CLKO_ODDSET = %01X\n", (regVal>>12)&0x03)
		if debug {
			fmt.Printf("  BIT[11:8]  Reserved = 0x%01X\n", (regVal>>8)&0x0f)
		}
		fmt.Printf("  BIT[7:0]   CLKO_DIV = 0x%02X\n", regVal&0xff)
	case UARTRelay:
		fmt.Printf("UART Relay : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31:16] GAP_CNT = 0x%04X\n", (regVal>>16)&0xffff)
		if debug {
			fmt.Printf("  BIT[15:2]  Reserved = 0x%04X\n", (regVal>>2)&0x3fff)
		}
		fmt.Printf("  BIT[1]     RO_RELAY_EN = %01X\n", (regVal>>1)&0x01)
		fmt.Printf("  BIT[0]     CO_RELAY_EN = %01X\n", regVal&0x01)
	case TicketMask2:
		fmt.Printf("Ticket Mask2 : 0x%08X\n", regVal)
	case ExternalTemperatureSensorRead:
		fmt.Printf("External Temperature Sensor Read : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31:24] LOCAL_TEMP_ADDR = 0x%02X\n", (regVal>>24)&0xff)
		fmt.Printf("  BIT[23:16] LOCAL_TEMP_DATA = 0x%02X\n", (regVal>>16)&0xff)
		fmt.Printf("  BIT[15:8]  EXTERNAL_TEMP_ADDR = 0x%02X\n", (regVal>>8)&0xff)
		fmt.Printf("  BIT[7:0]   EXTERNAL_TEMP_DATA = 0x%02X\n", regVal&0xff)
	case ErrorFlag:
		fmt.Printf("Error Flag : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31:24] CMD_ERR_CNT = 0x%02X\n", regVal&0xff)
		fmt.Printf("  BIT[23:16] WORK_ERR_CNT = 0x%02X\n", (regVal>>8)&0xff)
		if debug {
			fmt.Printf("  BIT[15:8]  Reserved = 0x%02X\n", (regVal>>16)&0xff)
		}
		fmt.Printf("  BIT[7:0]   CORE_RESP_ERR = 0x%02X\n", (regVal>>24)&0xff)
	case NonceErrorCounter:
		fmt.Printf("Nonce Error Counter : 0x%08X\n", regVal)
	case NonceOverflowCounter:
		fmt.Printf("Nonce Overflow Counter : 0x%08X\n", regVal)
	case AnalogMuxControl:
		fmt.Printf("Analog Mux Control : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  BIT[31:3] Reserved = 0x%08X\n", regVal>>3)
		}
		fmt.Printf("  BIT[2:0] DIODE_VDD_MUX_SEL = %01X\n", regVal&0x07)
	case IoDriverStrenghtConfiguration:
		fmt.Printf("Io Driver Strenght Configuration : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  BIT[31:28] Reserved = 0x%01X\n", (regVal>>28)&0x0f)
		}
		fmt.Printf("  BIT[27:24] RF_DS = 0x%01X\n", (regVal>>24)&0x0f)
		fmt.Printf("  BIT[23]    D3RS_DISA = %01X\n", (regVal>>23)&0x01)
		fmt.Printf("  BIT[22]    D2RS_DISA = %01X\n", (regVal>>22)&0x01)
		fmt.Printf("  BIT[21]    D1RS_DISA = %01X\n", (regVal>>21)&0x01)
		fmt.Printf("  BIT[20]    D0RS_EN = %01X\n", (regVal>>20)&0x01)
		fmt.Printf("  BIT[19:16] R0_DS = 0x%01X\n", (regVal>>16)&0x0f)
		fmt.Printf("  BIT[15:12] CLKO_DS = 0x%01X\n", (regVal>>12)&0x0f)
		fmt.Printf("  BIT[11:8]  NRSTO_DS = 0x%01X\n", (regVal>>8)&0x0f)
		fmt.Printf("  BIT[7:4]   BO_DS = 0x%01X\n", (regVal>>4)&0x0f)
		fmt.Printf("  BIT[3:0]   CO_DS = 0x%01X\n", regVal&0x0f)
	case TimeOut:
		fmt.Printf("Time Out : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  BIT[31:16] Reserved = 0x%04X\n", (regVal>>16)&0xffff)
		}
		fmt.Printf("  BIT[15:0]  TMOUT = 0x%04X\n", regVal&0xffff)
	case PLL1Parameter:
		fmt.Printf("PLL1 Parameter : 0x%08X\n", regVal)
		dumpPLLParam(regVal, debug)
	case PLL2Parameter:
		fmt.Printf("PLL2 Parameter : 0x%08X\n", regVal)
		dumpPLLParam(regVal, debug)
	case PLL3Parameter:
		fmt.Printf("PLL3 Parameter : 0x%08X\n", regVal)
		dumpPLLParam(regVal, debug)
	case OrderedClockMonitor:
		fmt.Printf("Ordered Clock Monitor : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31]    START = 0x%01X\n", (regVal>>31)&0x01)
		if debug {
			fmt.Printf("  BIT[30:28] Reserved = 0x%01X\n", (regVal>>28)&0x07)
		}
		fmt.Printf("  BIT[27:24] CLK_SEL = 0x%01X\n", (regVal>>24)&0x0f)
		if debug {
			fmt.Printf("  BIT[23:16] Reserved = 0x%02X\n", (regVal>>16)&0xff)
		}
		fmt.Printf("  BIT[15:0]  CLK_COUNT = 0x%04X\n", regVal&0xffff)
	case Pll0Divider:
		dumpPLLDiv(regVal, debug)
	case Pll1Divider:
		fmt.Printf("Pll1 Divider : 0x%08X\n", regVal)
		dumpPLLDiv(regVal, debug)
	case Pll2Divider:
		fmt.Printf("Pll2 Divider : 0x%08X\n", regVal)
		dumpPLLDiv(regVal, debug)
	case Pll3Divider:
		fmt.Printf("Pll3 Divider : 0x%08X\n", regVal)
		dumpPLLDiv(regVal, debug)
	case ClockOrderControl0:
		fmt.Printf("Clock Order Control0 : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31:28] CLK7_SEL = 0x%01X\n", (regVal>>28)&0x0f)
		fmt.Printf("  BIT[27:24] CLK6_SEL = 0x%01X\n", (regVal>>24)&0x0f)
		fmt.Printf("  BIT[23:20] CLK5_SEL = 0x%01X\n", (regVal>>20)&0x0f)
		fmt.Printf("  BIT[19:16] CLK4_SEL = 0x%01X\n", (regVal>>16)&0x0f)
		fmt.Printf("  BIT[15:12] CLK3_SEL = 0x%01X\n", (regVal>>12)&0x0f)
		fmt.Printf("  BIT[11:8]  CLK2_SEL = 0x%01X\n", (regVal>>8)&0x0f)
		fmt.Printf("  BIT[7:4]   CLK1_SEL = 0x%01X\n", (regVal>>4)&0x0f)
		fmt.Printf("  BIT[3:0]   CLK0_SEL = 0x%01X\n", regVal&0x0f)
	case ClockOrderControl1:
		fmt.Printf("Clock Order Control1 : 0x%08X\n", regVal)
		fmt.Printf("  BIT[31:28] CLK15_SEL = 0x%01X\n", (regVal>>28)&0x0f)
		fmt.Printf("  BIT[27:24] CLK14_SEL = 0x%01X\n", (regVal>>24)&0x0f)
		fmt.Printf("  BIT[23:20] CLK13_SEL = 0x%01X\n", (regVal>>20)&0x0f)
		fmt.Printf("  BIT[19:16] CLK12_SEL = 0x%01X\n", (regVal>>16)&0x0f)
		fmt.Printf("  BIT[15:12] CLK11_SEL = 0x%01X\n", (regVal>>12)&0x0f)
		fmt.Printf("  BIT[11:8]  CLK10_SEL = 0x%01X\n", (regVal>>8)&0x0f)
		fmt.Printf("  BIT[7:4]   CLK9_SEL = 0x%01X\n", (regVal>>4)&0x0f)
		fmt.Printf("  BIT[3:0]   CLK8_SEL = 0x%01X\n", regVal&0x0f)
	case ClockOrderStatus:
		fmt.Printf("Clock Order Status : 0x%08X\n", regVal)
	case FrequencySweepControl1:
		fmt.Printf("Frequency Sweep Control1 : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  BIT[31:27] Reserved = 0x%02X\n", regVal>>27)
		}
		fmt.Printf("  BIT[26:24] SWEEP_STATE = %01X\n", (regVal>>24)&0x07)
		if debug {
			fmt.Printf("  BIT[23:21] Reserved = %01X\n", (regVal>>21)&0x07)
		}
		fmt.Printf("  BIT[20:16] SWEEP_ST_ADDR = 0x%02X\n", (regVal>>21)&0x07)
		if debug {
			fmt.Printf("  BIT[15:14] Reserved = %01X\n", (regVal>>14)&0x03)
		}
		fmt.Printf("  BIT[13]    ALL_CORE_CLK_SEL_CHANGE_ST = %01X\n", (regVal>>13)&0x01)
		fmt.Printf("  BIT[12]    SWEEP_FAIL_LOCK_EN = %01X\n", (regVal>>12)&0x01)
		fmt.Printf("  BIT[11]    SWEEP_RESET = %01X\n", (regVal>>11)&0x01)
		fmt.Printf("  BIT[10:8]  CURR_PAT_ADDR = %01X\n", (regVal>>8)&0x07)
		fmt.Printf("  BIT[7]     SWP_ONE_PAT_DONE = %01X\n", (regVal>>7)&0x01)
		fmt.Printf("  BIT[6:4]   SWP_PAD_ADDR = %01X\n", (regVal>>4)&0x07)
		fmt.Printf("  BIT[3]     SWP_DONE_ALL = %01X\n", (regVal>>3)&0x01)
		fmt.Printf("  BIT[2]     SWP_ONGOING = %01X\n", (regVal>>2)&0x01)
		fmt.Printf("  BIT[1]     SWP_TRIG = %01X\n", (regVal>>1)&0x01)
		fmt.Printf("  BIT[0]     SWP_EN = %01X\n", regVal&0x01)
	case GoldenNonceForSweepReturn:
		fmt.Printf("Golden Nonce For Sweep Return : 0x%08X\n", regVal)
	case ReturnedGroupPatternStatus:
		fmt.Printf("Returned Group Pattern Status : 0x%08X\n", regVal)
	case NonceReturnedTimeout:
		fmt.Printf("Nonce Returned Timeout : 0x%08X\n", regVal)
		if debug {
			fmt.Printf("  BIT[31:16] Reserved = 0x%04X\n", (regVal>>16)&0xffff)
		}
		fmt.Printf("  BIT[15:0]  SWEEP_TIMEOUT = 0x%04X\n", regVal&0xffff)
	case ReturnedSinglePatternStatus:
		fmt.Printf("Returned Single Pattern Status : 0x%08X\n", regVal)
	default:
		fmt.Printf("Unknown Register 0x%02X : 0x%08X\n", byte(regAddr), regVal)
	}
}

func dumpPLLParam(regVal uint32, debug bool) {
	fmt.Printf("  BIT[31] LOCKED = %01X\n", (regVal>>31)&0x01)
	fmt.Printf("  BIT[30] PLLEN = %01X\n", (regVal>>30)&0x01)
	if debug {
		fmt.Printf("  BIT[29:28] Reserved = %01X\n", (regVal>>28)&0x03)
	}
	fmt.Printf("  BIT[27:16] FBDIV = 0x%03X\n", (regVal>>16)&0xfff)
	if debug {
		fmt.Printf("  BIT[15:14] Reserved = %01X\n", (regVal>>14)&0x03)
	}
	fmt.Printf("  BIT[13:8] REFDIV = 0x%02X\n", (regVal>>8)&0x3f)
	if debug {
		fmt.Printf("  BIT[7] Reserved = %01X\n", (regVal>>7)&0x01)
	}
	fmt.Printf("  BIT[6:4] POSTDIV1 = %01X\n", (regVal>>4)&0x07)
	if debug {
		fmt.Printf("  BIT[3] Reserved = %01X\n", (regVal>>3)&0x01)
	}
	fmt.Printf("  BIT[2:0] POSTDIV2 = %01X\n", regVal&0x07)
}

func dumpPLLDiv(regVal uint32, debug bool) {
	fmt.Printf("Pll0 Divider : 0x%08X\n", regVal)
	if debug {
		fmt.Printf("  BIT[31:28] Reserved = 0x%01X\n", (regVal>>28)&0x0f)
	}
	fmt.Printf("  BIT[27:24] PLL_DIV3 = 0x%01X\n", (regVal>>24)&0x0f)
	if debug {
		fmt.Printf("  BIT[23:20] Reserved = 0x%01X\n", (regVal>>20)&0x0f)
	}
	fmt.Printf("  BIT[19:16] PLL_DIV2 = 0x%01X\n", (regVal>>16)&0x0f)
	if debug {
		fmt.Printf("  BIT[15:12] Reserved = 0x%01X\n", (regVal>>12)&0x0f)
	}
	fmt.Printf("  BIT[11:8]  PLL_DIV1 = 0x%01X\n", (regVal>>8)&0x0f)
	if debug {
		fmt.Printf("  BIT[7:4]   Reserved = 0x%01X\n", (regVal>>4)&0x0f)
	}
	fmt.Printf("  BIT[3:0]   PLL_DIV0 = 0x%01X\n", regVal&0x0f)
}
