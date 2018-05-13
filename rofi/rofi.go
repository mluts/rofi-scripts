package rofi

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Option interface {
	PrintedName() string
	Match(string) bool
	Activate()
}

func showOptions(opts []Option) {
	for _, opt := range opts {
		fmt.Println(opt.PrintedName())
	}
}

func activateOption(val string, opts []Option) {
	for _, opt := range opts {
		if opt.Match(val) {
			opt.Activate()
		}
	}
}

func ReadYamlFile(path string, out interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal([]byte(content), out)
}

func Exec(opts []Option, args []string) {
	if len(args) > 0 {
		activateOption(args[0], opts)
	} else {
		showOptions(opts)
	}
}
