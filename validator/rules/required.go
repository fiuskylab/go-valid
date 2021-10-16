package validator

const (
	fieldRequired = `field required`
)

func Required(v interface{}) string {
	switch v.(type) {
	case string:
		if v.(string) == "" {
			return fieldRequired
		}
	case []string:
		if len(v.([]string)) == 0 {
			return fieldRequired
		}
	case []int:
		if len(v.([]int)) == 0 {
			return fieldRequired
		}
	case []int64:
		if len(v.([]int64)) == 0 {
			return fieldRequired
		}
	case []int32:
		if len(v.([]int32)) == 0 {
			return fieldRequired
		}
	case []int16:
		if len(v.([]int16)) == 0 {
			return fieldRequired
		}
	case []int8:
		if len(v.([]int8)) == 0 {
			return fieldRequired
		}
	case []uintptr:
		if len(v.([]uintptr)) == 0 {
			return fieldRequired
		}
	case []uint:
		if len(v.([]uint)) == 0 {
			return fieldRequired
		}
	case []uint64:
		if len(v.([]uint64)) == 0 {
			return fieldRequired
		}
	case []uint32:
		if len(v.([]uint32)) == 0 {
			return fieldRequired
		}
	case []uint16:
		if len(v.([]uint16)) == 0 {
			return fieldRequired
		}
	case []uint8:
		if len(v.([]uint8)) == 0 {
			return fieldRequired
		}
	case []float32:
		if len(v.([]float32)) == 0 {
			return fieldRequired
		}
	case []float64:
		if len(v.([]float64)) == 0 {
			return fieldRequired
		}
	case []complex64:
		if len(v.([]complex64)) == 0 {
			return fieldRequired
		}
	case []complex128:
		if len(v.([]complex128)) == 0 {
			return fieldRequired
		}
	case []interface{}:
		if len(v.([]interface{})) == 0 {
			return fieldRequired
		}
	}
	return ""
}
