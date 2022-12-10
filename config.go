package goconfig

import (
	"encoding/json"
	"os"
)

type ConfigManager struct {
	path string
}

// Creates new config manager for given local directory path, also creates directories specified in path if necessary
func Manage(path string) (ConfigManager, error) {
	var cm ConfigManager
	cm.path = path
	err := os.MkdirAll(cm.path, 0750)
	if err != nil && !os.IsExist(err) {
		return cm, err
	}
	return cm, nil
}

// Writes provided config to specified file
func (cm ConfigManager) Write(name string, config any) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(cm.path+"/"+name+".json", data, 0664)
}

// Reads config from specified file
func (cm ConfigManager) Read(name string, config any) error {
	data, err := os.ReadFile(cm.path + "/" + name + ".json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &config)
}
