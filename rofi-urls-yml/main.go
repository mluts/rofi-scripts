package main

import (
	"github.com/mluts/rofi-scripts/rofi"
	"os"
	"os/exec"
	"strings"
)

const OPTS_FILE = "$HOME/.rofi-urls.yml"

type UrlOption struct {
	Name string `yaml:"name,omitempty"`
	Val  string `yaml:"url"`
}

func (o *UrlOption) PrintedName() string {
	return o.Name
}

func (o *UrlOption) Match(name string) bool {
	return strings.TrimSpace(name) == strings.TrimSpace(o.Name)
}

func (o *UrlOption) Activate() {
	exec.Command("/usr/bin/google-chrome-stable", o.Val).Run()
}

func (o *UrlOption) Format() {
	if len(o.Name) == 0 && len(o.Val) > 0 {
		o.Name = o.Val
	}
}

func formatOptions(opts []*UrlOption) {
	for _, o := range opts {
		o.Format()
	}
}

func getRofiOptions() (rofiOpts []rofi.Option) {
	opts := []*UrlOption{}
	err := rofi.ReadYamlFile(os.ExpandEnv(OPTS_FILE), &opts)
	if err != nil {
		panic(err)
	}

	for _, opt := range opts {
		rofiOpts = append(rofiOpts, rofi.Option(opt))
	}

	return rofiOpts
}

func main() {
	opts := getRofiOptions()
	rofi.Exec(opts, os.Args[1:])
	os.Exit(0)
}
