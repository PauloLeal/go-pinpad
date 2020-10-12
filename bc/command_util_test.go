package bc

import (
	"fmt"
	"reflect"
	"testing"
)

func auxTestCommand_GetStatus(cmd CommandResponse, expectdStatus int) func(*testing.T) {
	return func(t *testing.T) {
		s := cmd.GetStatus()
		if s != expectdStatus {
			t.Errorf("Expected GetStatus() of `%#v` to be %d. Got %d.", cmd, expectdStatus, s)
		}
	}
}

func auxTestCommand_Validate(cmd Command, expectError bool) func(*testing.T) {
	return func(t *testing.T) {
		err := cmd.Validate()
		if err == nil && expectError == true {
			t.Errorf("Expected to return error")
		} else if err != nil && expectError == false {
			t.Errorf("Expected to not return error")
		}
	}
}

func auxTestCommand_Parse(rawData string, target Command, expectedResult Command, expectError bool) func(*testing.T) {
	return func(t *testing.T) {
		err := target.Parse(rawData)

		if err == nil && expectError == true {
			t.Errorf("Expected to return error")
		} else if err != nil && expectError == false {
			t.Errorf("Expected to not return error")
		}

		if err != nil && expectError == true {
			return
		}

		if !reflect.DeepEqual(target, expectedResult) {
			t.Error(fmt.Sprintf("Expected Parse() of '%s' to be '%#v'. Got '%#v'.", rawData, expectedResult, target))
		}
	}
}

func auxTestCommand_String(cmd Command, expectedResult string) func(*testing.T) {
	return func(t *testing.T) {
		s := cmd.String()
		if s != expectedResult {
			t.Error(fmt.Sprintf("Expected String() of '%#v' to be '%s'. Got '%s'.", cmd, expectedResult, s))
		}
	}
}
