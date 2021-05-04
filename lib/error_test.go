package lib

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsError(t *testing.T) {
	output := captureOutput(func() {
		IsError(errors.New("Test Error"), "[TEST]")
	})

	if output.Len() == 0 {
		assert.GreaterOrEqual(t, output.Len(), 0, "No information logged to STDOUT")
		assert.GreaterOrEqual(t, strings.Count(output.String(), "\n"), 1, "Expected only a single line of log output")
	}
}

func TestIsErrorBool(t *testing.T) {
	output := captureOutput(func() {
		IsErrorBool(errors.New("Test Error"), "[TEST]")
	})

	assert.GreaterOrEqual(t, output.Len(), 0, "No information logged to STDOUT")
	assert.GreaterOrEqual(t, strings.Count(output.String(), "\n"), 1, "Expected only a single line of log output")
}

func TestIfErrorLog(t *testing.T) {
	output := captureOutput(func() {
		IfErrorLog(errors.New("Test Error"), "[TEST]")
	})

	assert.GreaterOrEqual(t, output.Len(), 0, "No information logged to STDOUT")
	assert.GreaterOrEqual(t, strings.Count(output.String(), "\n"), 1, "Expected only a single line of log output")
}

func captureOutput(f func()) bytes.Buffer {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf
}
