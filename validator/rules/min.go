package rules

import "fmt"

const (
	fieldLengthMin = `field lenght is less than %d`
	fieldValueMin  = `field is less than %d`
)

func Min(v interface{}, min int) string {
	switch v.(type) {
	case int:
		if v.(int) < min {
			return fmt.Sprintf(fieldValueMin, min)
		}
	case int64:
		if v.(int64) < int64(min) {
			return fmt.Sprintf(fieldValueMin, min)
		}
	case int32:
		if v.(int32) < int32(min) {
			return fmt.Sprintf(fieldValueMin, min)
		}
	case int16:
		if v.(int16) < int16(min) {
			return fmt.Sprintf(fieldValueMin, min)
		}
	case int8:
		if v.(int8) < int8(min) {
			return fmt.Sprintf(fieldValueMin, min)
		}
	case float64:
		if v.(float64) < float64(min) {
			return fmt.Sprintf(fieldValueMin, min)
		}
	case float32:
		if v.(float32) < float32(min) {
			return fmt.Sprintf(fieldValueMin, min)
		}
	case string:
		if l := len(v.(string)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []string:
		if l := len(v.([]string)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []int:
		if l := len(v.([]int)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []int64:
		if l := len(v.([]int64)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []int32:
		if l := len(v.([]int32)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []int16:
		if l := len(v.([]int16)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []int8:
		if l := len(v.([]int8)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []uintptr:
		if l := len(v.([]uintptr)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []uint:
		if l := len(v.([]uint)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []uint64:
		if l := len(v.([]uint64)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []uint32:
		if l := len(v.([]uint32)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []uint16:
		if l := len(v.([]uint16)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []uint8:
		if l := len(v.([]uint8)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []float32:
		if l := len(v.([]float32)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []float64:
		if l := len(v.([]float64)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []complex64:
		if l := len(v.([]complex64)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []complex128:
		if l := len(v.([]complex128)); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	case []interface{}:
		if l := len(v.([]interface{})); l < min {
			return fmt.Sprintf(fieldLengthMin, min)
		}
	}
	return ""
}
