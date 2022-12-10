package goconfig

import (
	"os"
	"testing"
)

type DemoConfig1 struct {
	Key1 string
	Key2 string
	Key3 string
}

func TestManage(t *testing.T) {
	os.RemoveAll("cm_test")
	_, err0 := Manage("cm_test")
	if err0 != nil {
		t.Error("Managing not existing config dir failed: " + err0.Error())
	}
	_, err1 := Manage("cm_test")
	if err1 != nil {
		t.Error("Managing existing config dir failed: " + err1.Error())
	}
	_, err2 := Manage("cm_test/sub")
	if err2 != nil {
		t.Error("Managing not existing config sub dir failed: " + err2.Error())
	}
	_, err3 := Manage("cm_test/sub")
	if err3 != nil {
		t.Error("Managing existing config sub dir failed: " + err3.Error())
	}
	os.RemoveAll("cm_test")
	_, err4 := Manage("cm_test/sub")
	if err4 != nil {
		t.Error("Managing not existing config sub dir in not existing dir failed: " + err4.Error())
	}
	os.RemoveAll("cm_test")
}

func TestWrite(t *testing.T) {
	os.RemoveAll("cm_test")
	cm, err0 := Manage("cm_test")
	if err0 != nil {
		t.Error("Managing not existing config dir failed: " + err0.Error())
	}
	var config DemoConfig1
	err1 := cm.Write("test", config)
	if err1 != nil {
		t.Error("Failed writing new empty config: " + err1.Error())
	}
	config.Key1 = "val1"
	config.Key2 = "val2"
	config.Key3 = "val3"
	err2 := cm.Write("test", config)
	if err2 != nil {
		t.Error("Failed writing existing config: " + err2.Error())
	}
	os.RemoveAll("cm_test")
}

func TestRead(t *testing.T) {
	os.RemoveAll("cm_test")
	cm, err0 := Manage("cm_test")
	if err0 != nil {
		t.Error("Managing not existing config dir failed: " + err0.Error())
	}
	var config DemoConfig1
	err1 := cm.Read("test", &config)
	if !(err1 != nil && os.IsNotExist(err1)) {
		t.Error("Reading not existing config did not throw expected error")
	}
	config = DemoConfig1{}
	cm.Write("test", config)
	var rconfig DemoConfig1
	err2 := cm.Read("test", &rconfig)
	if err2 != nil {
		t.Error("Failed reading empty existing config: " + err2.Error())
	}
	if config != rconfig {
		t.Error("Written and read empty config are not the same")
	}
	var config2 DemoConfig1
	config2.Key1 = "val1"
	config2.Key2 = "val2"
	config2.Key3 = "val3"
	cm.Write("test", config2)
	var rconfig2 DemoConfig1
	err3 := cm.Read("test", &rconfig2)
	if err3 != nil {
		t.Error("Failed reading existing config: " + err3.Error())
	}
	if config2 != rconfig2 {
		t.Error("Written and read config are not the same")
	}
	os.RemoveAll("cm_test")
}
