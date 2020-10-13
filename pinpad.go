package go_pinpad

import (
	"bytes"
	"errors"
	"strings"
)

const (
	EOT     byte = 0x07
	ACK     byte = 0x06
	NAK     byte = 0x15
	SYN     byte = 0x16
	ETB     byte = 0x17
	CAN     byte = 0x18
	TIMEOUT byte = 0x22
)

type IPinpad interface {
	Open() error
	Close() error
	Write(data []byte) error
	IsPinpadCommand(command string) bool
	Read() ([]byte, error)
}

type Pinpad struct {
	serial ISerial
	NoRmc  bool
}

func NewPinpad(port string) Pinpad {
	p := Pinpad{serial: &Serial{portName: port}}
	return p
}

func (p *Pinpad) IsPinpadCommand(command string) bool {
	var ppCommands = []string{"OPN", "CLO", "DSP", "DEX", "GKY", "GPN", "RMC", "GEN", "CKE", "GCR",
		"GOC", "FNC", "CHP", "CNG", "GIN", "ENB", "TLI", "TLR", "TLE", "GDU", "GTS", "DWK"}

	for _, c := range ppCommands {
		if strings.Index(command, c) == 0 {
			return true
		}
	}

	return false
}

func (p *Pinpad) Open() error {
	err := p.serial.Open()
	if err != nil {
		return err
	}
	return nil
}

func (p *Pinpad) Close() error {
	err := p.serial.Close()
	if err != nil {
		return err
	}
	return nil
}

func (p *Pinpad) Write(data []byte) error {
	bufferWithETB := make([]byte, 0)
	for _, bwe := range data {
		bufferWithETB = append(bufferWithETB, bwe)
	}
	bufferWithETB = append(bufferWithETB, ETB)

	lrc := calcLRC(bufferWithETB)
	lrcBytes := getLrcBytes(lrc)

	b := make([]byte, 0)
	b = append(b, SYN)
	b = append(b, bufferWithETB...)
	b = append(b, lrcBytes[0])
	b = append(b, lrcBytes[1])

	_, err := p.serial.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pinpad) Read() ([]byte, error) {
	readBuffer := make([]byte, 0)

	b := make([]byte, 1)

	etbFound := false
	for i := 0; i < 3; {
		_, err := p.serial.Read(b)
		if err != nil {
			return []byte{}, err
		}

		readBuffer = append(readBuffer, b[0])

		if !etbFound && bytes.IndexByte([]byte{NAK, CAN, TIMEOUT}, b[0]) >= 0 {
			break
		}

		if b[0] == ETB {
			etbFound = true
		}

		if etbFound {
			i++
		}
	}

	for {
		if readBuffer[0] == 0 || readBuffer[0] == ACK {
			readBuffer = readBuffer[1:]
		} else {
			break
		}
	}

	switch readBuffer[0] {
	case SYN:
		return parseReadData(readBuffer)
	case NAK:
		return nil, errors.New("serial Read NAK")
	case TIMEOUT:
		return nil, errors.New("serial Read Timeout")
	default:
		return nil, errors.New("serial Read error")
	}
}

func parseReadData(data []byte) ([]byte, error) {
	if data[0] != SYN {
		return nil, errors.New("serial Read not SYN")
	}

	tmp := make([]byte, 0)
	p := 0
	for i, d := range data[1:] {
		if d == ETB {
			p = i
			break
		}
		tmp = append(tmp, d)
	}

	dataLrcByte0 := data[p+1+1]
	dataLrcByte1 := data[p+1+2]

	lrcData := append(tmp, ETB)
	lrcBytes := getLrcBytes(calcLRC(lrcData))

	if dataLrcByte0 != lrcBytes[0] || dataLrcByte1 != lrcBytes[1] {
		return nil, errors.New("data LRC mismatch")
	}

	return tmp, nil
}

const CRCMASK int = 4129

func calcLRC(buffer []byte) int {
	var crc = 0
	var wData = 0
	for i := 0; i < len(buffer); i++ {
		wData = int(buffer[i])
		wData <<= 8
		wData &= 0xffff
		for j := 0; j < 8; j++ {
			if ((crc ^ wData) & 0x8000) != 0 {
				temp := crc << 1 & 0xffff
				crc = temp ^ CRCMASK
			} else {
				crc <<= 1
			}
			crc &= 0xffff
			wData <<= 1
			wData &= 0xffff
		}
	}
	crc &= 0xffff
	return crc
}

func getLrcBytes(crc int) []byte {
	b := make([]byte, 2)
	b[0] = byte((crc & 0xff00) >> 8)
	b[1] = byte(crc & 0xff)
	return b
}
