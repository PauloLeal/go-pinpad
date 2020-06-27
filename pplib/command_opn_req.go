package pplib

import (
	"errors"
	"fmt"
	"strconv"
)

type OpnRequest struct {
	PsCom int    `json:"psCom"`
}

func (cmd *OpnRequest) GetName() string {
	return "OPN"
}

func (cmd *OpnRequest) Validate() error {
	if cmd.PsCom < 0 || cmd.PsCom > 99 {
		return errors.New("invalid PsCom value")
	}
	return nil
}

func (cmd *OpnRequest) Parse(rawData string) error {
	pr := NewPositionalReader(rawData)

	cmdName := pr.Read(3)
	if cmdName != "OPN" {
		return errors.New("cannot parse cmd command")
	}

	size, err := strconv.Atoi(pr.Read(3))
	if err != nil {
		return err
	}

	cmd.PsCom, err = strconv.Atoi(pr.Read(size))
	if err != nil {
		return err
	}

	return cmd.Validate()
}

func (cmd *OpnRequest) String() string {
	err := cmd.Validate()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s002%02d", cmd.GetName(), cmd.PsCom)
}
