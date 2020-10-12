package bc

import (
	"testing"
)

const goodRmcSample = "RETIRE O CARTAO                 "

func TestRmcRequest_Validate(t *testing.T) {
	t.Run("32", auxTestCommand_Validate(&RmcRequest{Message: goodRmcSample}, false))
	t.Run("!32", auxTestCommand_Validate(&RmcRequest{Message: "RETIRE O CARTAO"}, true))
}

func TestRmcRequest_Parse(t *testing.T) {
	t.Run("good", auxTestCommand_Parse("RMC032"+goodRmcSample, &RmcRequest{}, &RmcRequest{Message: goodRmcSample}, false))
	t.Run("no params 15", auxTestCommand_Parse("RMC015RETIRE O CARTAO", &RmcRequest{}, &RmcRequest{}, true))
	t.Run("no params 32", auxTestCommand_Parse("RMC032RETIRE O CARTAO", &RmcRequest{}, &RmcRequest{}, true))
	t.Run("bad size", auxTestCommand_Parse("RMCA32"+goodRmcSample, &RmcRequest{}, &RmcRequest{Message: goodRmcSample}, true))
	t.Run("bad name", auxTestCommand_Parse("RMX032"+goodRmcSample, &RmcRequest{}, &RmcRequest{Message: goodRmcSample}, true))

}

func TestRmcRequest_String(t *testing.T) {
	t.Run("00", auxTestCommand_String(&RmcRequest{Message: goodRmcSample}, "RMC032"+goodRmcSample))
	t.Run("bad size", auxTestCommand_String(&RmcRequest{Message: "RETIRE O CARTAO"}, ""))
}
