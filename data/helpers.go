package data

import (
	"fmt"
	"time"
)

func ValidatePropertyValue(t PropertyValueType, v interface{}) error {
	var ok bool
	expect := ""
	switch t {
	case PVTBool:
		_, ok = v.(bool)
		expect = "bool"
	case PVTString:
		_, ok = v.(string)
		expect = "string"
	case PVTInteger:
		_, ok = v.(int)
		expect = "int"
	case PVTReal:
		_, ok = v.(float64)
		expect = "float64"
	case PVTPercentage:
		_, ok = v.(float64)
		expect = "float64"
	case PVTVersion:
		_, ok = v.(string)
		expect = "string"
	case PVTTimestamp:
		_, ok = v.(time.Time)
		expect = "time.Time"
	case PVTType:
		_, ok = v.(string)
		expect = "string"
	case PVTFile:
		_, ok = v.(string)
		expect = "string"
	case PVTObject:
		_, ok = v.(ObjectLink)
		expect = "data.ObjectLink"
	}
	if !ok {
		return fmt.Errorf(" Expected property type %s, got %T", expect, v)
	}
	return nil
}
