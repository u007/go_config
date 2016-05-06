package go_config

import (
	"github.com/go-ini/ini"
	"fmt"
	"strings"
	"os"
	"path/filepath"
)

type IniConfigLoader struct {
	file string
	loader *ini.File
}

func NewConfigLoader(_type string, file string) (*IniConfigLoader, error) {
  if(_type == "ini") {
		path, err := filepath.Abs(file)
		if (err != nil) {
			err_f := fmt.Errorf("unable to get abs path: %s", file)
			return nil, err_f
		}
		fmt.Printf("path: %s\n", path)
		_, err = os.Stat(path);
		if  err != nil {
			fmt.Printf("ERROR!!!!: %s\n", err.Error())
		  return nil, err
		}
    config := new(IniConfigLoader)
  	_, err = config.load(file)
		return config, err
  } else {
    return nil, fmt.Errorf("NewConfigLoader: unsupported format: %s", _type)
  }
}

func (this *IniConfigLoader) Section(name string) *ini.Section {
	// fmt.Printf("loader config %v", this)
	section, err := this.loader.GetSection(name)
	if err != nil {
		section, err = this.loader.GetSection("")
		if err != nil {
			err_f := fmt.Errorf("Section %s does not exists", name)
			fmt.Printf(err_f.Error())
			return nil
		}
		fmt.Printf("section fallback to root for %s", name)
		return section
	}
	return section
}

func (this *IniConfigLoader) load(filename string) (*ini.File, error) {
	this.file = filename
	loader, err := ini.LooseLoad(filename)
	this.loader = loader
	return this.loader, err
}

func (this *IniConfigLoader) String(section string, key string, default_value string) string {
	// fmt.Printf("Getting string: %s > %s\n", section, key)
	if this.Section(section).HasKey(key) {
		return this.Section(section).Key(key).MustString(default_value)
	} else if (this.loader.Section("").HasKey(key)) {
		return this.loader.Section("").Key(key).MustString(default_value)
	} else {
		return default_value
	}
}

func (this *IniConfigLoader) Int(section string, key string, default_value int) int {
	if this.Section(section).HasKey(key) {
		return this.Section(section).Key(key).MustInt(default_value)
	} else if (this.loader.Section("").HasKey(key)) {
		return this.loader.Section("").Key(key).MustInt()
	} else {
		return default_value
	}
}

func (this *IniConfigLoader) Float(section string, key string, default_value float64) float64 {
	if this.Section(section).HasKey(key) {
		return this.Section(section).Key(key).MustFloat64(default_value)
	} else if (this.loader.Section("").HasKey(key)) {
		return this.loader.Section("").Key(key).MustFloat64()
	} else {
		return default_value
	}
}

func (this *IniConfigLoader) Boolean(section string, key string, default_value bool) bool {
	default_string := "false"
	if (default_value){
		default_string = "true"
	}
	value := this.String(section, key, default_string)

	if(strings.ToLower(value) == "true"){
		return true
	} else {
		return false
	}
}
