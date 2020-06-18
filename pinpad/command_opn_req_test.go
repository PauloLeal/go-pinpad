package pinpad

import (
	"fmt"
	"testing"
)

func TestOpnRequest(t *testing.T) {
	opn := OpnRequest{}

	if _, ok := interface{}(&opn).(Command); !ok {
		t.Error(fmt.Sprintf("Expected OpnRequest to implement interface Command"))
	}
}

func TestOpnRequest_Validate(t *testing.T) {
	t.Run("00", auxTestOpnRequest_Validate(OpnRequest{PsCom: 0}, false))
	t.Run("-1", auxTestOpnRequest_Validate(OpnRequest{PsCom: -1}, true))
	t.Run("100", auxTestOpnRequest_Validate(OpnRequest{PsCom: 100}, true))
	t.Run("1000", auxTestOpnRequest_Validate(OpnRequest{PsCom: 1000}, true))
}

func TestOpnRequest_Parse(t *testing.T) {
	t.Run("00", auxTestOpnRequest_Parse("OPN00200", OpnRequest{PsCom: 0}, false))
	t.Run("99", auxTestOpnRequest_Parse("OPN00299", OpnRequest{PsCom: 99}, false))
	t.Run("100", auxTestOpnRequest_Parse("OPN003100", OpnRequest{PsCom: 100}, true))
	t.Run("Non OPN", auxTestOpnRequest_Parse("XPN00200", OpnRequest{PsCom: 0}, true))
	t.Run("Bad size", auxTestOpnRequest_Parse("OPNAAA00", OpnRequest{PsCom: 0}, true))
	t.Run("Bad psCom", auxTestOpnRequest_Parse("OPN002AA", OpnRequest{PsCom: 0}, true))
}

func TestOpnRequest_String(t *testing.T) {
	t.Run("00", auxTestOpnRequest_String(OpnRequest{PsCom: 0}, "OPN00200"))
	t.Run("99", auxTestOpnRequest_String(OpnRequest{PsCom: 99}, "OPN00299"))
	t.Run("100", auxTestOpnRequest_String(OpnRequest{PsCom: 100}, ""))
}

func auxTestOpnRequest_Validate(opn OpnRequest, expectError bool) func(*testing.T) {
	return func(t *testing.T) {
		err := opn.Validate()
		if err == nil && expectError == true {
			t.Errorf("Expected to return error")
		} else if err != nil && expectError == false {
			t.Errorf("Expected to not return error")
		}
	}
}

func auxTestOpnRequest_Parse(rawData string, expectedResult OpnRequest, expectError bool) func(*testing.T) {
	return func(t *testing.T) {
		o := OpnRequest{}
		err := o.Parse(rawData)

		if err == nil && expectError == true {
			t.Errorf("Expected to return error")
		} else if err != nil && expectError == false {
			t.Errorf("Expected to not return error")
		}

		if o != expectedResult {
			t.Error(fmt.Sprintf("Expected Parse() of %s to be %+v. Got %+v", rawData, expectedResult, o))
		}
	}
}

func auxTestOpnRequest_String(opn OpnRequest, expectedResult string) func(*testing.T) {
	return func(t *testing.T) {
		s := opn.String()
		if s != expectedResult {
			t.Error(fmt.Sprintf("Expected String() of %+v to be %s. Got %s", opn, expectedResult, s))
		}
	}
}
