package pplib

import (
	"errors"
	"fmt"
	"strconv"
)

type OpnRequest struct {
	PsCom string `json:"psCom"`
}

func (cmd *OpnRequest) GetName() string {
	return "OPN"
}

func (cmd *OpnRequest) Validate() error {
	psc, err := strconv.Atoi(cmd.PsCom)
	if err != nil {
		return err
	}

	if psc < 0 || psc > 99 {
		return errors.New("invalid PsCom value")
	}
	return nil
}

func (cmd *OpnRequest) Parse(rawData string) error {
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
