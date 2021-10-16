package validator

import "fmt"

const (
	fieldMin = `field lenght is less than %d`
)

func Min(v interface{}, min int) string {
	switch v.(type) {
	case string:
		if l := len(v.(string)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []string:
		if l := len(v.([]string)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []int:
		if l := len(v.([]int)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []int64:
		if l := len(v.([]int64)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []int32:
		if l := len(v.([]int32)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []int16:
		if l := len(v.([]int16)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []int8:
		if l := len(v.([]int8)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []uintptr:
		if l := len(v.([]uintptr)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []uint:
		if l := len(v.([]uint)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []uint64:
		if l := len(v.([]uint64)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []uint32:
		if l := len(v.([]uint32)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []uint16:
		if l := len(v.([]uint16)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []uint8:
		if l := len(v.([]uint8)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []float32:
		if l := len(v.([]float32)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []float64:
		if l := len(v.([]float64)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []complex64:
		if l := len(v.([]complex64)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []complex128:
		if l := len(v.([]complex128)); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	case []interface{}:
		if l := len(v.([]interface{})); l < min {
			return fmt.Sprintf(fieldMin, min)
		}
	}
	return ""
}
