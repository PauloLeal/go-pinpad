package pinpad

import "testing"

func TestOpnResponse_GetStatus(t *testing.T) {
	t.Run("status 0", auxTestCommand_GetStatus(&OpnResponse{Status: PP_OK}, PP_OK))
	t.Run("status 14", auxTestCommand_GetStatus(&OpnResponse{Status: PP_ALREADYOPEN}, PP_ALREADYOPEN))
}

func TestOpnResponse_Validate(t *testing.T) {
	t.Run("0", auxTestCommand_Validate(&OpnResponse{Status: 0}, false))
}

func TestOpnResponse_Parse(t *testing.T) {
	t.Run("status 0", auxTestCommand_Parse("OPN000000", &OpnResponse{}, &OpnResponse{Status: PP_OK}, false))
	t.Run("status 14", auxTestCommand_Parse("OPN014000", &OpnResponse{}, &OpnResponse{Status: PP_ALREADYOPEN}, false))
	t.Run("bad command", auxTestCommand_Parse("XPN000000", &OpnResponse{}, &OpnResponse{Status: 0}, true))
	t.Run("bad status", auxTestCommand_Parse("OPN00A000", &OpnResponse{}, &OpnResponse{Status: 0}, true))
	t.Run("bad output", auxTestCommand_Parse("OPN00000A", &OpnResponse{}, &OpnResponse{Status: 0}, true))
}

func TestOpnResponse_String(t *testing.T) {
	t.Run("status 0", auxTestCommand_String(&OpnResponse{Status: PP_OK}, "OPN000000"))
	t.Run("status 14", auxTestCommand_String(&OpnResponse{Status: PP_ALREADYOPEN}, "OPN014000"))
}
