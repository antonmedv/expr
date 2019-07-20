package vm

import (
	"fmt"
	"reflect"
)

func equal(a, b interface{}) interface{} {
	switch x := a.(type) {
	case uint:
		switch y := b.(type) {
		case uint:
			return x == y
		case uint8:
			return uint8(x) == y
		case uint16:
			return uint16(x) == y
		case uint32:
			return uint32(x) == y
		case uint64:
			return uint64(x) == y
		case int:
			return int(x) == y
		case int8:
			return int8(x) == y
		case int16:
			return int16(x) == y
		case int32:
			return int32(x) == y
		case int64:
			return int64(x) == y
		case float32:
			return float32(x) == y
		case float64:
			return float64(x) == y
		}
	case uint8:
		switch y := b.(type) {
		case uint:
			return x == uint8(y)
		case uint8:
			return x == y
		case uint16:
			return uint16(x) == y
		case uint32:
			return uint32(x) == y
		case uint64:
			return uint64(x) == y
		case int:
			return int(x) == y
		case int8:
			return int8(x) == y
		case int16:
			return int16(x) == y
		case int32:
			return int32(x) == y
		case int64:
			return int64(x) == y
		case float32:
			return float32(x) == y
		case float64:
			return float64(x) == y
		}
	case uint16:
		switch y := b.(type) {
		case uint:
			return x == uint16(y)
		case uint8:
			return x == uint16(y)
		case uint16:
			return x == y
		case uint32:
			return uint32(x) == y
		case uint64:
			return uint64(x) == y
		case int:
			return int(x) == y
		case int8:
			return int8(x) == y
		case int16:
			return int16(x) == y
		case int32:
			return int32(x) == y
		case int64:
			return int64(x) == y
		case float32:
			return float32(x) == y
		case float64:
			return float64(x) == y
		}
	case uint32:
		switch y := b.(type) {
		case uint:
			return x == uint32(y)
		case uint8:
			return x == uint32(y)
		case uint16:
			return x == uint32(y)
		case uint32:
			return x == y
		case uint64:
			return uint64(x) == y
		case int:
			return int(x) == y
		case int8:
			return int8(x) == y
		case int16:
			return int16(x) == y
		case int32:
			return int32(x) == y
		case int64:
			return int64(x) == y
		case float32:
			return float32(x) == y
		case float64:
			return float64(x) == y
		}
	case uint64:
		switch y := b.(type) {
		case uint:
			return x == uint64(y)
		case uint8:
			return x == uint64(y)
		case uint16:
			return x == uint64(y)
		case uint32:
			return x == uint64(y)
		case uint64:
			return x == y
		case int:
			return int(x) == y
		case int8:
			return int8(x) == y
		case int16:
			return int16(x) == y
		case int32:
			return int32(x) == y
		case int64:
			return int64(x) == y
		case float32:
			return float32(x) == y
		case float64:
			return float64(x) == y
		}
	case int:
		switch y := b.(type) {
		case uint:
			return x == int(y)
		case uint8:
			return x == int(y)
		case uint16:
			return x == int(y)
		case uint32:
			return x == int(y)
		case uint64:
			return x == int(y)
		case int:
			return x == y
		case int8:
			return int8(x) == y
		case int16:
			return int16(x) == y
		case int32:
			return int32(x) == y
		case int64:
			return int64(x) == y
		case float32:
			return float32(x) == y
		case float64:
			return float64(x) == y
		}
	case int8:
		switch y := b.(type) {
		case uint:
			return x == int8(y)
		case uint8:
			return x == int8(y)
		case uint16:
			return x == int8(y)
		case uint32:
			return x == int8(y)
		case uint64:
			return x == int8(y)
		case int:
			return x == int8(y)
		case int8:
			return x == y
		case int16:
			return int16(x) == y
		case int32:
			return int32(x) == y
		case int64:
			return int64(x) == y
		case float32:
			return float32(x) == y
		case float64:
			return float64(x) == y
		}
	case int16:
		switch y := b.(type) {
		case uint:
			return x == int16(y)
		case uint8:
			return x == int16(y)
		case uint16:
			return x == int16(y)
		case uint32:
			return x == int16(y)
		case uint64:
			return x == int16(y)
		case int:
			return x == int16(y)
		case int8:
			return x == int16(y)
		case int16:
			return x == y
		case int32:
			return int32(x) == y
		case int64:
			return int64(x) == y
		case float32:
			return float32(x) == y
		case float64:
			return float64(x) == y
		}
	case int32:
		switch y := b.(type) {
		case uint:
			return x == int32(y)
		case uint8:
			return x == int32(y)
		case uint16:
			return x == int32(y)
		case uint32:
			return x == int32(y)
		case uint64:
			return x == int32(y)
		case int:
			return x == int32(y)
		case int8:
			return x == int32(y)
		case int16:
			return x == int32(y)
		case int32:
			return x == y
		case int64:
			return int64(x) == y
		case float32:
			return float32(x) == y
		case float64:
			return float64(x) == y
		}
	case int64:
		switch y := b.(type) {
		case uint:
			return x == int64(y)
		case uint8:
			return x == int64(y)
		case uint16:
			return x == int64(y)
		case uint32:
			return x == int64(y)
		case uint64:
			return x == int64(y)
		case int:
			return x == int64(y)
		case int8:
			return x == int64(y)
		case int16:
			return x == int64(y)
		case int32:
			return x == int64(y)
		case int64:
			return x == y
		case float32:
			return float32(x) == y
		case float64:
			return float64(x) == y
		}
	case float32:
		switch y := b.(type) {
		case uint:
			return x == float32(y)
		case uint8:
			return x == float32(y)
		case uint16:
			return x == float32(y)
		case uint32:
			return x == float32(y)
		case uint64:
			return x == float32(y)
		case int:
			return x == float32(y)
		case int8:
			return x == float32(y)
		case int16:
			return x == float32(y)
		case int32:
			return x == float32(y)
		case int64:
			return x == float32(y)
		case float32:
			return x == y
		case float64:
			return float64(x) == y
		}
	case float64:
		switch y := b.(type) {
		case uint:
			return x == float64(y)
		case uint8:
			return x == float64(y)
		case uint16:
			return x == float64(y)
		case uint32:
			return x == float64(y)
		case uint64:
			return x == float64(y)
		case int:
			return x == float64(y)
		case int8:
			return x == float64(y)
		case int16:
			return x == float64(y)
		case int32:
			return x == float64(y)
		case int64:
			return x == float64(y)
		case float32:
			return x == float64(y)
		case float64:
			return x == y
		}
	case string:
		switch y := b.(type) {
		case string:
			return x == y
		}
	}
	// Two nil values should be considered as equal.
	if isNil(a) && isNil(b) {
		return true
	}
	return reflect.DeepEqual(a, b)
}

