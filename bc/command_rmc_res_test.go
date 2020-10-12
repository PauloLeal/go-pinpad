package bc

import "testing"

func TestRmcResponse_GetStatus(t *testing.T) {
	t.Run("status 0", auxTestCommand_GetStatus(&RmcResponse{status: PP_OK}, PP_OK))
}

func TestRmcResponse_Validate(t *testing.T) {
	t.Run("empty", auxTestCommand_Validate(&RmcResponse{status: 0, NotifyMessage: ""}, false))
	t.Run("good notify", auxTestCommand_Validate(&RmcResponse{status: PP_NOTIFY, NotifyMessage: goodRmcSample}, false))
	t.Run("<bad notify", auxTestCommand_Validate(&RmcResponse{status: PP_NOTIFY}, true))
}

func TestRmcResponse_Parse(t *testing.T) {
	t.Run("status 0", auxTestCommand_Parse("RMC000000", &RmcResponse{}, &RmcResponse{status: PP_OK}, false))
	t.Run("status 0", auxTestCommand_Parse("RMC000032"+goodRmcSample, &RmcResponse{}, &RmcResponse{status: PP_OK, NotifyMessage: goodRmcSample}, false))
	t.Run("status 0", auxTestCommand_Parse("RMC002032"+goodRmcSample, &RmcResponse{}, &RmcResponse{status: PP_NOTIFY, NotifyMessage: goodRmcSample}, false))
	t.Run("bad status", auxTestCommand_Parse("RMC00A0100000000000", &RmcResponse{}, &RmcResponse{}, true))
	t.Run("Bad command name", auxTestCommand_Parse("RMX000000", &RmcResponse{}, &RmcResponse{}, true))
	t.Run("bad command size", auxTestCommand_Parse("RMC00001A0000000000", &RmcResponse{}, &RmcResponse{}, true))
	t.Run("bad param size", auxTestCommand_Parse("RMC0000110000000000", &RmcResponse{}, &RmcResponse{}, true))
}

func TestRmcResponse_String(t *testing.T) {
	t.Run("status 0", auxTestCommand_String(&RmcResponse{status: PP_OK}, "RMC000000"))
}
