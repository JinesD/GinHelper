package ginhelper

import (
	"log"

	"github.com/lestrrat-go/file-rotatelogs"
)

func SetupLogger(file string) error {
	rl, err := rotatelogs.New(file)
	if err != nil {
		return err
	}

	log.SetOutput(rl)
	return nil
}
