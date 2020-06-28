package pplib

import "regexp"

type PP_DataType string

const ALPHA, NUMBER, HEX PP_DataType = "A", "N", "H"

func isValidDataType(data string, dataType PP_DataType, size int) bool {
	if len(data) != size {
		return false
	}

	switch dataType {
	case ALPHA:
		return isDataTypeAlpha(data)
	case NUMBER:
		return isDataTypeNumber(data)
	case HEX:
		return isDataTypeHex(data)
	}

	return false
}

func isDataTypeAlpha(data string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]*$", data)
	return matched
}

func isDataTypeNumber(data string) bool {
	matched, _ := regexp.MatchString("^[0-9]*$", data)
	return matched
}

func isDataTypeHex(data string) bool {
	matched, _ := regexp.MatchString("^[a-fA-F0-9]*$", data)
	return matched
}
