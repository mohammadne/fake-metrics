package http

type Config struct {
	Address string `koanf:"address"`
	Port    int    `koanf:"port"`
}
