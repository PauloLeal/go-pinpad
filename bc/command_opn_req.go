package bc

import (
	"errors"
	"fmt"
	"strconv"
)

type OpnRequest struct {
	PsCom string `json:"psCom" pp_dataType:"N2"`
}

func (cmd *OpnRequest) GetName() string {
	return "OPN"
}

func (cmd *OpnRequest) Validate() error {
	return ValidateFieldsByDataTypeTag(*cmd)
}

func (cmd *OpnRequest) Parse(rawData string) (err error) {
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

	cmd.PsCom = pr.Read(size)

	return cmd.Validate()
}

func (cmd *OpnRequest) String() string {
	err := cmd.Validate()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s002%02s", cmd.GetName(), cmd.PsCom)
}
