package bc

import "testing"

func TestIsDataTypeAlpha(t *testing.T) {
	tf := func(data string, expectResult bool) func(*testing.T) {
		return func(t *testing.T) {
			if isDataTypeAlpha(data) != expectResult {
				t.Errorf("Expected to return %+v", expectResult)
			}
		}
	}

	t.Run("Good", tf("A12344545FZ", true))
	t.Run("Special char", tf("A123\x1244545FZ", true))
	t.Run("line break", tf("A123\n44545FZ", false))
}

func TestIsDataTypeNumber(t *testing.T) {
	tf := func(data string, expectResult bool) func(*testing.T) {
		return func(t *testing.T) {
			if isDataTypeNumber(data) != expectResult {
				t.Errorf("Expected to return %+v", expectResult)
			}
		}
	}

	t.Run("Good", tf("316235736451723", true))
	t.Run("Special char", tf("3162\x1235736451723", false))
	t.Run("line break", tf("31623\n5736451723", false))
	t.Run("Alpha", tf("31623573645A1723", false))
}

func TestIsDataTypeHex(t *testing.T) {
	tf := func(data string, expectResult bool) func(*testing.T) {
		return func(t *testing.T) {
			if isDataTypeHex(data) != expectResult {
				t.Errorf("Expected to return %+v", expectResult)
			}
		}
	}

	t.Run("Good", tf("1234567890ABCDEFabcdef", true))
	t.Run("Special char", tf("123456\x127890ABCDEFabcdef", false))
	t.Run("line break", tf("1234567890\nABCDEFabcdef", false))
	t.Run("bad hex char", tf("1234567890ABCDEFabcdefgG", false))
}

func TestIsValidDataType(t *testing.T) {
	tf := func(data string, dataType PP_DataType, size int, expectResult bool) func(*testing.T) {
		return func(t *testing.T) {
			if isValidDataType(data, dataType, size) != expectResult {
				t.Errorf("Expected to return %+v", expectResult)
			}
		}
	}

	t.Run("Good alpha", tf("A12344545FZ", ALPHA, 11, true))
	t.Run("Good number", tf("316235736451723", NUMBER, 15, true))
	t.Run("Good hex", tf("1234567890ABCDEFabcdef", HEX, 22, true))
	t.Run("bad size", tf("ASDASD3123123asdcvmxNB", ALPHA, 21, false))
	t.Run("bad dataType", tf("A12344545FZ", "X", 11, false))
}
