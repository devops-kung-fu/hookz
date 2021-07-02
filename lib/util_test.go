package lib

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintIf(t *testing.T) {
	result, err := CaptureStdout(func() { fmt.Println("Test") })

	assert.Equal(t, "Test\n", result, "Should match the string Test")
	assert.NoError(t, err, "No error should have been generated")
}

func CaptureStdout(f func()) (captured string, err error) {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	origStdout := os.Stdout
	os.Stdout = w

	f()

	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = origStdout
	captured = string(buf[:n])
	return
}