func less(a, b interface{}) interface{} {
	switch x := a.(type) {
	case uint:
		switch y := b.(type) {
		case uint:
			return x < y
		case uint8:
			return uint8(x) < y
		case uint16:
			return uint16(x) < y
		case uint32:
			return uint32(x) < y
		case uint64:
			return uint64(x) < y
		case int:
			return int(x) < y
		case int8:
			return int8(x) < y
		case int16:
			return int16(x) < y
		case int32:
			return int32(x) < y
		case int64:
			return int64(x) < y
		case float32:
			return float32(x) < y
		case float64:
			return float64(x) < y
		}
	case uint8:
		switch y := b.(type) {
		case uint:
			return x < uint8(y)
		case uint8:
			return x < y
		case uint16:
			return uint16(x) < y
		case uint32:
			return uint32(x) < y
		case uint64:
			return uint64(x) < y
		case int:
			return int(x) < y
		case int8:
			return int8(x) < y
		case int16:
			return int16(x) < y
		case int32:
			return int32(x) < y
		case int64:
			return int64(x) < y
		case float32:
			return float32(x) < y
		case float64:
			return float64(x) < y
		}
	case uint16:
		switch y := b.(type) {
		case uint:
			return x < uint16(y)
		case uint8:
			return x < uint16(y)
		case uint16:
			return x < y
		case uint32:
			return uint32(x) < y
		case uint64:
			return uint64(x) < y
		case int:
			return int(x) < y
		case int8:
			return int8(x) < y
		case int16:
			return int16(x) < y
		case int32:
			return int32(x) < y
		case int64:
			return int64(x) < y
		case float32:
			return float32(x) < y
		case float64:
			return float64(x) < y
		}
	case uint32:
		switch y := b.(type) {
		case uint:
			return x < uint32(y)
		case uint8:
			return x < uint32(y)
		case uint16:
			return x < uint32(y)
		case uint32:
			return x < y
		case uint64:
			return uint64(x) < y
		case int:
			return int(x) < y
		case int8:
			return int8(x) < y
		case int16:
			return int16(x) < y
		case int32:
			return int32(x) < y
		case int64:
			return int64(x) < y
		case float32:
			return float32(x) < y
		case float64:
			return float64(x) < y
		}
	case uint64:
		switch y := b.(type) {
		case uint:
			return x < uint64(y)
		case uint8:
			return x < uint64(y)
		case uint16:
			return x < uint64(y)
		case uint32:
			return x < uint64(y)
		case uint64:
			return x < y
		case int:
			return int(x) < y
		case int8:
			return int8(x) < y
		case int16:
			return int16(x) < y
		case int32:
			return int32(x) < y
		case int64:
			return int64(x) < y
		case float32:
			return float32(x) < y
		case float64:
			return float64(x) < y
		}
	case int:
		switch y := b.(type) {
		case uint:
			return x < int(y)
		case uint8:
			return x < int(y)
		case uint16:
			return x < int(y)
		case uint32:
			return x < int(y)
		case uint64:
			return x < int(y)
		case int:
			return x < y
		case int8:
			return int8(x) < y
		case int16:
			return int16(x) < y
		case int32:
			return int32(x) < y
		case int64:
			return int64(x) < y
		case float32:
			return float32(x) < y
		case float64:
			return float64(x) < y
		}
	case int8:
		switch y := b.(type) {
		case uint:
			return x < int8(y)
		case uint8:
			return x < int8(y)
		case uint16:
			return x < int8(y)
		case uint32:
			return x < int8(y)
		case uint64:
			return x < int8(y)
		case int:
			return x < int8(y)
		case int8:
			return x < y
		case int16:
			return int16(x) < y
		case int32:
			return int32(x) < y
		case int64:
			return int64(x) < y
		case float32:
			return float32(x) < y
		case float64:
			return float64(x) < y
		}
	case int16:
		switch y := b.(type) {
		case uint:
			return x < int16(y)
		case uint8:
			return x < int16(y)
		case uint16:
			return x < int16(y)
		case uint32:
			return x < int16(y)
		case uint64:
			return x < int16(y)
		case int:
			return x < int16(y)
		case int8:
			return x < int16(y)
		case int16:
			return x < y
		case int32:
			return int32(x) < y
		case int64:
			return int64(x) < y
		case float32:
			return float32(x) < y
		case float64:
			return float64(x) < y
		}
	case int32:
		switch y := b.(type) {
		case uint:
			return x < int32(y)
		case uint8:
			return x < int32(y)
		case uint16:
			return x < int32(y)
		case uint32:
			return x < int32(y)
		case uint64:
			return x < int32(y)
		case int:
			return x < int32(y)
		case int8:
			return x < int32(y)
		case int16:
			return x < int32(y)
		case int32:
			return x < y
		case int64:
			return int64(x) < y
		case float32:
			return float32(x) < y
		case float64:
			return float64(x) < y
		}
	case int64:
		switch y := b.(type) {
		case uint:
			return x < int64(y)
		case uint8:
			return x < int64(y)
		case uint16:
			return x < int64(y)
		case uint32:
			return x < int64(y)
		case uint64:
			return x < int64(y)
		case int:
			return x < int64(y)
		case int8:
			return x < int64(y)
		case int16:
			return x < int64(y)
		case int32:
			return x < int64(y)
		case int64:
			return x < y
		case float32:
			return float32(x) < y
		case float64:
			return float64(x) < y
		}
	case float32:
		switch y := b.(type) {
		case uint:
			return x < float32(y)
		case uint8:
			return x < float32(y)
		case uint16:
			return x < float32(y)
		case uint32:
			return x < float32(y)
		case uint64:
			return x < float32(y)
		case int:
			return x < float32(y)
		case int8:
			return x < float32(y)
		case int16:
			return x < float32(y)
		case int32:
			return x < float32(y)
		case int64:
			return x < float32(y)
		case float32:
			return x < y
		case float64:
			return float64(x) < y
		}
	case float64:
		switch y := b.(type) {
		case uint:
			return x < float64(y)
		case uint8:
			return x < float64(y)
		case uint16:
			return x < float64(y)
		case uint32:
			return x < float64(y)
		case uint64:
			return x < float64(y)
		case int:
			return x < float64(y)
		case int8:
			return x < float64(y)
		case int16:
			return x < float64(y)
		case int32:
			return x < float64(y)
		case int64:
			return x < float64(y)
		case float32:
			return x < float64(y)
		case float64:
			return x < y
		}
	case string:
		switch y := b.(type) {
		case string:
			return x < y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "<", b))
}

