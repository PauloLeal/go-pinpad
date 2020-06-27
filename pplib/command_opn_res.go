package pplib

import (
	"errors"
	"fmt"
	"strconv"
)

type OpnResponse struct {
	status int
}

func (opn *OpnResponse) GetName() string {
	return "OPN"
}

func (opn *OpnResponse) Validate() error {
	return nil
}

func (opn *OpnResponse) Parse(rawData string) error {
	pr := NewPositionalReader(rawData)

	cmd := pr.Read(3)
	if cmd != "OPN" {
		return errors.New("cannot parse opn command")
	}

	status, err := strconv.Atoi(pr.Read(3))
	if err != nil {
		return err
	}

	zeroes := pr.Read(3)
	if zeroes != "000" {
		return errors.New("cannot parse opn command")
	}

	opn.status = status

	return opn.Validate()
}

func (opn *OpnResponse) String() string {
	return fmt.Sprintf("%s%03d000", opn.GetName(), opn.status)
}

func (opn *OpnResponse) GetStatus() int {
	return opn.status
}
