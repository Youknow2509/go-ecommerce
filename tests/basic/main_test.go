package basic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Cách tạo file out: go test -coverprofile=coverprofile.out
 * Chuyển file out thành file html: go tool cover -html=coverprofile.out -o coverprofile.html
 */

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 2
	// )

	// if result := AddOne(input); result != output {
	// 	t.Errorf("AddOne(%d) = %d, expected %d", input, result, output)
	// }
	// t.Log("All tests passed")

	assert.Equal(t, AddOne(2), 3, "AddOne(2) should return 3")
}
