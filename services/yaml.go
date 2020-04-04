package services

import (
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
)

type GitConf struct {
	Client_id string `yaml:client_id`
	Client_secret string `yaml:client_secret`
	Url string `yaml:url`
}

func GetGits() (setting *GitConf, e error) {
    config, err := ioutil.ReadFile("./config/github.yaml")
    if err != nil {
        return nil, err
    }
	err = yaml.Unmarshal(config, &setting)
    if err != nil {
        return nil, err
    }
	return setting,nil
}
