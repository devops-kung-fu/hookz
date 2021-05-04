//Package lib Functionality for the Hookz CLI
package lib

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNewDeps(t *testing.T) {
	f := FileSystem{}

	var i interface{} = NewDeps()
	var fs interface{} = afero.NewOsFs()

	assert.IsType(t, f, i, "NewDeps is not returning a Deps struct")
	assert.IsType(t, fs, NewDeps().fs, "fs should be an afero.OsFs")
}
