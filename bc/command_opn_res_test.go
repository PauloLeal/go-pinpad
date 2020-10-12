package bc

import "testing"

func TestOpnResponse_GetStatus(t *testing.T) {
	t.Run("status 0", auxTestCommand_GetStatus(&OpnResponse{status: PP_OK}, PP_OK))
	t.Run("status 14", auxTestCommand_GetStatus(&OpnResponse{status: PP_ALREADYOPEN}, PP_ALREADYOPEN))
}

func TestOpnResponse_Validate(t *testing.T) {
	t.Run("0", auxTestCommand_Validate(&OpnResponse{status: 0}, false))
}

func TestOpnResponse_Parse(t *testing.T) {
	t.Run("status 0", auxTestCommand_Parse("OPN000000", &OpnResponse{}, &OpnResponse{status: PP_OK}, false))
	t.Run("status 14", auxTestCommand_Parse("OPN014000", &OpnResponse{}, &OpnResponse{status: PP_ALREADYOPEN}, false))
	t.Run("bad command", auxTestCommand_Parse("XPN000000", &OpnResponse{}, &OpnResponse{}, true))
	t.Run("bad status", auxTestCommand_Parse("OPN00A000", &OpnResponse{}, &OpnResponse{}, true))
	t.Run("bad output", auxTestCommand_Parse("OPN00000A", &OpnResponse{}, &OpnResponse{}, true))
	t.Run("bad param size", auxTestCommand_Parse("OPN000001", &OpnResponse{}, &OpnResponse{}, true))
}

func TestOpnResponse_String(t *testing.T) {
	t.Run("status 0", auxTestCommand_String(&OpnResponse{status: PP_OK}, "OPN000000"))
	t.Run("status 14", auxTestCommand_String(&OpnResponse{status: PP_ALREADYOPEN}, "OPN014000"))
}
