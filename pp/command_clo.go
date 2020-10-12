package pp

import (
	"errors"
	"fmt"
	"strconv"
)

type CloRequest struct {
	IdleMessage string `json:"idle_message" pp_dataType:"A32"`
}

func (cmd *CloRequest) GetName() string {
	return "CLO"
}

func (cmd *CloRequest) Validate() error {
	return ValidateFieldsByDataTypeTag(*cmd)
}

func (cmd *CloRequest) Parse(rawData string) (err error) {
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

	cmd.IdleMessage = pr.Read(size)

	return cmd.Validate()
}

func (cmd *CloRequest) String() string {
	err := cmd.Validate()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s032%32s", cmd.GetName(), cmd.IdleMessage)
}

type CloResponse struct {
	status int
}

func (cmd *CloResponse) GetName() string {
	return "CLO"
}

func (cmd *CloResponse) Validate() error {
	return nil
}

func (cmd *CloResponse) Parse(rawData string) (err error) {
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

func (cmd *CloResponse) String() string {
	return fmt.Sprintf("%s%03d000", cmd.GetName(), cmd.status)
}

func (cmd *CloResponse) GetStatus() int {
	return cmd.status
}
