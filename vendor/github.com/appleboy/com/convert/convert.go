package convert

import (
	"fmt"
	"math"
	"strconv"
)

// ToString convert any type to string
func ToString(value interface{}) interface{} {
	if v, ok := value.(*string); ok {
		return *v
	}
	return fmt.Sprintf("%v", value)
}

// ToBool convert any type to boolean
func ToBool(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		return value
	case *bool:
		return *value
	case string:
		switch value {
		case "", "false":
			return false
		}
		return true
	case *string:
		return ToBool(*value)
	case float64:
		if value != 0 {
			return true
		}
		return false
	case *float64:
		return ToBool(*value)
	case float32:
		if value != 0 {
			return true
		}
		return false
	case *float32:
		return ToBool(*value)
	case int:
		if value != 0 {
			return true
		}
		return false
	case *int:
		return ToBool(*value)
	}
	return false
}

// ToInt convert any type to int
func ToInt(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return 1
		}
		return 0
	case int:
		if value < int(math.MinInt32) || value > int(math.MaxInt32) {
			return nil
		}
		return value
	case *int:
		return ToInt(*value)
	case int8:
		return int(value)
	case *int8:
		return int(*value)
	case int16:
		return int(value)
	case *int16:
		return int(*value)
	case int32:
		return int(value)
	case *int32:
		return int(*value)
	case int64:
		if value < int64(math.MinInt32) || value > int64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case *int64:
		return ToInt(*value)
	case uint:
		if value > math.MaxInt32 {
			return nil
		}
		return int(value)
	case *uint:
		return ToInt(*value)
	case uint8:
		return int(value)
	case *uint8:
		return int(*value)
	case uint16:
		return int(value)
	case *uint16:
		return int(*value)
	case uint32:
		if value > uint32(math.MaxInt32) {
			return nil
		}
		return int(value)
	case *uint32:
		return ToInt(*value)
	case uint64:
		if value > uint64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case *uint64:
		return ToInt(*value)
	case float32:
		if value < float32(math.MinInt32) || value > float32(math.MaxInt32) {
			return nil
		}
		return int(value)
	case *float32:
		return ToInt(*value)
	case float64:
		if value < float64(math.MinInt32) || value > float64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case *float64:
		return ToInt(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return ToInt(val)
	case *string:
		return ToInt(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

// ToFloat convert any type to float
func ToFloat(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return 1.0
		}
		return 0.0
	case *bool:
		return ToFloat(*value)
	case int:
		return float64(value)
	case *int32:
		return ToFloat(*value)
	case float32:
		return value
	case *float32:
		return ToFloat(*value)
	case float64:
		return value
	case *float64:
		return ToFloat(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return val
	case *string:
		return ToFloat(*value)
	}
	return 0.0
}
