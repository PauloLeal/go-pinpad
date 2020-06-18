package pinpad

import (
	"errors"
	"fmt"
	"strconv"
)

type OpnRequest struct {
	PsCom int `json:"psCom"`
}

func (opn *OpnRequest) Validate() error {
	if opn.PsCom < 0 || opn.PsCom > 99 {
		return errors.New("invalid PsCom value")
	}
	return nil
}

func (opn *OpnRequest) Parse(rawData string) error {
	cursor := 0
	readN := func(n int) string {
		r := rawData[cursor : cursor+n]
		cursor += n
		return r
	}

	cmd := readN(3)
	if cmd != "OPN" {
		return errors.New("cannot parse opn command")
	}

	size, err := strconv.Atoi(readN(3))
	if err != nil {
		return err
	}

	opn.PsCom, err = strconv.Atoi(readN(size))
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

	return fmt.Sprintf("OPN002%02d", opn.PsCom)
}
