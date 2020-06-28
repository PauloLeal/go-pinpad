package pplib

import (
	"testing"
)

func TestOpnRequest_Validate(t *testing.T) {
	t.Run("00", auxTestCommand_Validate(&OpnRequest{PsCom: "00"}, false))
	t.Run("-1", auxTestCommand_Validate(&OpnRequest{PsCom: "-01"}, true))
	t.Run("100", auxTestCommand_Validate(&OpnRequest{PsCom: "100"}, true))
	t.Run("1000", auxTestCommand_Validate(&OpnRequest{PsCom: "1000"}, true))
}

func TestOpnRequest_Parse(t *testing.T) {
	t.Run("00", auxTestCommand_Parse("OPN00200", &OpnRequest{}, &OpnRequest{PsCom: "00"}, false))
	t.Run("99", auxTestCommand_Parse("OPN00299", &OpnRequest{}, &OpnRequest{PsCom: "99"}, false))
	t.Run("100", auxTestCommand_Parse("OPN003100", &OpnRequest{}, &OpnRequest{PsCom: "100"}, true))
	t.Run("Non OPN", auxTestCommand_Parse("XPN00200", &OpnRequest{}, &OpnRequest{PsCom: ""}, true))
	t.Run("Bad size", auxTestCommand_Parse("OPNAAA00", &OpnRequest{}, &OpnRequest{PsCom: ""}, true))
	t.Run("Bad psCom", auxTestCommand_Parse("OPN002AA", &OpnRequest{}, &OpnRequest{PsCom: "AA"}, true))
}

func TestOpnRequest_String(t *testing.T) {
	t.Run("00", auxTestCommand_String(&OpnRequest{PsCom: "0"}, "OPN00200"))
	t.Run("99", auxTestCommand_String(&OpnRequest{PsCom: "99"}, "OPN00299"))
	t.Run("100", auxTestCommand_String(&OpnRequest{PsCom: "100"}, ""))
}
