package gql

import (
	"strconv"
	"time"

	"github.com/graph-gophers/graphql-go"

	"github.com/adyatlov/xp/data"
)

type propertyResolver struct {
	property data.Property
	id       graphql.ID
}

func (r *propertyResolver) Id() graphql.ID {
	return r.id
}

func (r *propertyResolver) Type() *propertyTypeResolver {
	return &propertyTypeResolver{r.property.Type()}
}

func (r *propertyResolver) Value() (value string) {
	v := r.property.Value()
	switch r.property.Type().Type {
	case data.PVTBool:
		value = strconv.FormatBool(v.(bool))
	case data.PVTString:
		value = v.(string)
	case data.PVTInteger:
		value = strconv.Itoa(v.(int))
	case data.PVTReal:
		value = strconv.FormatFloat(v.(float64), 'f', 10, 64)
	case data.PVTPercentage:
		value = strconv.FormatFloat(v.(float64), 'f', 10, 64)
	case data.PVTVersion:
		value = v.(string)
	case data.PVTTimestamp:
		value = strconv.FormatInt(v.(time.Time).UnixNano()/1e6, 10)
	case data.PVTType:
		value = v.(string)
	case data.PVTFile:
		value = v.(string)
	case data.PVTObject:
		o := v.(data.Object)
		pId := decodeId(r.id).(propertyId)
		oId := encodeId(objectId{
			datasetId:      pId.datasetId,
			ObjectTypeName: o.Type().Name,
			ObjectId:       o.Id(),
		})
		value = string(oId)
	default:
		panic("unknown property value type: " + strconv.Itoa(int(r.property.Type().Type)))
	}
	return
}
