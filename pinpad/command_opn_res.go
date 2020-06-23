package pinpad

import (
	"errors"
	"fmt"
	"strconv"
)

type OpnResponse struct {
	Status int `json:"status"`
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

	opn.Status = status

	return opn.Validate()
}

func (opn *OpnResponse) String() string {
	return fmt.Sprintf("OPN%03d000", opn.Status)
}

func (opn *OpnResponse) GetStatus() int {
	return opn.Status
}
