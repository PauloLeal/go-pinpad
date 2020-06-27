package pplib

import (
	"errors"
	"fmt"
	"strconv"
)

type GtsRequest struct {
	AcquirerIndex string `json:"acquirer_index"`
}

func (cmd *GtsRequest) GetName() string {
	return "GTS"
}

func (cmd *GtsRequest) Validate() error {
	ac, err := strconv.Atoi(cmd.AcquirerIndex)
	if err != nil {
		return err
	}

	if ac < 0 || ac > 99 {
		return errors.New("invalid AcquirerIndex value")
	}
	return nil
}

func (cmd *GtsRequest) Parse(rawData string) error {
	pr := NewPositionalReader(rawData)

	cmdName := pr.Read(3)
	if cmdName != cmd.GetName() {
		return errors.New("cannot parse cmd command")
	}

	size, err := strconv.Atoi(pr.Read(3))
	if err != nil {
		return err
	}

	cmd.AcquirerIndex = pr.Read(size)

	return cmd.Validate()
}

func (cmd *GtsRequest) String() string {
	err := cmd.Validate()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s002%02s", cmd.GetName(), cmd.AcquirerIndex)
}
