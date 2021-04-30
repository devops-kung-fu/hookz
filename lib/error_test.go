package lib

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
	"testing"
)

func TestIsError(t *testing.T) {
	output := captureOutput(func() {
		IsError(errors.New("Test Error"), "[TEST]")
	})

	if output.Len() == 0 {
		t.Error("No information logged to STDOUT")
	}
	if strings.Count(output.String(), "\n") > 1 {
		t.Error("Expected only a single line of log output")
	}
}

func TestIsErrorBool(t *testing.T) {
	output := captureOutput(func() {
		IsErrorBool(errors.New("Test Error"), "[TEST]")
	})

	if output.Len() == 0 {
		t.Error("No information logged to STDOUT")
	}
	if strings.Count(output.String(), "\n") > 1 {
		t.Error("Expected only a single line of log output")
	}
}

func TestIfErrorLog(t *testing.T) {
	output := captureOutput(func() {
		IfErrorLog(errors.New("Test Error"), "[TEST]")
	})

	if output.Len() == 0 {
		t.Error("No information logged to STDOUT")
	}
	if strings.Count(output.String(), "\n") > 1 {
		t.Error("Expected only a single line of log output")
	}
}

func captureOutput(f func()) bytes.Buffer {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf
}