func more(a, b interface{}) interface{} {
	switch x := a.(type) {
	case uint:
		switch y := b.(type) {
		case uint:
			return x > y
		case uint8:
			return uint8(x) > y
		case uint16:
			return uint16(x) > y
		case uint32:
			return uint32(x) > y
		case uint64:
			return uint64(x) > y
		case int:
			return int(x) > y
		case int8:
			return int8(x) > y
		case int16:
			return int16(x) > y
		case int32:
			return int32(x) > y
		case int64:
			return int64(x) > y
		case float32:
			return float32(x) > y
		case float64:
			return float64(x) > y
		}
	case uint8:
		switch y := b.(type) {
		case uint:
			return x > uint8(y)
		case uint8:
			return x > y
		case uint16:
			return uint16(x) > y
		case uint32:
			return uint32(x) > y
		case uint64:
			return uint64(x) > y
		case int:
			return int(x) > y
		case int8:
			return int8(x) > y
		case int16:
			return int16(x) > y
		case int32:
			return int32(x) > y
		case int64:
			return int64(x) > y
		case float32:
			return float32(x) > y
		case float64:
			return float64(x) > y
		}
	case uint16:
		switch y := b.(type) {
		case uint:
			return x > uint16(y)
		case uint8:
			return x > uint16(y)
		case uint16:
			return x > y
		case uint32:
			return uint32(x) > y
		case uint64:
			return uint64(x) > y
		case int:
			return int(x) > y
		case int8:
			return int8(x) > y
		case int16:
			return int16(x) > y
		case int32:
			return int32(x) > y
		case int64:
			return int64(x) > y
		case float32:
			return float32(x) > y
		case float64:
			return float64(x) > y
		}
	case uint32:
		switch y := b.(type) {
		case uint:
			return x > uint32(y)
		case uint8:
			return x > uint32(y)
		case uint16:
			return x > uint32(y)
		case uint32:
			return x > y
		case uint64:
			return uint64(x) > y
		case int:
			return int(x) > y
		case int8:
			return int8(x) > y
		case int16:
			return int16(x) > y
		case int32:
			return int32(x) > y
		case int64:
			return int64(x) > y
		case float32:
			return float32(x) > y
		case float64:
			return float64(x) > y
		}
	case uint64:
		switch y := b.(type) {
		case uint:
			return x > uint64(y)
		case uint8:
			return x > uint64(y)
		case uint16:
			return x > uint64(y)
		case uint32:
			return x > uint64(y)
		case uint64:
			return x > y
		case int:
			return int(x) > y
		case int8:
			return int8(x) > y
		case int16:
			return int16(x) > y
		case int32:
			return int32(x) > y
		case int64:
			return int64(x) > y
		case float32:
			return float32(x) > y
		case float64:
			return float64(x) > y
		}
	case int:
		switch y := b.(type) {
		case uint:
			return x > int(y)
		case uint8:
			return x > int(y)
		case uint16:
			return x > int(y)
		case uint32:
			return x > int(y)
		case uint64:
			return x > int(y)
		case int:
			return x > y
		case int8:
			return int8(x) > y
		case int16:
			return int16(x) > y
		case int32:
			return int32(x) > y
		case int64:
			return int64(x) > y
		case float32:
			return float32(x) > y
		case float64:
			return float64(x) > y
		}
	case int8:
		switch y := b.(type) {
		case uint:
			return x > int8(y)
		case uint8:
			return x > int8(y)
		case uint16:
			return x > int8(y)
		case uint32:
			return x > int8(y)
		case uint64:
			return x > int8(y)
		case int:
			return x > int8(y)
		case int8:
			return x > y
		case int16:
			return int16(x) > y
		case int32:
			return int32(x) > y
		case int64:
			return int64(x) > y
		case float32:
			return float32(x) > y
		case float64:
			return float64(x) > y
		}
	case int16:
		switch y := b.(type) {
		case uint:
			return x > int16(y)
		case uint8:
			return x > int16(y)
		case uint16:
			return x > int16(y)
		case uint32:
			return x > int16(y)
		case uint64:
			return x > int16(y)
		case int:
			return x > int16(y)
		case int8:
			return x > int16(y)
		case int16:
			return x > y
		case int32:
			return int32(x) > y
		case int64:
			return int64(x) > y
		case float32:
			return float32(x) > y
		case float64:
			return float64(x) > y
		}
	case int32:
		switch y := b.(type) {
		case uint:
			return x > int32(y)
		case uint8:
			return x > int32(y)
		case uint16:
			return x > int32(y)
		case uint32:
			return x > int32(y)
		case uint64:
			return x > int32(y)
		case int:
			return x > int32(y)
		case int8:
			return x > int32(y)
		case int16:
			return x > int32(y)
		case int32:
			return x > y
		case int64:
			return int64(x) > y
		case float32:
			return float32(x) > y
		case float64:
			return float64(x) > y
		}
	case int64:
		switch y := b.(type) {
		case uint:
			return x > int64(y)
		case uint8:
			return x > int64(y)
		case uint16:
			return x > int64(y)
		case uint32:
			return x > int64(y)
		case uint64:
			return x > int64(y)
		case int:
			return x > int64(y)
		case int8:
			return x > int64(y)
		case int16:
			return x > int64(y)
		case int32:
			return x > int64(y)
		case int64:
			return x > y
		case float32:
			return float32(x) > y
		case float64:
			return float64(x) > y
		}
	case float32:
		switch y := b.(type) {
		case uint:
			return x > float32(y)
		case uint8:
			return x > float32(y)
		case uint16:
			return x > float32(y)
		case uint32:
			return x > float32(y)
		case uint64:
			return x > float32(y)
		case int:
			return x > float32(y)
		case int8:
			return x > float32(y)
		case int16:
			return x > float32(y)
		case int32:
			return x > float32(y)
		case int64:
			return x > float32(y)
		case float32:
			return x > y
		case float64:
			return float64(x) > y
		}
	case float64:
		switch y := b.(type) {
		case uint:
			return x > float64(y)
		case uint8:
			return x > float64(y)
		case uint16:
			return x > float64(y)
		case uint32:
			return x > float64(y)
		case uint64:
			return x > float64(y)
		case int:
			return x > float64(y)
		case int8:
			return x > float64(y)
		case int16:
			return x > float64(y)
		case int32:
			return x > float64(y)
		case int64:
			return x > float64(y)
		case float32:
			return x > float64(y)
		case float64:
			return x > y
		}
	case string:
		switch y := b.(type) {
		case string:
			return x > y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, ">", b))
}

func lessOrEqual(a, b interface{}) interface{} {
	switch x := a.(type) {
	case uint:
		switch y := b.(type) {
		case uint:
			return x <= y
		case uint8:
			return uint8(x) <= y
		case uint16:
			return uint16(x) <= y
		case uint32:
			return uint32(x) <= y
		case uint64:
			return uint64(x) <= y
		case int:
			return int(x) <= y
		case int8:
			return int8(x) <= y
		case int16:
			return int16(x) <= y
		case int32:
			return int32(x) <= y
		case int64:
			return int64(x) <= y
		case float32:
			return float32(x) <= y
		case float64:
			return float64(x) <= y
		}
	case uint8:
		switch y := b.(type) {
		case uint:
			return x <= uint8(y)
		case uint8:
			return x <= y
		case uint16:
			return uint16(x) <= y
		case uint32:
			return uint32(x) <= y
		case uint64:
			return uint64(x) <= y
		case int:
			return int(x) <= y
		case int8:
			return int8(x) <= y
		case int16:
			return int16(x) <= y
		case int32:
			return int32(x) <= y
		case int64:
			return int64(x) <= y
		case float32:
			return float32(x) <= y
		case float64:
			return float64(x) <= y
		}
	case uint16:
		switch y := b.(type) {
		case uint:
			return x <= uint16(y)
		case uint8:
			return x <= uint16(y)
		case uint16:
			return x <= y
		case uint32:
			return uint32(x) <= y
		case uint64:
			return uint64(x) <= y
		case int:
			return int(x) <= y
		case int8:
			return int8(x) <= y
		case int16:
			return int16(x) <= y
		case int32:
			return int32(x) <= y
		case int64:
			return int64(x) <= y
		case float32:
			return float32(x) <= y
		case float64:
			return float64(x) <= y
		}
	case uint32:
		switch y := b.(type) {
		case uint:
			return x <= uint32(y)
		case uint8:
			return x <= uint32(y)
		case uint16:
			return x <= uint32(y)
		case uint32:
			return x <= y
		case uint64:
			return uint64(x) <= y
		case int:
			return int(x) <= y
		case int8:
			return int8(x) <= y
		case int16:
			return int16(x) <= y
		case int32:
			return int32(x) <= y
		case int64:
			return int64(x) <= y
		case float32:
			return float32(x) <= y
		case float64:
			return float64(x) <= y
		}
	case uint64:
		switch y := b.(type) {
		case uint:
			return x <= uint64(y)
		case uint8:
			return x <= uint64(y)
		case uint16:
			return x <= uint64(y)
		case uint32:
			return x <= uint64(y)
		case uint64:
			return x <= y
		case int:
			return int(x) <= y
		case int8:
			return int8(x) <= y
		case int16:
			return int16(x) <= y
		case int32:
			return int32(x) <= y
		case int64:
			return int64(x) <= y
		case float32:
			return float32(x) <= y
		case float64:
			return float64(x) <= y
		}
	case int:
		switch y := b.(type) {
		case uint:
			return x <= int(y)
		case uint8:
			return x <= int(y)
		case uint16:
			return x <= int(y)
		case uint32:
			return x <= int(y)
		case uint64:
			return x <= int(y)
		case int:
			return x <= y
		case int8:
			return int8(x) <= y
		case int16:
			return int16(x) <= y
		case int32:
			return int32(x) <= y
		case int64:
			return int64(x) <= y
		case float32:
			return float32(x) <= y
		case float64:
			return float64(x) <= y
		}
	case int8:
		switch y := b.(type) {
		case uint:
			return x <= int8(y)
		case uint8:
			return x <= int8(y)
		case uint16:
			return x <= int8(y)
		case uint32:
			return x <= int8(y)
		case uint64:
			return x <= int8(y)
		case int:
			return x <= int8(y)
		case int8:
			return x <= y
		case int16:
			return int16(x) <= y
		case int32:
			return int32(x) <= y
		case int64:
			return int64(x) <= y
		case float32:
			return float32(x) <= y
		case float64:
			return float64(x) <= y
		}
	case int16:
		switch y := b.(type) {
		case uint:
			return x <= int16(y)
		case uint8:
			return x <= int16(y)
		case uint16:
			return x <= int16(y)
		case uint32:
			return x <= int16(y)
		case uint64:
			return x <= int16(y)
		case int:
			return x <= int16(y)
		case int8:
			return x <= int16(y)
		case int16:
			return x <= y
		case int32:
			return int32(x) <= y
		case int64:
			return int64(x) <= y
		case float32:
			return float32(x) <= y
		case float64:
			return float64(x) <= y
		}
	case int32:
		switch y := b.(type) {
		case uint:
			return x <= int32(y)
		case uint8:
			return x <= int32(y)
		case uint16:
			return x <= int32(y)
		case uint32:
			return x <= int32(y)
		case uint64:
			return x <= int32(y)
		case int:
			return x <= int32(y)
		case int8:
			return x <= int32(y)
		case int16:
			return x <= int32(y)
		case int32:
			return x <= y
		case int64:
			return int64(x) <= y
		case float32:
			return float32(x) <= y
		case float64:
			return float64(x) <= y
		}
	case int64:
		switch y := b.(type) {
		case uint:
			return x <= int64(y)
		case uint8:
			return x <= int64(y)
		case uint16:
			return x <= int64(y)
		case uint32:
			return x <= int64(y)
		case uint64:
			return x <= int64(y)
		case int:
			return x <= int64(y)
		case int8:
			return x <= int64(y)
		case int16:
			return x <= int64(y)
		case int32:
			return x <= int64(y)
		case int64:
			return x <= y
		case float32:
			return float32(x) <= y
		case float64:
			return float64(x) <= y
		}
	case float32:
		switch y := b.(type) {
		case uint:
			return x <= float32(y)
		case uint8:
			return x <= float32(y)
		case uint16:
			return x <= float32(y)
		case uint32:
			return x <= float32(y)
		case uint64:
			return x <= float32(y)
		case int:
			return x <= float32(y)
		case int8:
			return x <= float32(y)
		case int16:
			return x <= float32(y)
		case int32:
			return x <= float32(y)
		case int64:
			return x <= float32(y)
		case float32:
			return x <= y
		case float64:
			return float64(x) <= y
		}
	case float64:
		switch y := b.(type) {
		case uint:
			return x <= float64(y)
		case uint8:
			return x <= float64(y)
		case uint16:
			return x <= float64(y)
		case uint32:
			return x <= float64(y)
		case uint64:
			return x <= float64(y)
		case int:
			return x <= float64(y)
		case int8:
			return x <= float64(y)
		case int16:
			return x <= float64(y)
		case int32:
			return x <= float64(y)
		case int64:
			return x <= float64(y)
		case float32:
			return x <= float64(y)
		case float64:
			return x <= y
		}
	case string:
		switch y := b.(type) {
		case string:
			return x <= y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "<=", b))
}

