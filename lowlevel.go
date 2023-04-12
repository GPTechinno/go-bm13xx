package bm13xx

import (
	"encoding/binary"
	"fmt"

	"github.com/snksoft/crc"
)

type cmd byte

const (
	sendJob       cmd = 0x21
	setChipAddr   cmd = 0x40
	writeRegister cmd = 0x41
	readRegister  cmd = 0x42
	chainInactive cmd = 0x43
)

func crc5(data []byte) byte {
	crc5 := crc.NewHash(&crc.Parameters{Width: 5, Polynomial: 0x05, Init: 0x1F, ReflectIn: false, ReflectOut: false, FinalXor: 0x00})
	return byte(crc5.CalculateCRC(data))
}

func crc16(data []byte) uint16 {
	crc16 := crc.NewHash(&crc.Parameters{Width: 16, Polynomial: 0x1021, Init: 0xD165, ReflectIn: false, ReflectOut: false, FinalXor: 0x0000})
	return uint16(crc16.CalculateCRC(data))
}

func (c *Chain) sendCommand(cmd cmd, all bool, chipAddr byte, regAddr RegAddr, data []byte) (int, error) {
	frame := []byte{byte(cmd), 0x05, chipAddr, byte(regAddr)}
	if all {
		frame[0] += 0x10
	}
	frame[1] = 0x05 + byte(len(data))
	frame = append(frame, data...)
	frame = append(frame, crc5(frame))
	if c.is139x {
		frame = append([]byte{0x55, 0xAA}, frame...)
	}
	return c.port.Write(frame)
}

func (c *Chain) SetChipAddr(chipAddr byte) error {
	_, err := c.sendCommand(setChipAddr, false, chipAddr, 0, nil)
	return err
}

func (c *Chain) WriteRegister(all bool, chipAddr byte, regAddr RegAddr, regVal uint32) error {
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, regVal)
	_, err := c.sendCommand(writeRegister, all, chipAddr, regAddr, data)
	return err
}

func (c *Chain) ReadRegister(all bool, chipAddr byte, regAddr RegAddr) error {
	_, err := c.sendCommand(readRegister, all, chipAddr, regAddr, nil)
	return err
}

func (c *Chain) GetResponse() (uint32, byte, byte, error) {
	respLen := 7
	if c.is139x {
		respLen += 2
	}
	resp := make([]byte, respLen)
	i, err := c.port.Read(resp)
	if err != nil {
		return 0, 0, 0, err
	}
	if i != respLen {
		return 0, 0, 0, fmt.Errorf("uncomplete resp")
	}
	if c.is139x {
		if resp[0] != 0xAA || resp[1] != 0x55 {
			return 0, 0, 0, fmt.Errorf("bad preamble")
		}
		resp = resp[2:]
	}
	if crc5(resp) != 0x00 {
		return 0, 0, 0, fmt.Errorf("bad crc5")
	}
	return binary.BigEndian.Uint32(resp), resp[4], resp[5], nil
}

func (c *Chain) Inactive() error {
	_, err := c.sendCommand(chainInactive, true, 0, 0, nil)
	return err
}
