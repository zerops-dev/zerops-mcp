package httpClient

import "time"

type Config struct {
	Timeout time.Duration `json:"timeout"`
}

func NewConfig() Config {
	return Config{
		Timeout: time.Second * 30,
	}
}