func moreOrEqual(a, b interface{}) interface{} {
	switch x := a.(type) {
	case uint:
		switch y := b.(type) {
		case uint:
			return x >= y
		case uint8:
			return uint8(x) >= y
		case uint16:
			return uint16(x) >= y
		case uint32:
			return uint32(x) >= y
		case uint64:
			return uint64(x) >= y
		case int:
			return int(x) >= y
		case int8:
			return int8(x) >= y
		case int16:
			return int16(x) >= y
		case int32:
			return int32(x) >= y
		case int64:
			return int64(x) >= y
		case float32:
			return float32(x) >= y
		case float64:
			return float64(x) >= y
		}
	case uint8:
		switch y := b.(type) {
		case uint:
			return x >= uint8(y)
		case uint8:
			return x >= y
		case uint16:
			return uint16(x) >= y
		case uint32:
			return uint32(x) >= y
		case uint64:
			return uint64(x) >= y
		case int:
			return int(x) >= y
		case int8:
			return int8(x) >= y
		case int16:
			return int16(x) >= y
		case int32:
			return int32(x) >= y
		case int64:
			return int64(x) >= y
		case float32:
			return float32(x) >= y
		case float64:
			return float64(x) >= y
		}
	case uint16:
		switch y := b.(type) {
		case uint:
			return x >= uint16(y)
		case uint8:
			return x >= uint16(y)
		case uint16:
			return x >= y
		case uint32:
			return uint32(x) >= y
		case uint64:
			return uint64(x) >= y
		case int:
			return int(x) >= y
		case int8:
			return int8(x) >= y
		case int16:
			return int16(x) >= y
		case int32:
			return int32(x) >= y
		case int64:
			return int64(x) >= y
		case float32:
			return float32(x) >= y
		case float64:
			return float64(x) >= y
		}
	case uint32:
		switch y := b.(type) {
		case uint:
			return x >= uint32(y)
		case uint8:
			return x >= uint32(y)
		case uint16:
			return x >= uint32(y)
		case uint32:
			return x >= y
		case uint64:
			return uint64(x) >= y
		case int:
			return int(x) >= y
		case int8:
			return int8(x) >= y
		case int16:
			return int16(x) >= y
		case int32:
			return int32(x) >= y
		case int64:
			return int64(x) >= y
		case float32:
			return float32(x) >= y
		case float64:
			return float64(x) >= y
		}
	case uint64:
		switch y := b.(type) {
		case uint:
			return x >= uint64(y)
		case uint8:
			return x >= uint64(y)
		case uint16:
			return x >= uint64(y)
		case uint32:
			return x >= uint64(y)
		case uint64:
			return x >= y
		case int:
			return int(x) >= y
		case int8:
			return int8(x) >= y
		case int16:
			return int16(x) >= y
		case int32:
			return int32(x) >= y
		case int64:
			return int64(x) >= y
		case float32:
			return float32(x) >= y
		case float64:
			return float64(x) >= y
		}
	case int:
		switch y := b.(type) {
		case uint:
			return x >= int(y)
		case uint8:
			return x >= int(y)
		case uint16:
			return x >= int(y)
		case uint32:
			return x >= int(y)
		case uint64:
			return x >= int(y)
		case int:
			return x >= y
		case int8:
			return int8(x) >= y
		case int16:
			return int16(x) >= y
		case int32:
			return int32(x) >= y
		case int64:
			return int64(x) >= y
		case float32:
			return float32(x) >= y
		case float64:
			return float64(x) >= y
		}
	case int8:
		switch y := b.(type) {
		case uint:
			return x >= int8(y)
		case uint8:
			return x >= int8(y)
		case uint16:
			return x >= int8(y)
		case uint32:
			return x >= int8(y)
		case uint64:
			return x >= int8(y)
		case int:
			return x >= int8(y)
		case int8:
			return x >= y
		case int16:
			return int16(x) >= y
		case int32:
			return int32(x) >= y
		case int64:
			return int64(x) >= y
		case float32:
			return float32(x) >= y
		case float64:
			return float64(x) >= y
		}
	case int16:
		switch y := b.(type) {
		case uint:
			return x >= int16(y)
		case uint8:
			return x >= int16(y)
		case uint16:
			return x >= int16(y)
		case uint32:
			return x >= int16(y)
		case uint64:
			return x >= int16(y)
		case int:
			return x >= int16(y)
		case int8:
			return x >= int16(y)
		case int16:
			return x >= y
		case int32:
			return int32(x) >= y
		case int64:
			return int64(x) >= y
		case float32:
			return float32(x) >= y
		case float64:
			return float64(x) >= y
		}
	case int32:
		switch y := b.(type) {
		case uint:
			return x >= int32(y)
		case uint8:
			return x >= int32(y)
		case uint16:
			return x >= int32(y)
		case uint32:
			return x >= int32(y)
		case uint64:
			return x >= int32(y)
		case int:
			return x >= int32(y)
		case int8:
			return x >= int32(y)
		case int16:
			return x >= int32(y)
		case int32:
			return x >= y
		case int64:
			return int64(x) >= y
		case float32:
			return float32(x) >= y
		case float64:
			return float64(x) >= y
		}
	case int64:
		switch y := b.(type) {
		case uint:
			return x >= int64(y)
		case uint8:
			return x >= int64(y)
		case uint16:
			return x >= int64(y)
		case uint32:
			return x >= int64(y)
		case uint64:
			return x >= int64(y)
		case int:
			return x >= int64(y)
		case int8:
			return x >= int64(y)
		case int16:
			return x >= int64(y)
		case int32:
			return x >= int64(y)
		case int64:
			return x >= y
		case float32:
			return float32(x) >= y
		case float64:
			return float64(x) >= y
		}
	case float32:
		switch y := b.(type) {
		case uint:
			return x >= float32(y)
		case uint8:
			return x >= float32(y)
		case uint16:
			return x >= float32(y)
		case uint32:
			return x >= float32(y)
		case uint64:
			return x >= float32(y)
		case int:
			return x >= float32(y)
		case int8:
			return x >= float32(y)
		case int16:
			return x >= float32(y)
		case int32:
			return x >= float32(y)
		case int64:
			return x >= float32(y)
		case float32:
			return x >= y
		case float64:
			return float64(x) >= y
		}
	case float64:
		switch y := b.(type) {
		case uint:
			return x >= float64(y)
		case uint8:
			return x >= float64(y)
		case uint16:
			return x >= float64(y)
		case uint32:
			return x >= float64(y)
		case uint64:
			return x >= float64(y)
		case int:
			return x >= float64(y)
		case int8:
			return x >= float64(y)
		case int16:
			return x >= float64(y)
		case int32:
			return x >= float64(y)
		case int64:
			return x >= float64(y)
		case float32:
			return x >= float64(y)
		case float64:
			return x >= y
		}
	case string:
		switch y := b.(type) {
		case string:
			return x >= y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, ">=", b))
}

