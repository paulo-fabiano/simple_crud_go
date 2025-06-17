package config

import (
	"io"
	"log"
)

type Logger struct {
	debug 	*log.Logger
	info	*log.Logger
	warning	*log.Logger
	err		*log.Logger
	writer	io.Writer
}

func NewLogger() {

}