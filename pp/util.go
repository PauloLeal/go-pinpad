package pp

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

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
	matched, _ := regexp.MatchString("^.*$", data)
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

func ValidateFieldsByDataTypeTag(cmd interface{}) error {
	t := reflect.TypeOf(cmd)
	v := reflect.ValueOf(cmd)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		value := v.FieldByName(field.Name).String()
		tag := field.Tag.Get("pp_dataType")
		if len(tag) == 0 {
			continue
		}

		var dataType = PP_DataType(tag[:1])
		dataSize, _ := strconv.Atoi(tag[1:])

		if !isValidDataType(value, dataType, dataSize) {
			return errors.New(fmt.Sprintf("invalid %s value", field.Name))
		}
	}

	return nil
}
