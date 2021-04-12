package cliargs

import (
	"strings"
)

func (o CmdOptions) Has(key string) bool {
	_, ok := o.options[key]
	return ok
}

func (o CmdOptions) First(key string) string {
	return o.N(key, 0)
}

func (o CmdOptions) N(key string, index int) string {
	if value, ok := o.options[key]; ok {
		if index >= len(value) {
			return ""
		}
		return value[index]
	}
	return ""
}

func (o CmdOptions) Get(key string) []string {
	if value, ok := o.options[key]; ok {
		return value
	}
	return []string{}
}

func (o CmdOptions) Add(arg string) {
	params := strings.SplitN(arg, "=", 2)
	key := ""
	value := ""
	if strings.HasPrefix(params[0], "--") {
		key = params[0][2:]
		if len(params) == 1 {
			value = "true"
		} else {
			value = params[1]
		}
	}

	if !o.Has(key) {
		o.options[key] = []string{}
	}
	o.options[key] = append(o.options[key], value)
}
