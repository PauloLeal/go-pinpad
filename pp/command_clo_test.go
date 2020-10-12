package pp

import (
	"testing"
)

const goodCloSample = "     go emv                     "

func TestCloRequest_Validate(t *testing.T) {
	t.Run("32", auxTestCommand_Validate(&CloRequest{IdleMessage: goodCloSample}, false))
	t.Run("!32", auxTestCommand_Validate(&CloRequest{IdleMessage: "RETIRE O CARTAO"}, true))
}

func TestCloRequest_Parse(t *testing.T) {
	t.Run("good", auxTestCommand_Parse("CLO032"+goodCloSample, &CloRequest{}, &CloRequest{IdleMessage: goodCloSample}, false))
	t.Run("no params 15", auxTestCommand_Parse("CLO015RETIRE O CARTAO", &CloRequest{}, &CloRequest{}, true))
	t.Run("no params 32", auxTestCommand_Parse("CLO032RETIRE O CARTAO", &CloRequest{}, &CloRequest{}, true))
	t.Run("bad size", auxTestCommand_Parse("CLOA32"+goodCloSample, &CloRequest{}, &CloRequest{IdleMessage: goodCloSample}, true))
	t.Run("bad name", auxTestCommand_Parse("RMX032"+goodCloSample, &CloRequest{}, &CloRequest{IdleMessage: goodCloSample}, true))

}

func TestCloRequest_String(t *testing.T) {
	t.Run("00", auxTestCommand_String(&CloRequest{IdleMessage: goodCloSample}, "CLO032"+goodCloSample))
	t.Run("bad size", auxTestCommand_String(&CloRequest{IdleMessage: "RETIRE O CARTAO"}, ""))
}

func TestCloResponse_GetStatus(t *testing.T) {
	t.Run("status 0", auxTestCommand_GetStatus(&CloResponse{status: PP_OK}, PP_OK))
}

func TestCloResponse_Validate(t *testing.T) {
	t.Run("empty", auxTestCommand_Validate(&CloResponse{status: 0}, false))
}

func TestCloResponse_Parse(t *testing.T) {
	t.Run("status 0", auxTestCommand_Parse("CLO000000", &CloResponse{}, &CloResponse{status: PP_OK}, false))
	t.Run("bad status", auxTestCommand_Parse("CLO00A0100000000000", &CloResponse{}, &CloResponse{}, true))
	t.Run("Bad command name", auxTestCommand_Parse("RMX000000", &CloResponse{}, &CloResponse{}, true))
	t.Run("bad command size", auxTestCommand_Parse("CLO00001A0000000000", &CloResponse{}, &CloResponse{}, true))
	t.Run("bad param size", auxTestCommand_Parse("CLO0000110000000000", &CloResponse{}, &CloResponse{}, true))
}

func TestCloResponse_String(t *testing.T) {
	t.Run("status 0", auxTestCommand_String(&CloResponse{status: PP_OK}, "CLO000000"))
}
