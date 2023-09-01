package conf

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Server   ServerConf   `yaml:"server"`
	Database DatabaseConf `yaml:"database"`
	Common   CommonConf   `yaml:"common"`
}

type ServerConf struct {
	Port string `yaml:"port"`
}

type DatabaseConf struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

type CommonConf struct {
	ENV          string `yaml:"ENV"`
	PROJECT_ROOT string `yaml:"PROJECT_ROOT"`
	UPLOADS_PATH string `yaml:"UPLOADS_PATH"`
}

var Settings Conf

func init() {
	temp, err := Readfile("./conf/config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	_, filename, _, _ := runtime.Caller(0)
	temp.Common.PROJECT_ROOT = path.Dir(path.Dir(filename))
	if temp.Common.UPLOADS_PATH[0:1] == "/" {
		temp.Common.UPLOADS_PATH = temp.Common.PROJECT_ROOT + temp.Common.UPLOADS_PATH
	} else {
		temp.Common.UPLOADS_PATH = temp.Common.PROJECT_ROOT + "/" + temp.Common.UPLOADS_PATH
	}
	Settings = temp
}

func Readfile(filePath string) (Conf, error) {
	var settings Conf
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return settings, err
	}
	//yaml文件内容影射到结构体中
	err1 := yaml.Unmarshal(yamlFile, &settings)
	if err1 != nil {
		return settings, err1
	}
	return settings, nil
}
