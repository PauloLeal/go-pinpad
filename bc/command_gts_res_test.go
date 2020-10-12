package bc

import "testing"

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
	t.Run("bad status", auxTestCommand_Parse("GTS00A0100000000000", &GtsResponse{}, &GtsResponse{status: PP_OK, Timestamp: ""}, true))
	t.Run("Bad command name", auxTestCommand_Parse("GTX000000", &GtsResponse{}, &GtsResponse{status: PP_OK}, true))
	t.Run("bad command size", auxTestCommand_Parse("GTS00001A0000000000", &GtsResponse{}, &GtsResponse{status: PP_OK, Timestamp: ""}, true))
}

func TestGtsResponse_String(t *testing.T) {
	t.Run("status 0", auxTestCommand_String(&GtsResponse{status: PP_OK}, "GTS000000"))
}
