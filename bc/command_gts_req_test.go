package bc

import (
	"testing"
)

func TestGtsRequest_Validate(t *testing.T) {
	t.Run("00", auxTestCommand_Validate(&GtsRequest{AcquirerIndex: "01"}, false))
	t.Run("00", auxTestCommand_Validate(&GtsRequest{AcquirerIndex: "00"}, false))
	t.Run("00", auxTestCommand_Validate(&GtsRequest{AcquirerIndex: "100"}, true))
}

func TestGtsRequest_Parse(t *testing.T) {
	t.Run("00", auxTestCommand_Parse("GTS00200", &GtsRequest{}, &GtsRequest{AcquirerIndex: "00"}, false))
	t.Run("01", auxTestCommand_Parse("GTS00201", &GtsRequest{}, &GtsRequest{AcquirerIndex: "01"}, false))
	t.Run("no params", auxTestCommand_Parse("GTS000", &GtsRequest{}, &GtsRequest{}, true))
	t.Run("bad size", auxTestCommand_Parse("GTS003100", &GtsRequest{}, &GtsRequest{}, true))
	t.Run("bad size format", auxTestCommand_Parse("GTS0AA00", &GtsRequest{}, &GtsRequest{}, true))
	t.Run("bad name", auxTestCommand_Parse("GTX00200", &GtsRequest{}, &GtsRequest{}, true))
	t.Run("bad param size", auxTestCommand_Parse("GTS00300", &GtsRequest{}, &GtsRequest{}, true))
}

func TestGtsRequest_String(t *testing.T) {
	t.Run("00", auxTestCommand_String(&GtsRequest{AcquirerIndex: "01"}, "GTS00201"))
	t.Run("bad size", auxTestCommand_String(&GtsRequest{AcquirerIndex: "100"}, ""))
}