func add(a, b interface{}) interface{} {
	switch x := a.(type) {
	case uint:
		switch y := b.(type) {
		case uint:
			return x + y
		case uint8:
			return uint8(x) + y
		case uint16:
			return uint16(x) + y
		case uint32:
			return uint32(x) + y
		case uint64:
			return uint64(x) + y
		case int:
			return int(x) + y
		case int8:
			return int8(x) + y
		case int16:
			return int16(x) + y
		case int32:
			return int32(x) + y
		case int64:
			return int64(x) + y
		case float32:
			return float32(x) + y
		case float64:
			return float64(x) + y
		}
	case uint8:
		switch y := b.(type) {
		case uint:
			return x + uint8(y)
		case uint8:
			return x + y
		case uint16:
			return uint16(x) + y
		case uint32:
			return uint32(x) + y
		case uint64:
			return uint64(x) + y
		case int:
			return int(x) + y
		case int8:
			return int8(x) + y
		case int16:
			return int16(x) + y
		case int32:
			return int32(x) + y
		case int64:
			return int64(x) + y
		case float32:
			return float32(x) + y
		case float64:
			return float64(x) + y
		}
	case uint16:
		switch y := b.(type) {
		case uint:
			return x + uint16(y)
		case uint8:
			return x + uint16(y)
		case uint16:
			return x + y
		case uint32:
			return uint32(x) + y
		case uint64:
			return uint64(x) + y
		case int:
			return int(x) + y
		case int8:
			return int8(x) + y
		case int16:
			return int16(x) + y
		case int32:
			return int32(x) + y
		case int64:
			return int64(x) + y
		case float32:
			return float32(x) + y
		case float64:
			return float64(x) + y
		}
	case uint32:
		switch y := b.(type) {
		case uint:
			return x + uint32(y)
		case uint8:
			return x + uint32(y)
		case uint16:
			return x + uint32(y)
		case uint32:
			return x + y
		case uint64:
			return uint64(x) + y
		case int:
			return int(x) + y
		case int8:
			return int8(x) + y
		case int16:
			return int16(x) + y
		case int32:
			return int32(x) + y
		case int64:
			return int64(x) + y
		case float32:
			return float32(x) + y
		case float64:
			return float64(x) + y
		}
	case uint64:
		switch y := b.(type) {
		case uint:
			return x + uint64(y)
		case uint8:
			return x + uint64(y)
		case uint16:
			return x + uint64(y)
		case uint32:
			return x + uint64(y)
		case uint64:
			return x + y
		case int:
			return int(x) + y
		case int8:
			return int8(x) + y
		case int16:
			return int16(x) + y
		case int32:
			return int32(x) + y
		case int64:
			return int64(x) + y
		case float32:
			return float32(x) + y
		case float64:
			return float64(x) + y
		}
	case int:
		switch y := b.(type) {
		case uint:
			return x + int(y)
		case uint8:
			return x + int(y)
		case uint16:
			return x + int(y)
		case uint32:
			return x + int(y)
		case uint64:
			return x + int(y)
		case int:
			return x + y
		case int8:
			return int8(x) + y
		case int16:
			return int16(x) + y
		case int32:
			return int32(x) + y
		case int64:
			return int64(x) + y
		case float32:
			return float32(x) + y
		case float64:
			return float64(x) + y
		}
	case int8:
		switch y := b.(type) {
		case uint:
			return x + int8(y)
		case uint8:
			return x + int8(y)
		case uint16:
			return x + int8(y)
		case uint32:
			return x + int8(y)
		case uint64:
			return x + int8(y)
		case int:
			return x + int8(y)
		case int8:
			return x + y
		case int16:
			return int16(x) + y
		case int32:
			return int32(x) + y
		case int64:
			return int64(x) + y
		case float32:
			return float32(x) + y
		case float64:
			return float64(x) + y
		}
	case int16:
		switch y := b.(type) {
		case uint:
			return x + int16(y)
		case uint8:
			return x + int16(y)
		case uint16:
			return x + int16(y)
		case uint32:
			return x + int16(y)
		case uint64:
			return x + int16(y)
		case int:
			return x + int16(y)
		case int8:
			return x + int16(y)
		case int16:
			return x + y
		case int32:
			return int32(x) + y
		case int64:
			return int64(x) + y
		case float32:
			return float32(x) + y
		case float64:
			return float64(x) + y
		}
	case int32:
		switch y := b.(type) {
		case uint:
			return x + int32(y)
		case uint8:
			return x + int32(y)
		case uint16:
			return x + int32(y)
		case uint32:
			return x + int32(y)
		case uint64:
			return x + int32(y)
		case int:
			return x + int32(y)
		case int8:
			return x + int32(y)
		case int16:
			return x + int32(y)
		case int32:
			return x + y
		case int64:
			return int64(x) + y
		case float32:
			return float32(x) + y
		case float64:
			return float64(x) + y
		}
	case int64:
		switch y := b.(type) {
		case uint:
			return x + int64(y)
		case uint8:
			return x + int64(y)
		case uint16:
			return x + int64(y)
		case uint32:
			return x + int64(y)
		case uint64:
			return x + int64(y)
		case int:
			return x + int64(y)
		case int8:
			return x + int64(y)
		case int16:
			return x + int64(y)
		case int32:
			return x + int64(y)
		case int64:
			return x + y
		case float32:
			return float32(x) + y
		case float64:
			return float64(x) + y
		}
	case float32:
		switch y := b.(type) {
		case uint:
			return x + float32(y)
		case uint8:
			return x + float32(y)
		case uint16:
			return x + float32(y)
		case uint32:
			return x + float32(y)
		case uint64:
			return x + float32(y)
		case int:
			return x + float32(y)
		case int8:
			return x + float32(y)
		case int16:
			return x + float32(y)
		case int32:
			return x + float32(y)
		case int64:
			return x + float32(y)
		case float32:
			return x + y
		case float64:
			return float64(x) + y
		}
	case float64:
		switch y := b.(type) {
		case uint:
			return x + float64(y)
		case uint8:
			return x + float64(y)
		case uint16:
			return x + float64(y)
		case uint32:
			return x + float64(y)
		case uint64:
			return x + float64(y)
		case int:
			return x + float64(y)
		case int8:
			return x + float64(y)
		case int16:
			return x + float64(y)
		case int32:
			return x + float64(y)
		case int64:
			return x + float64(y)
		case float32:
			return x + float64(y)
		case float64:
			return x + y
		}
	case string:
		switch y := b.(type) {
		case string:
			return x + y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "+", b))
}

