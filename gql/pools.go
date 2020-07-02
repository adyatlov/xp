package gql

import (
	"sync"

	"github.com/adyatlov/xp/data"
)

// stores *[]data.Object
var objectPool = sync.Pool{New: func() interface{} {
	buffer := make([]data.Object, 0, 1024)
	return &buffer
}}

// stores property values: *[]interface{}
var propertyPool = sync.Pool{New: func() interface{} {
	buffer := make([]interface{}, 0, 512)
	return &buffer
}}
