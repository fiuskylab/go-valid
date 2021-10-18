package rules

import "fmt"

const (
	fieldLengthMax = `field lenght is more than %d`
	fieldValueMax  = `field is more than %d`
)

func Max(v interface{}, max int) string {
	switch v.(type) {
	case int:
		if v.(int) > max {
			return fmt.Sprintf(fieldValueMax, max)
		}
	case int64:
		if v.(int64) > int64(max) {
			return fmt.Sprintf(fieldValueMax, max)
		}
	case int32:
		if v.(int32) > int32(max) {
			return fmt.Sprintf(fieldValueMax, max)
		}
	case int16:
		if v.(int16) > int16(max) {
			return fmt.Sprintf(fieldValueMax, max)
		}
	case int8:
		if v.(int8) > int8(max) {
			return fmt.Sprintf(fieldValueMax, max)
		}
	case float64:
		if v.(float64) > float64(max) {
			return fmt.Sprintf(fieldValueMax, max)
		}
	case float32:
		if v.(float32) > float32(max) {
			return fmt.Sprintf(fieldValueMax, max)
		}
	case string:
		if l := len(v.(string)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []string:
		if l := len(v.([]string)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []int:
		if l := len(v.([]int)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []int64:
		if l := len(v.([]int64)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []int32:
		if l := len(v.([]int32)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []int16:
		if l := len(v.([]int16)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []int8:
		if l := len(v.([]int8)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []uintptr:
		if l := len(v.([]uintptr)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []uint:
		if l := len(v.([]uint)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []uint64:
		if l := len(v.([]uint64)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []uint32:
		if l := len(v.([]uint32)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []uint16:
		if l := len(v.([]uint16)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []uint8:
		if l := len(v.([]uint8)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []float32:
		if l := len(v.([]float32)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []float64:
		if l := len(v.([]float64)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []complex64:
		if l := len(v.([]complex64)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []complex128:
		if l := len(v.([]complex128)); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	case []interface{}:
		if l := len(v.([]interface{})); l > max {
			return fmt.Sprintf(fieldLengthMax, max)
		}
	}
	return ""
}
