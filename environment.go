package ginhelper

import (
	"github.com/go-ini/ini"
)

func EnvLoader() (*ini.File, error) {
	envFile := RootPath() + "env.ini"

	loader, err := ini.Load(envFile)
	if err != nil {
		return nil, err
	}

	return loader, nil
}
