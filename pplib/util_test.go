package pplib

import "testing"

func TestIsDataTypeAlpha(t *testing.T) {
	t.Run("Good", auxTestIsDataTypeAlpha("A12344545FZ", true))
	t.Run("Special char", auxTestIsDataTypeAlpha("A123\x1244545FZ", false))
	t.Run("line break", auxTestIsDataTypeAlpha("A123\n44545FZ", false))
}

func TestIsDataTypeNumber(t *testing.T) {
	t.Run("Good", auxTestIsDataTypeNumber("316235736451723", true))
	t.Run("Special char", auxTestIsDataTypeNumber("3162\x1235736451723", false))
	t.Run("line break", auxTestIsDataTypeNumber("31623\n5736451723", false))
	t.Run("Alpha", auxTestIsDataTypeNumber("31623573645A1723", false))
}

func TestIsDataTypeHex(t *testing.T) {
	t.Run("Good", auxTestIsDataTypeHex("1234567890ABCDEFabcdef", true))
	t.Run("Special char", auxTestIsDataTypeHex("123456\x127890ABCDEFabcdef", false))
	t.Run("line break", auxTestIsDataTypeHex("1234567890\nABCDEFabcdef", false))
	t.Run("bad hex char", auxTestIsDataTypeHex("1234567890ABCDEFabcdefgG", false))
}
func TestIsValidDataType(t *testing.T) {
	t.Run("Good alpha", auxTestIsValidDataType("A12344545FZ", ALPHA, 11, true))
	t.Run("Good number", auxTestIsValidDataType("316235736451723", NUMBER, 15, true))
	t.Run("Good hex", auxTestIsValidDataType("1234567890ABCDEFabcdef", HEX, 22, true))
	t.Run("bad size", auxTestIsValidDataType("ASDASD3123123asdcvmxNB", ALPHA, 21, false))
	t.Run("bad dataType", auxTestIsValidDataType("A12344545FZ", "X", 11, false))
}

func auxTestIsDataTypeAlpha(data string, expectResult bool) func(*testing.T) {
	return func(t *testing.T) {
		if isDataTypeAlpha(data) != expectResult {
			t.Errorf("Expected to return %+v", expectResult)
		}
	}
}

func auxTestIsDataTypeNumber(data string, expectResult bool) func(*testing.T) {
	return func(t *testing.T) {
		if isDataTypeNumber(data) != expectResult {
			t.Errorf("Expected to return %+v", expectResult)
		}
	}
}

func auxTestIsDataTypeHex(data string, expectResult bool) func(*testing.T) {
	return func(t *testing.T) {
		if isDataTypeHex(data) != expectResult {
			t.Errorf("Expected to return %+v", expectResult)
		}
	}
}

func auxTestIsValidDataType(data string, dataType PP_DataType, size int, expectResult bool) func(*testing.T) {
	return func(t *testing.T) {
		if isValidDataType(data, dataType, size) != expectResult {
			t.Errorf("Expected to return %+v", expectResult)
		}
	}
}