func subtract(a, b interface{}) interface{} {
	switch x := a.(type) {
	case uint:
		switch y := b.(type) {
		case uint:
			return x - y
		case uint8:
			return uint8(x) - y
		case uint16:
			return uint16(x) - y
		case uint32:
			return uint32(x) - y
		case uint64:
			return uint64(x) - y
		case int:
			return int(x) - y
		case int8:
			return int8(x) - y
		case int16:
			return int16(x) - y
		case int32:
			return int32(x) - y
		case int64:
			return int64(x) - y
		case float32:
			return float32(x) - y
		case float64:
			return float64(x) - y
		}
	case uint8:
		switch y := b.(type) {
		case uint:
			return x - uint8(y)
		case uint8:
			return x - y
		case uint16:
			return uint16(x) - y
		case uint32:
			return uint32(x) - y
		case uint64:
			return uint64(x) - y
		case int:
			return int(x) - y
		case int8:
			return int8(x) - y
		case int16:
			return int16(x) - y
		case int32:
			return int32(x) - y
		case int64:
			return int64(x) - y
		case float32:
			return float32(x) - y
		case float64:
			return float64(x) - y
		}
	case uint16:
		switch y := b.(type) {
		case uint:
			return x - uint16(y)
		case uint8:
			return x - uint16(y)
		case uint16:
			return x - y
		case uint32:
			return uint32(x) - y
		case uint64:
			return uint64(x) - y
		case int:
			return int(x) - y
		case int8:
			return int8(x) - y
		case int16:
			return int16(x) - y
		case int32:
			return int32(x) - y
		case int64:
			return int64(x) - y
		case float32:
			return float32(x) - y
		case float64:
			return float64(x) - y
		}
	case uint32:
		switch y := b.(type) {
		case uint:
			return x - uint32(y)
		case uint8:
			return x - uint32(y)
		case uint16:
			return x - uint32(y)
		case uint32:
			return x - y
		case uint64:
			return uint64(x) - y
		case int:
			return int(x) - y
		case int8:
			return int8(x) - y
		case int16:
			return int16(x) - y
		case int32:
			return int32(x) - y
		case int64:
			return int64(x) - y
		case float32:
			return float32(x) - y
		case float64:
			return float64(x) - y
		}
	case uint64:
		switch y := b.(type) {
		case uint:
			return x - uint64(y)
		case uint8:
			return x - uint64(y)
		case uint16:
			return x - uint64(y)
		case uint32:
			return x - uint64(y)
		case uint64:
			return x - y
		case int:
			return int(x) - y
		case int8:
			return int8(x) - y
		case int16:
			return int16(x) - y
		case int32:
			return int32(x) - y
		case int64:
			return int64(x) - y
		case float32:
			return float32(x) - y
		case float64:
			return float64(x) - y
		}
	case int:
		switch y := b.(type) {
		case uint:
			return x - int(y)
		case uint8:
			return x - int(y)
		case uint16:
			return x - int(y)
		case uint32:
			return x - int(y)
		case uint64:
			return x - int(y)
		case int:
			return x - y
		case int8:
			return int8(x) - y
		case int16:
			return int16(x) - y
		case int32:
			return int32(x) - y
		case int64:
			return int64(x) - y
		case float32:
			return float32(x) - y
		case float64:
			return float64(x) - y
		}
	case int8:
		switch y := b.(type) {
		case uint:
			return x - int8(y)
		case uint8:
			return x - int8(y)
		case uint16:
			return x - int8(y)
		case uint32:
			return x - int8(y)
		case uint64:
			return x - int8(y)
		case int:
			return x - int8(y)
		case int8:
			return x - y
		case int16:
			return int16(x) - y
		case int32:
			return int32(x) - y
		case int64:
			return int64(x) - y
		case float32:
			return float32(x) - y
		case float64:
			return float64(x) - y
		}
	case int16:
		switch y := b.(type) {
		case uint:
			return x - int16(y)
		case uint8:
			return x - int16(y)
		case uint16:
			return x - int16(y)
		case uint32:
			return x - int16(y)
		case uint64:
			return x - int16(y)
		case int:
			return x - int16(y)
		case int8:
			return x - int16(y)
		case int16:
			return x - y
		case int32:
			return int32(x) - y
		case int64:
			return int64(x) - y
		case float32:
			return float32(x) - y
		case float64:
			return float64(x) - y
		}
	case int32:
		switch y := b.(type) {
		case uint:
			return x - int32(y)
		case uint8:
			return x - int32(y)
		case uint16:
			return x - int32(y)
		case uint32:
			return x - int32(y)
		case uint64:
			return x - int32(y)
		case int:
			return x - int32(y)
		case int8:
			return x - int32(y)
		case int16:
			return x - int32(y)
		case int32:
			return x - y
		case int64:
			return int64(x) - y
		case float32:
			return float32(x) - y
		case float64:
			return float64(x) - y
		}
	case int64:
		switch y := b.(type) {
		case uint:
			return x - int64(y)
		case uint8:
			return x - int64(y)
		case uint16:
			return x - int64(y)
		case uint32:
			return x - int64(y)
		case uint64:
			return x - int64(y)
		case int:
			return x - int64(y)
		case int8:
			return x - int64(y)
		case int16:
			return x - int64(y)
		case int32:
			return x - int64(y)
		case int64:
			return x - y
		case float32:
			return float32(x) - y
		case float64:
			return float64(x) - y
		}
	case float32:
		switch y := b.(type) {
		case uint:
			return x - float32(y)
		case uint8:
			return x - float32(y)
		case uint16:
			return x - float32(y)
		case uint32:
			return x - float32(y)
		case uint64:
			return x - float32(y)
		case int:
			return x - float32(y)
		case int8:
			return x - float32(y)
		case int16:
			return x - float32(y)
		case int32:
			return x - float32(y)
		case int64:
			return x - float32(y)
		case float32:
			return x - y
		case float64:
			return float64(x) - y
		}
	case float64:
		switch y := b.(type) {
		case uint:
			return x - float64(y)
		case uint8:
			return x - float64(y)
		case uint16:
			return x - float64(y)
		case uint32:
			return x - float64(y)
		case uint64:
			return x - float64(y)
		case int:
			return x - float64(y)
		case int8:
			return x - float64(y)
		case int16:
			return x - float64(y)
		case int32:
			return x - float64(y)
		case int64:
			return x - float64(y)
		case float32:
			return x - float64(y)
		case float64:
			return x - y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "-", b))
}

