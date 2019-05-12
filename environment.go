package ginhelper

import (
	"github.com/go-ini/ini"
)

func EnvLoader(file string) (*ini.File, error) {

	loader, err := ini.Load(file)
	if err != nil {
		return nil, err
	}

	return loader, nil
}
