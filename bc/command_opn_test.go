package bc

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
	t.Run("100", auxTestCommand_Parse("OPN003100", &OpnRequest{}, &OpnRequest{}, true))
	t.Run("Non OPN", auxTestCommand_Parse("XPN00200", &OpnRequest{}, &OpnRequest{}, true))
	t.Run("Bad size", auxTestCommand_Parse("OPNAAA00", &OpnRequest{}, &OpnRequest{}, true))
	t.Run("Bad psCom", auxTestCommand_Parse("OPN002AA", &OpnRequest{}, &OpnRequest{}, true))
	t.Run("Bad param size", auxTestCommand_Parse("OPN00300", &OpnRequest{}, &OpnRequest{}, true))
}

func TestOpnRequest_String(t *testing.T) {
	t.Run("00", auxTestCommand_String(&OpnRequest{PsCom: "00"}, "OPN00200"))
	t.Run("99", auxTestCommand_String(&OpnRequest{PsCom: "99"}, "OPN00299"))
	t.Run("100", auxTestCommand_String(&OpnRequest{PsCom: "100"}, ""))
}

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