func multiply(a, b interface{}) interface{} {
	switch x := a.(type) {
	case uint:
		switch y := b.(type) {
		case uint:
			return x * y
		case uint8:
			return uint8(x) * y
		case uint16:
			return uint16(x) * y
		case uint32:
			return uint32(x) * y
		case uint64:
			return uint64(x) * y
		case int:
			return int(x) * y
		case int8:
			return int8(x) * y
		case int16:
			return int16(x) * y
		case int32:
			return int32(x) * y
		case int64:
			return int64(x) * y
		case float32:
			return float32(x) * y
		case float64:
			return float64(x) * y
		}
	case uint8:
		switch y := b.(type) {
		case uint:
			return x * uint8(y)
		case uint8:
			return x * y
		case uint16:
			return uint16(x) * y
		case uint32:
			return uint32(x) * y
		case uint64:
			return uint64(x) * y
		case int:
			return int(x) * y
		case int8:
			return int8(x) * y
		case int16:
			return int16(x) * y
		case int32:
			return int32(x) * y
		case int64:
			return int64(x) * y
		case float32:
			return float32(x) * y
		case float64:
			return float64(x) * y
		}
	case uint16:
		switch y := b.(type) {
		case uint:
			return x * uint16(y)
		case uint8:
			return x * uint16(y)
		case uint16:
			return x * y
		case uint32:
			return uint32(x) * y
		case uint64:
			return uint64(x) * y
		case int:
			return int(x) * y
		case int8:
			return int8(x) * y
		case int16:
			return int16(x) * y
		case int32:
			return int32(x) * y
		case int64:
			return int64(x) * y
		case float32:
			return float32(x) * y
		case float64:
			return float64(x) * y
		}
	case uint32:
		switch y := b.(type) {
		case uint:
			return x * uint32(y)
		case uint8:
			return x * uint32(y)
		case uint16:
			return x * uint32(y)
		case uint32:
			return x * y
		case uint64:
			return uint64(x) * y
		case int:
			return int(x) * y
		case int8:
			return int8(x) * y
		case int16:
			return int16(x) * y
		case int32:
			return int32(x) * y
		case int64:
			return int64(x) * y
		case float32:
			return float32(x) * y
		case float64:
			return float64(x) * y
		}
	case uint64:
		switch y := b.(type) {
		case uint:
			return x * uint64(y)
		case uint8:
			return x * uint64(y)
		case uint16:
			return x * uint64(y)
		case uint32:
			return x * uint64(y)
		case uint64:
			return x * y
		case int:
			return int(x) * y
		case int8:
			return int8(x) * y
		case int16:
			return int16(x) * y
		case int32:
			return int32(x) * y
		case int64:
			return int64(x) * y
		case float32:
			return float32(x) * y
		case float64:
			return float64(x) * y
		}
	case int:
		switch y := b.(type) {
		case uint:
			return x * int(y)
		case uint8:
			return x * int(y)
		case uint16:
			return x * int(y)
		case uint32:
			return x * int(y)
		case uint64:
			return x * int(y)
		case int:
			return x * y
		case int8:
			return int8(x) * y
		case int16:
			return int16(x) * y
		case int32:
			return int32(x) * y
		case int64:
			return int64(x) * y
		case float32:
			return float32(x) * y
		case float64:
			return float64(x) * y
		}
	case int8:
		switch y := b.(type) {
		case uint:
			return x * int8(y)
		case uint8:
			return x * int8(y)
		case uint16:
			return x * int8(y)
		case uint32:
			return x * int8(y)
		case uint64:
			return x * int8(y)
		case int:
			return x * int8(y)
		case int8:
			return x * y
		case int16:
			return int16(x) * y
		case int32:
			return int32(x) * y
		case int64:
			return int64(x) * y
		case float32:
			return float32(x) * y
		case float64:
			return float64(x) * y
		}
	case int16:
		switch y := b.(type) {
		case uint:
			return x * int16(y)
		case uint8:
			return x * int16(y)
		case uint16:
			return x * int16(y)
		case uint32:
			return x * int16(y)
		case uint64:
			return x * int16(y)
		case int:
			return x * int16(y)
		case int8:
			return x * int16(y)
		case int16:
			return x * y
		case int32:
			return int32(x) * y
		case int64:
			return int64(x) * y
		case float32:
			return float32(x) * y
		case float64:
			return float64(x) * y
		}
	case int32:
		switch y := b.(type) {
		case uint:
			return x * int32(y)
		case uint8:
			return x * int32(y)
		case uint16:
			return x * int32(y)
		case uint32:
			return x * int32(y)
		case uint64:
			return x * int32(y)
		case int:
			return x * int32(y)
		case int8:
			return x * int32(y)
		case int16:
			return x * int32(y)
		case int32:
			return x * y
		case int64:
			return int64(x) * y
		case float32:
			return float32(x) * y
		case float64:
			return float64(x) * y
		}
	case int64:
		switch y := b.(type) {
		case uint:
			return x * int64(y)
		case uint8:
			return x * int64(y)
		case uint16:
			return x * int64(y)
		case uint32:
			return x * int64(y)
		case uint64:
			return x * int64(y)
		case int:
			return x * int64(y)
		case int8:
			return x * int64(y)
		case int16:
			return x * int64(y)
		case int32:
			return x * int64(y)
		case int64:
			return x * y
		case float32:
			return float32(x) * y
		case float64:
			return float64(x) * y
		}
	case float32:
		switch y := b.(type) {
		case uint:
			return x * float32(y)
		case uint8:
			return x * float32(y)
		case uint16:
			return x * float32(y)
		case uint32:
			return x * float32(y)
		case uint64:
			return x * float32(y)
		case int:
			return x * float32(y)
		case int8:
			return x * float32(y)
		case int16:
			return x * float32(y)
		case int32:
			return x * float32(y)
		case int64:
			return x * float32(y)
		case float32:
			return x * y
		case float64:
			return float64(x) * y
		}
	case float64:
		switch y := b.(type) {
		case uint:
			return x * float64(y)
		case uint8:
			return x * float64(y)
		case uint16:
			return x * float64(y)
		case uint32:
			return x * float64(y)
		case uint64:
			return x * float64(y)
		case int:
			return x * float64(y)
		case int8:
			return x * float64(y)
		case int16:
			return x * float64(y)
		case int32:
			return x * float64(y)
		case int64:
			return x * float64(y)
		case float32:
			return x * float64(y)
		case float64:
			return x * y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "*", b))
}

