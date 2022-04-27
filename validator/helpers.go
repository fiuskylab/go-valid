package validator

// ToPtr converts any value
// to a pointer.
func ToPtr[T any](v T) *T {
	return &v
}
