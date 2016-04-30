import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/config"
	"github.com/go-ini/ini"
	"fmt"
	"strings"
)

type IniConfigLoader struct {
	file string
	loader *ini.File
}

func NewConfigLoader(type string, file string) (*IniConfigLoader, error) {
  if(type == "ini") {
    config := new(IniConfigLoader)
  	_, err := config.load(file)
  } else {
    return _, fmt.Errorf("NewConfigLoader: unsupported format: %s", type)
  }
	
	return config, err
}

func (this *IniConfigLoader) GetSection(name string) (*ini.Section, error) {
	return this.loader.GetSection(name)
}

func (this *IniConfigLoader) load(filename string) (*ini.File, error) {
	this.file = filename
	loader, err := ini.LooseLoad(filename)
	this.loader = loader
	return this.loader, err
}

func (this *IniConfigLoader) String(section string, key string, default_value string) string {
	if this.loader.Section(section).HasKey(key) {
		return this.loader.Section(section).Key(key).String()
	} else if this.loader.Section("").HasKey(key) {
		return this.loader.Section("").Key(key).String()
	} else {
		return default_value
	}	
}

func (this *IniConfigLoader) Int(section string, key string, default_value int) int {
	if this.loader.Section(section).HasKey(key) {
		return this.loader.Section(section).Key(key).MustInt(default_value)
	} else if this.loader.Section("").HasKey(key) {
		return this.loader.Section("").Key(key).MustInt(default_value)
	} else {
		return (default_value)
	}	
}

func (this *IniConfigLoader) Float(section string, key string, default_value float64) float64 {
	if this.loader.Section(section).HasKey(key) {
		return this.loader.Section(section).Key(key).MustFloat64(default_value)
	} else if this.loader.Section("").HasKey(key) {
		return this.loader.Section("").Key(key).MustFloat64(default_value)
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