func divide(a, b interface{}) interface{} {
	switch x := a.(type) {
	case uint:
		switch y := b.(type) {
		case uint:
			return x / y
		case uint8:
			return uint8(x) / y
		case uint16:
			return uint16(x) / y
		case uint32:
			return uint32(x) / y
		case uint64:
			return uint64(x) / y
		case int:
			return int(x) / y
		case int8:
			return int8(x) / y
		case int16:
			return int16(x) / y
		case int32:
			return int32(x) / y
		case int64:
			return int64(x) / y
		case float32:
			return float32(x) / y
		case float64:
			return float64(x) / y
		}
	case uint8:
		switch y := b.(type) {
		case uint:
			return x / uint8(y)
		case uint8:
			return x / y
		case uint16:
			return uint16(x) / y
		case uint32:
			return uint32(x) / y
		case uint64:
			return uint64(x) / y
		case int:
			return int(x) / y
		case int8:
			return int8(x) / y
		case int16:
			return int16(x) / y
		case int32:
			return int32(x) / y
		case int64:
			return int64(x) / y
		case float32:
			return float32(x) / y
		case float64:
			return float64(x) / y
		}
	case uint16:
		switch y := b.(type) {
		case uint:
			return x / uint16(y)
		case uint8:
			return x / uint16(y)
		case uint16:
			return x / y
		case uint32:
			return uint32(x) / y
		case uint64:
			return uint64(x) / y
		case int:
			return int(x) / y
		case int8:
			return int8(x) / y
		case int16:
			return int16(x) / y
		case int32:
			return int32(x) / y
		case int64:
			return int64(x) / y
		case float32:
			return float32(x) / y
		case float64:
			return float64(x) / y
		}
	case uint32:
		switch y := b.(type) {
		case uint:
			return x / uint32(y)
		case uint8:
			return x / uint32(y)
		case uint16:
			return x / uint32(y)
		case uint32:
			return x / y
		case uint64:
			return uint64(x) / y
		case int:
			return int(x) / y
		case int8:
			return int8(x) / y
		case int16:
			return int16(x) / y
		case int32:
			return int32(x) / y
		case int64:
			return int64(x) / y
		case float32:
			return float32(x) / y
		case float64:
			return float64(x) / y
		}
	case uint64:
		switch y := b.(type) {
		case uint:
			return x / uint64(y)
		case uint8:
			return x / uint64(y)
		case uint16:
			return x / uint64(y)
		case uint32:
			return x / uint64(y)
		case uint64:
			return x / y
		case int:
			return int(x) / y
		case int8:
			return int8(x) / y
		case int16:
			return int16(x) / y
		case int32:
			return int32(x) / y
		case int64:
			return int64(x) / y
		case float32:
			return float32(x) / y
		case float64:
			return float64(x) / y
		}
	case int:
		switch y := b.(type) {
		case uint:
			return x / int(y)
		case uint8:
			return x / int(y)
		case uint16:
			return x / int(y)
		case uint32:
			return x / int(y)
		case uint64:
			return x / int(y)
		case int:
			return x / y
		case int8:
			return int8(x) / y
		case int16:
			return int16(x) / y
		case int32:
			return int32(x) / y
		case int64:
			return int64(x) / y
		case float32:
			return float32(x) / y
		case float64:
			return float64(x) / y
		}
	case int8:
		switch y := b.(type) {
		case uint:
			return x / int8(y)
		case uint8:
			return x / int8(y)
		case uint16:
			return x / int8(y)
		case uint32:
			return x / int8(y)
		case uint64:
			return x / int8(y)
		case int:
			return x / int8(y)
		case int8:
			return x / y
		case int16:
			return int16(x) / y
		case int32:
			return int32(x) / y
		case int64:
			return int64(x) / y
		case float32:
			return float32(x) / y
		case float64:
			return float64(x) / y
		}
	case int16:
		switch y := b.(type) {
		case uint:
			return x / int16(y)
		case uint8:
			return x / int16(y)
		case uint16:
			return x / int16(y)
		case uint32:
			return x / int16(y)
		case uint64:
			return x / int16(y)
		case int:
			return x / int16(y)
		case int8:
			return x / int16(y)
		case int16:
			return x / y
		case int32:
			return int32(x) / y
		case int64:
			return int64(x) / y
		case float32:
			return float32(x) / y
		case float64:
			return float64(x) / y
		}
	case int32:
		switch y := b.(type) {
		case uint:
			return x / int32(y)
		case uint8:
			return x / int32(y)
		case uint16:
			return x / int32(y)
		case uint32:
			return x / int32(y)
		case uint64:
			return x / int32(y)
		case int:
			return x / int32(y)
		case int8:
			return x / int32(y)
		case int16:
			return x / int32(y)
		case int32:
			return x / y
		case int64:
			return int64(x) / y
		case float32:
			return float32(x) / y
		case float64:
			return float64(x) / y
		}
	case int64:
		switch y := b.(type) {
		case uint:
			return x / int64(y)
		case uint8:
			return x / int64(y)
		case uint16:
			return x / int64(y)
		case uint32:
			return x / int64(y)
		case uint64:
			return x / int64(y)
		case int:
			return x / int64(y)
		case int8:
			return x / int64(y)
		case int16:
			return x / int64(y)
		case int32:
			return x / int64(y)
		case int64:
			return x / y
		case float32:
			return float32(x) / y
		case float64:
			return float64(x) / y
		}
	case float32:
		switch y := b.(type) {
		case uint:
			return x / float32(y)
		case uint8:
			return x / float32(y)
		case uint16:
			return x / float32(y)
		case uint32:
			return x / float32(y)
		case uint64:
			return x / float32(y)
		case int:
			return x / float32(y)
		case int8:
			return x / float32(y)
		case int16:
			return x / float32(y)
		case int32:
			return x / float32(y)
		case int64:
			return x / float32(y)
		case float32:
			return x / y
		case float64:
			return float64(x) / y
		}
	case float64:
		switch y := b.(type) {
		case uint:
			return x / float64(y)
		case uint8:
			return x / float64(y)
		case uint16:
			return x / float64(y)
		case uint32:
			return x / float64(y)
		case uint64:
			return x / float64(y)
		case int:
			return x / float64(y)
		case int8:
			return x / float64(y)
		case int16:
			return x / float64(y)
		case int32:
			return x / float64(y)
		case int64:
			return x / float64(y)
		case float32:
			return x / float64(y)
		case float64:
			return x / y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "/", b))
}

func modulo(a, b interface{}) interface{} {
	switch x := a.(type) {
	case uint:
		switch y := b.(type) {
		case uint:
			return x % y
		case uint8:
			return uint8(x) % y
		case uint16:
			return uint16(x) % y
		case uint32:
			return uint32(x) % y
		case uint64:
			return uint64(x) % y
		case int:
			return int(x) % y
		case int8:
			return int8(x) % y
		case int16:
			return int16(x) % y
		case int32:
			return int32(x) % y
		case int64:
			return int64(x) % y
		}
	case uint8:
		switch y := b.(type) {
		case uint:
			return x % uint8(y)
		case uint8:
			return x % y
		case uint16:
			return uint16(x) % y
		case uint32:
			return uint32(x) % y
		case uint64:
			return uint64(x) % y
		case int:
			return int(x) % y
		case int8:
			return int8(x) % y
		case int16:
			return int16(x) % y
		case int32:
			return int32(x) % y
		case int64:
			return int64(x) % y
		}
	case uint16:
		switch y := b.(type) {
		case uint:
			return x % uint16(y)
		case uint8:
			return x % uint16(y)
		case uint16:
			return x % y
		case uint32:
			return uint32(x) % y
		case uint64:
			return uint64(x) % y
		case int:
			return int(x) % y
		case int8:
			return int8(x) % y
		case int16:
			return int16(x) % y
		case int32:
			return int32(x) % y
		case int64:
			return int64(x) % y
		}
	case uint32:
		switch y := b.(type) {
		case uint:
			return x % uint32(y)
		case uint8:
			return x % uint32(y)
		case uint16:
			return x % uint32(y)
		case uint32:
			return x % y
		case uint64:
			return uint64(x) % y
		case int:
			return int(x) % y
		case int8:
			return int8(x) % y
		case int16:
			return int16(x) % y
		case int32:
			return int32(x) % y
		case int64:
			return int64(x) % y
		}
	case uint64:
		switch y := b.(type) {
		case uint:
			return x % uint64(y)
		case uint8:
			return x % uint64(y)
		case uint16:
			return x % uint64(y)
		case uint32:
			return x % uint64(y)
		case uint64:
			return x % y
		case int:
			return int(x) % y
		case int8:
			return int8(x) % y
		case int16:
			return int16(x) % y
		case int32:
			return int32(x) % y
		case int64:
			return int64(x) % y
		}
	case int:
		switch y := b.(type) {
		case uint:
			return x % int(y)
		case uint8:
			return x % int(y)
		case uint16:
			return x % int(y)
		case uint32:
			return x % int(y)
		case uint64:
			return x % int(y)
		case int:
			return x % y
		case int8:
			return int8(x) % y
		case int16:
			return int16(x) % y
		case int32:
			return int32(x) % y
		case int64:
			return int64(x) % y
		}
	case int8:
		switch y := b.(type) {
		case uint:
			return x % int8(y)
		case uint8:
			return x % int8(y)
		case uint16:
			return x % int8(y)
		case uint32:
			return x % int8(y)
		case uint64:
			return x % int8(y)
		case int:
			return x % int8(y)
		case int8:
			return x % y
		case int16:
			return int16(x) % y
		case int32:
			return int32(x) % y
		case int64:
			return int64(x) % y
		}
	case int16:
		switch y := b.(type) {
		case uint:
			return x % int16(y)
		case uint8:
			return x % int16(y)
		case uint16:
			return x % int16(y)
		case uint32:
			return x % int16(y)
		case uint64:
			return x % int16(y)
		case int:
			return x % int16(y)
		case int8:
			return x % int16(y)
		case int16:
			return x % y
		case int32:
			return int32(x) % y
		case int64:
			return int64(x) % y
		}
	case int32:
		switch y := b.(type) {
		case uint:
			return x % int32(y)
		case uint8:
			return x % int32(y)
		case uint16:
			return x % int32(y)
		case uint32:
			return x % int32(y)
		case uint64:
			return x % int32(y)
		case int:
			return x % int32(y)
		case int8:
			return x % int32(y)
		case int16:
			return x % int32(y)
		case int32:
			return x % y
		case int64:
			return int64(x) % y
		}
	case int64:
		switch y := b.(type) {
		case uint:
			return x % int64(y)
		case uint8:
			return x % int64(y)
		case uint16:
			return x % int64(y)
		case uint32:
			return x % int64(y)
		case uint64:
			return x % int64(y)
		case int:
			return x % int64(y)
		case int8:
			return x % int64(y)
		case int16:
			return x % int64(y)
		case int32:
			return x % int64(y)
		case int64:
			return x % y
		}
	}
	panic(fmt.Sprintf("invalid operation: %T %v %T", a, "%", b))
}

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}
	r := reflect.ValueOf(v)
	switch r.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return r.IsNil()
	default:
		return false
	}
}
