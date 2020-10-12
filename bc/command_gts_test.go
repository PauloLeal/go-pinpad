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

func TestGtsResponse_GetStatus(t *testing.T) {
	t.Run("status 0", auxTestCommand_GetStatus(&GtsResponse{status: PP_OK}, PP_OK))
}

func TestGtsResponse_Validate(t *testing.T) {
	t.Run("0", auxTestCommand_Validate(&GtsResponse{status: 0, Timestamp: "0000000000"}, false))
	t.Run("empty", auxTestCommand_Validate(&GtsResponse{status: 0, Timestamp: ""}, true))
	t.Run("<10", auxTestCommand_Validate(&GtsResponse{status: 0, Timestamp: "000000000"}, true))
	t.Run(">10", auxTestCommand_Validate(&GtsResponse{status: 0, Timestamp: "00000000000"}, true))
}

func TestGtsResponse_Parse(t *testing.T) {
	t.Run("status 0", auxTestCommand_Parse("GTS0000100000000000", &GtsResponse{}, &GtsResponse{status: PP_OK, Timestamp: "0000000000"}, false))
	t.Run("bad status", auxTestCommand_Parse("GTS00A0100000000000", &GtsResponse{}, &GtsResponse{}, true))
	t.Run("Bad command name", auxTestCommand_Parse("GTX000000", &GtsResponse{}, &GtsResponse{}, true))
	t.Run("bad command size", auxTestCommand_Parse("GTS00001A0000000000", &GtsResponse{}, &GtsResponse{}, true))
	t.Run("bad param size", auxTestCommand_Parse("GTS0000110000000000", &GtsResponse{}, &GtsResponse{}, true))
}

func TestGtsResponse_String(t *testing.T) {
	t.Run("status 0", auxTestCommand_String(&GtsResponse{status: PP_OK}, "GTS000000"))
}
