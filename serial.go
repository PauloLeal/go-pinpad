package go_pinpad

import (
	"github.com/jacobsa/go-serial/serial"
	"io"
)

type ISerial interface {
	Open() error
	Read(p []byte) (int, error)
	Write(p []byte) (int, error)
	Close() error
}

type Serial struct {
	portName string
	port     io.ReadWriteCloser
	isOpen   bool
}

func (s *Serial) Open() error {
	if s.isOpen {
		return nil
	}

	options := serial.OpenOptions{
		PortName:        s.portName,
		BaudRate:        115200,
		ParityMode:      serial.PARITY_ODD,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	var err error
	s.port, err = serial.Open(options)
	if err != nil {
		return err
	}

	s.isOpen = true
	return nil
}

func (s *Serial) Read(data []byte) (int, error) {
	return s.port.Read(data)
}

func (s *Serial) Write(bytes []byte) (n int, err error) {
	return s.port.Write(bytes)
}

func (s *Serial) Close() error {
	err := s.port.Close()

	if !s.isOpen {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}
