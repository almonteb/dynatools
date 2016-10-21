package streamer

import (
	"github.com/leevlad/dynago"
	"github.com/leevlad/dynago/streams"
)

type Config struct {
	arn      string
	executor streams.MakeRequester
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) WithExecutor(e *dynago.AwsExecutor) *Config {
	c.executor = e
	return &c
}

func (c Config) WithArn(arn string) *Config {
	c.arn = arn
	return &c
}
