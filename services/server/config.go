package server

type Config struct {
	Neco bool
}

func NewConfig() Config {
	return Config{
		Neco: true,
	}
}
