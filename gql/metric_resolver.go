package gql

import (
	"strconv"
	"time"

	"github.com/adyatlov/xp/data"
)

type propertyResolver struct {
	property data.Property
}

func (r *propertyResolver) Type() *propertyTypeResolver {
	return &propertyTypeResolver{r.property.Type()}
}

func (r *propertyResolver) Value() (value string) {
	v := r.property.Value()
	switch r.property.Type().ValueType {
	case data.MVTBool:
		value = strconv.FormatBool(v.(bool))
	case data.MVTString:
		value = v.(string)
	case data.MVTInteger:
		value = strconv.Itoa(v.(int))
	case data.MVTReal:
		value = strconv.FormatFloat(v.(float64), 'f', 10, 64)
	case data.MVTPercentage:
		value = strconv.FormatFloat(v.(float64), 'f', 10, 64)
	case data.MVTVersion:
		value = v.(string)
	case data.MVTTimestamp:
		value = strconv.FormatInt(v.(time.Time).UnixNano()/1e6, 10)
	case data.MVTType:
		value = v.(string)
	case data.MVTFile:
		value = v.(string)
	case data.MVTObject:
		o := v.(data.Object)
		value = string(encodeUniqueId(o.Type().Name, o.Id()))
	default:
		panic("unknown property value type: " + r.property.Type().ValueType)
	}
	return
}
