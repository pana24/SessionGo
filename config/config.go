package config

type Configuration struct {
	Active string `json:"active" yaml:"active"`
	Server Server `json:"server" yaml:"server"`
	Opa    Opa    `json:"opa" yaml:"opa"`
}

type Opa struct {
	Url string `json:"url" yaml:"url"`
}

type Server struct {
	Port int `json:"port" yaml:"port"`
}
