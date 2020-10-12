package bc

import (
	"errors"
	"fmt"
	"strconv"
)

type GtsRequest struct {
	AcquirerIndex string `json:"acquirer_index" pp_dataType:"N2"`
}

func (cmd *GtsRequest) GetName() string {
	return "GTS"
}

func (cmd *GtsRequest) Validate() error {
	return ValidateFieldsByDataTypeTag(*cmd)
}

func (cmd *GtsRequest) Parse(rawData string) error {
	pr := NewPositionalReader(rawData)

	cmdName := pr.Read(3)
	if cmdName != cmd.GetName() {
		return errors.New(fmt.Sprintf("cannot parse %s command", cmd.GetName()))
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
