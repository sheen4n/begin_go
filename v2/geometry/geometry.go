package geometry

import "errors"

// Cube Exported
func Cube(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	}
	return 0, errors.New("Zero length edge is not allowed")
}
