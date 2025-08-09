package config

type Config struct {
	Database mysql `json:"database"`
	Auth     Auth  `json:"auth"`
}

type mysql struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

type Auth struct {
	SecretKey string `json:"secretKey"`
}
