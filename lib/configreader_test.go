//Package lib Functionality for the Hookz CLI
package lib

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestDeps_ReadConfig(t *testing.T) {
	deps := Deps{
		fs: afero.NewMemMapFs(),
	}
	version := "1.0.0"

	config, err := createTestConfig(deps, version)
	assert.NoError(t, err, "creteTestConfig error should be nil")
	assert.Equal(t, version, config.Version, "Versions should match")

	// type args struct {
	// 	version string
	// }

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		d := Deps{
	// 			fs: tt.fields.fs,
	// 		}
	// 		gotConfig, err := d.ReadConfig(tt.args.version)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("Deps.ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
	// 			t.Errorf("Deps.ReadConfig() = %v, want %v", gotConfig, tt.wantConfig)
	// 		}
	// 	})
	// }
}

func TestDeps_checkVersion(t *testing.T) {

}

func createTestConfig(d Deps, version string) (config Configuration, err error) {
	command := "echo"
	config = Configuration{
		Version: version,
		Hooks: []Hook{
			{
				Type: "pre-commit",
				Actions: []Action{
					{
						Name: "Hello Hookz!",
						Exec: &command,
						Args: []string{"-e", "Hello Hookz!"},
					},
				},
			},
		},
	}

	file, merr := yaml.Marshal(config)
	if merr != nil {
		err = merr
		fmt.Println(err)
		return
	}

	err = d.Afero().WriteFile(".hooks.yaml", file, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
