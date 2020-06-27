package pplib

import (
	"errors"
	"fmt"
	"strconv"
)

type OpnRequest struct {
	PsCom int    `json:"psCom"`
}

func (opn *OpnRequest) GetName() string {
	return "OPN"
}

func (opn *OpnRequest) Validate() error {
	if opn.PsCom < 0 || opn.PsCom > 99 {
		return errors.New("invalid PsCom value")
	}
	return nil
}

func (opn *OpnRequest) Parse(rawData string) error {
	pr := NewPositionalReader(rawData)

	cmd := pr.Read(3)
	if cmd != "OPN" {
		return errors.New("cannot parse opn command")
	}

	size, err := strconv.Atoi(pr.Read(3))
	if err != nil {
		return err
	}

	opn.PsCom, err = strconv.Atoi(pr.Read(size))
	if err != nil {
		return err
	}

	return opn.Validate()
}

func (opn *OpnRequest) String() string {
	err := opn.Validate()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s002%02d", opn.GetName(), opn.PsCom)
}
