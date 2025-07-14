package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	testFunc := func() (string, error) {
		return "hello", nil
	}
	assert.Equal(t, Must(testFunc()), "hello")
}

func TestMustPanicWithError(t *testing.T) {
	testFunc := func() (string, error) {
		return "", fmt.Errorf("an error occurred")
	}
	assert.PanicsWithError(t, "an error occurred", func() {
		Must(testFunc())
	})
}
