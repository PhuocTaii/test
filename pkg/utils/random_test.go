package utils

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	length := 10
	randomString := GenerateRandomString(length)

	if len(randomString) != length*2 {
		t.Errorf("Expected length %d, but got %d", length*2, len(randomString))
	}
}
