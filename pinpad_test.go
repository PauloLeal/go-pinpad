package go_pinpad

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

type SerialMock struct {
	isOpen     bool
	forceError bool
	buffer     []byte
	rc         int
}

func (sm *SerialMock) Open() error {
	if sm.forceError {
		return errors.New("error")
	}

	sm.isOpen = true
	return nil
}

func (sm *SerialMock) Read(data []byte) (n int, err error) {
	if sm.forceError {
		return 0, errors.New("error")
	}

	data[0] = sm.buffer[sm.rc]
	sm.rc++
	return len(data), nil
}

func (sm *SerialMock) Write(data []byte) (n int, err error) {
	if sm.forceError {
		return 0, errors.New("error")
	}

	sm.buffer = make([]byte, len(data))
	copy(sm.buffer, data)
	sm.rc = 0
	return len(sm.buffer), nil
}

func (sm *SerialMock) Close() error {
	if sm.forceError {
		return errors.New("error")
	}

	sm.isOpen = false
	return nil
}

func TestNewPinpad(t *testing.T) {
	tf := func(portName string) func(*testing.T) {
		return func(t *testing.T) {
			p := NewPinpad(portName)
			if _, ok := interface{}(p.serial).(ISerial); !ok {
				t.Error(fmt.Sprintf("Expected NewPinpad to return a Pinpad with an embedded Serial"))
			}
		}
	}

	t.Run("New", tf("/dev/port"))
}

func TestPinpad_IsPinpadCommand(t *testing.T) {
	tf := func(command string, expectedResult bool) func(*testing.T) {
		return func(t *testing.T) {
			p := NewPinpad("")
			if p.IsPinpadCommand(command) != expectedResult {
				m := ""
				if !expectedResult {
					m = " not"
				}
				t.Errorf("Expected %s%s to be a zec command", command, m)
			}
		}
	}

	t.Run("OPN", tf("OPN", true))
	t.Run("GCR", tf("GCR", true))
	t.Run("GOC", tf("GOC", true))
	t.Run("DSA", tf("DSA", false))
	t.Run("CPS", tf("CPS", false))
	t.Run("CPS", tf("XXX", false))
}

func TestPinpad_Open(t *testing.T) {
	tf := func(forceError bool, expectError bool, isOpen bool) func(*testing.T) {
		return func(t *testing.T) {
			serial := SerialMock{forceError: forceError}
			p := Pinpad{serial: &serial}
			err := p.Open()
			if err == nil && expectError == true {
				t.Errorf("Expected Open to return error")
			} else if err != nil && expectError == false {
				t.Errorf("Expected Open to not return error")
			}

			if serial.isOpen != isOpen {
				t.Errorf("Expected isOpen to be %v. Got %v", isOpen, serial.isOpen)
			}
		}
	}

	t.Run("Success", tf(false, false, true))
	t.Run("Failure", tf(true, true, false))
}

func TestPinpad_Close(t *testing.T) {
	tf := func(forceError bool, expectError bool, isOpen bool) func(*testing.T) {
		return func(t *testing.T) {
			serial := SerialMock{forceError: forceError, isOpen: true}
			p := Pinpad{serial: &serial}

			err := p.Close()

			if err == nil && expectError == true {
				t.Errorf("Expected Open to return error")
			} else if err != nil && expectError == false {
				t.Errorf("Expected Open to not return error")
			}

			if serial.isOpen != isOpen {
				t.Errorf("Expected isOpen to be %v. Got %v", isOpen, serial.isOpen)
			}
		}
	}

	t.Run("Success", tf(false, false, false))
	t.Run("Failure", tf(true, true, true))
}

func TestPinpad_Write(t *testing.T) {
	tf := func(forceError bool, expectError bool, data []byte, serialData []byte) func(*testing.T) {
		return func(t *testing.T) {
			serial := SerialMock{forceError: forceError}
			p := Pinpad{serial: &serial}
			err := p.Write(data)

			if err == nil && expectError == true {
				t.Errorf("Expected Open to return error")
			} else if err != nil && expectError == false {
				t.Errorf("Expected Open to not return error")
			}

			if bytes.Compare(serialData, serial.buffer) != 0 {
				t.Errorf("Expected serial buffer to be %+v. Got %+v", data, serial.buffer)
			}
		}
	}

	t.Run("Success", tf(false, false, []byte("OPN00200"), []byte{22, 79, 80, 78, 48, 48, 50, 48, 48, 23, 163, 136}))
	t.Run("Failure", tf(true, true, []byte("OPN00200"), []byte{}))
}

