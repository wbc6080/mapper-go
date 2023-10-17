package grpcclient

import (
	"github.com/wbc6080/mapper-go/pkg/config"
)

var cfg *config.Config

func Init(c *config.Config) {
	cfg = &config.Config{}
	cfg = c
}
