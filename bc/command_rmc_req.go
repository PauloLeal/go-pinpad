package bc

import (
	"errors"
	"fmt"
	"strconv"
)

type RmcRequest struct {
	Message string `json:"message" pp_dataType:"A32"`
}

func (cmd *RmcRequest) GetName() string {
	return "RMC"
}

func (cmd *RmcRequest) Validate() error {
	return ValidateFieldsByDataTypeTag(*cmd)
}

func (cmd *RmcRequest) Parse(rawData string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	pr := NewPositionalReader(rawData)

	cmdName := pr.Read(3)
	if cmdName != cmd.GetName() {
		return errors.New(fmt.Sprintf("cannot parse %s command", cmd.GetName()))
	}

	size, err := strconv.Atoi(pr.Read(3))
	if err != nil {
		return err
	}

	cmd.Message = pr.Read(size)

	return cmd.Validate()
}

func (cmd *RmcRequest) String() string {
	err := cmd.Validate()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s032%32s", cmd.GetName(), cmd.Message)
}
