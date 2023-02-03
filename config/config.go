package config

type Configuration struct {
	Active string `json:"active" yaml:"active"`
	Server Server `json:"server" yaml:"server"`
}

type Server struct {
	Port int `json:"port" yaml:"port"`
}
