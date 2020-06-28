package pplib

import (
	"errors"
	"fmt"
	"strconv"
)

type GtsResponse struct {
	status    int
	Timestamp string `json:"timestamp"`
}

func (cmd *GtsResponse) GetName() string {
	return "GTS"
}

func (cmd *GtsResponse) Validate() error {
	if !isValidDataType(cmd.Timestamp, NUMBER, 10) {
		return errors.New("invalid Timestamp value")
	}
	return nil
}

func (cmd *GtsResponse) Parse(rawData string) error {
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
