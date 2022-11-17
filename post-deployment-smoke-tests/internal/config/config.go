package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "post-deployment-smoke-tests/types"
    "post-deployment-smoke-tests/internal/errors"
)

func Unmarshal(yamlFile []byte, config types.HttpConfig) (types.HttpConfig, error) {
    err := yaml.Unmarshal(yamlFile, &config)

	if err != nil {
        return config, err
    }

    return config, nil
}

func ReadFile(filename string) (yamlFile []byte) {
	yamlFile, err := ioutil.ReadFile(filename)
	errors.CheckError(err)
    return yamlFile
}

func GetHttpConfig() types.HttpConfig{
    yamlFile := ReadFile(types.HttpConfigName)
    var config types.HttpConfig
    config, err := Unmarshal(yamlFile, config)
    errors.CheckError(err)
	return config
}
