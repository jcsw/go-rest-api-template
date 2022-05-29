package system

import (
	ioutil "io/ioutil"
	os "os"

	yaml "github.com/go-yaml/yaml"
)

// AppProperties define the properties values
type AppProperties struct {
	ServerPort int    `yaml:"server.port"`
	Mariadb    string `yaml:"mariadb"`
	Mongodb    string `yaml:"mongodb"`
}

// Properties the loaded properties values
var Properties AppProperties

// LoadProperties load properties in by environment
func LoadProperties(env string) {
	Info("[Loading properties by env %s]", env)

	pwd, _ := os.Getwd()
	file, err := ioutil.ReadFile(pwd + "/properties/" + env + ".yaml")
	if err != nil {
		Fatal("[Could not load file of properties] err:%v", err)
	}

	err = yaml.UnmarshalStrict(file, &Properties)
	if err != nil {
		Fatal("[Could not load properties values] err:%v", err)
	}

	Info("[Properties loaded with successful]")
}
