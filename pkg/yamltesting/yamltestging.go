package yamltesting

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config *struct {
	NewVersion string `yaml:"new_version"`
	Oldversion string `yaml:"old_version"`
}

func Yamltesting() {
	something, err := versionCall()
	fmt.Println(something.Oldversion, something.NewVersion, err)
}

func versionCall() (Config, error) {
	data, err := os.ReadFile("testing.yaml")
	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
