package bc

import (
	"errors"
	"fmt"
	"strconv"
)

type OpnResponse struct {
	status int
}

func (cmd *OpnResponse) GetName() string {
	return "OPN"
}

func (cmd *OpnResponse) Validate() error {
	return nil
}

func (cmd *OpnResponse) Parse(rawData string) (err error) {
	pr := NewPositionalReader(rawData)

	cmdName := pr.Read(3)
	if cmdName != cmd.GetName() {
		return errors.New(fmt.Sprintf("cannot parse %s command", cmd.GetName()))
	}

	status, err := strconv.Atoi(pr.Read(3))
	if err != nil {
		return err
	}

	zeroes := pr.Read(3)
	if zeroes != "000" {
		return errors.New(fmt.Sprintf("cannot parse %s command", cmd.GetName()))
	}

	cmd.status = status

	return cmd.Validate()
}

func (cmd *OpnResponse) String() string {
	return fmt.Sprintf("%s%03d000", cmd.GetName(), cmd.status)
}

func (cmd *OpnResponse) GetStatus() int {
	return cmd.status
}
