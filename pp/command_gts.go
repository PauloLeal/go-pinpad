package pp

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

func (cmd *GtsRequest) Parse(rawData string) (err error) {
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

type GtsResponse struct {
	status    int
	Timestamp string `json:"timestamp" pp_dataType:"N10"`
}

func (cmd *GtsResponse) GetName() string {
	return "GTS"
}

func (cmd *GtsResponse) Validate() error {
	return ValidateFieldsByDataTypeTag(*cmd)
}

func (cmd *GtsResponse) Parse(rawData string) (err error) {
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

	status, err := strconv.Atoi(pr.Read(3))
	if err != nil {
		return err
	}

	timeStampSize, err := strconv.Atoi(pr.Read(3))
	if err != nil {
		return err
	}

	cmd.status = status
	cmd.Timestamp = pr.Read(timeStampSize)

	return cmd.Validate()
}

func (cmd *GtsResponse) String() string {
	return fmt.Sprintf("%s%03d%03d%s", cmd.GetName(), cmd.status, len(cmd.Timestamp), cmd.Timestamp)
}

func (cmd *GtsResponse) GetStatus() int {
	return cmd.status
}
