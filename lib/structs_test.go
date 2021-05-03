//Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeps(t *testing.T) {
	deps := Deps{}

	var i interface{} = NewDeps()

	assert.IsType(t, deps, i, "NewDeps is not returning a Deps struct")
}