func TestPinpad_Read(t *testing.T) {
	tf := func(forceError bool, expectError bool, serialData []byte, data []byte) func(*testing.T) {
		return func(t *testing.T) {
			serial := SerialMock{forceError: forceError, buffer: serialData}
			p := Pinpad{serial: &serial}
			d := make([]byte, 1024)
			d, err := p.Read()

			if err == nil && expectError == true {
				t.Errorf("Expected Open to return error")
			} else if err != nil && expectError == false {
				t.Errorf("Expected Open to not return error")
			}

			if bytes.Compare(d, data) != 0 {
				t.Errorf("Expected serial buffer to be %+v. Got %+v", data, d)
			}
		}
	}

	t.Run("Success ACK", tf(false, false, []byte{ACK, SYN, 79, 80, 78, 48, 48, 48, ETB, 119, 94}, []byte("OPN000")))
	t.Run("Success Multiple ACK", tf(false, false, []byte{ACK, ACK, ACK, SYN, 79, 80, 78, 48, 48, 48, ETB, 119, 94}, []byte("OPN000")))
	t.Run("Success Multiple 0 before ACK", tf(false, false, []byte{0, 0, 0, ACK, ACK, ACK, SYN, 79, 80, 78, 48, 48, 48, ETB, 119, 94}, []byte("OPN000")))
	t.Run("Success SYN", tf(false, false, []byte{SYN, 79, 80, 78, 48, 48, 48, ETB, 119, 94}, []byte("OPN000")))
	t.Run("Failure", tf(true, true, []byte{}, []byte{}))
	t.Run("Failure NAK", tf(false, true, []byte{NAK}, []byte{}))
	t.Run("Failure TIMEOUT", tf(false, true, []byte{TIMEOUT}, []byte{}))
	t.Run("Failure DEFAULT", tf(false, true, []byte{CAN}, []byte{}))
	t.Run("Failure ERROR", tf(true, true, []byte{0}, []byte{}))
}

func TestParseReadData(t *testing.T) {
	tf := func(expectError bool, serialData []byte, data []byte) func(*testing.T) {
		return func(t *testing.T) {
			d, err := parseReadData(serialData)

			if err == nil && expectError == true {
				t.Errorf("Expected Open to return error")
			} else if err != nil && expectError == false {
				t.Errorf("Expected Open to not return error")
			}

			if bytes.Compare(d, data) != 0 {
				t.Errorf("Expected parsed Data to be %+v. Got %+v", data, d)
			}
		}
	}

	t.Run("Success SYN start", tf(false, []byte{SYN, 79, 80, 78, 48, 48, 48, ETB, 119, 94}, []byte("OPN000")))
	t.Run("Failure No SYN Start", tf(true, []byte{ACK, 79, 80, 78, 48, 48, 48, ETB, 119, 94}, []byte{}))
	t.Run("Failure No ETB", tf(true, []byte{SYN, 79, 80, 78, 48, 48, 48, 119, 94}, []byte{}))
	t.Run("Failure Bad LRC byte 1", tf(true, []byte{SYN, 79, 80, 78, 48, 48, 48, ETB, 0, 94}, []byte{}))
	t.Run("Failure Bad LRC byte 2", tf(true, []byte{SYN, 79, 80, 78, 48, 48, 48, ETB, 119, 0}, []byte{}))
}

func TestPinpad_calcLRC(t *testing.T) {
	tf := func(buffer []byte, expectedResult int) func(*testing.T) {
		return func(t *testing.T) {
			b := buffer
			b = append(b, ETB)
			lrc := calcLRC(b)
			if lrc != expectedResult {
				t.Errorf("Expected LRC of %s to be %d. Got %d", buffer, expectedResult, lrc)
			}
		}
	}

	t.Run("OPN00200", tf([]byte("OPN00200"), 41864))
	t.Run("GIN00200", tf([]byte("GIN00200"), 61130))
}

func TestPinpad_getLrcBytes(t *testing.T) {
	tf := func(lrc int, expectedResult []byte) func(*testing.T) {
		return func(t *testing.T) {
			lrcBytes := getLrcBytes(lrc)

			if len(lrcBytes) != 2 {
				t.Errorf("Expected LRC lrcBytes to have size 2. Got size %d", len(lrcBytes))
			}

			for i, b := range lrcBytes {
				if b != expectedResult[i] {
					t.Errorf("Expected LRC lrcBytes of %d to be %+v. Got %+v", lrc, expectedResult, lrcBytes)
					break
				}
			}
		}
	}

	t.Run("41864", tf(41864, []byte{163, 136}))
	t.Run("61130", tf(61130, []byte{238, 202}))
}
