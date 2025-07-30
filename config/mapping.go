package config

type Config struct {
	Database mysql `json:"database"`
}

type mysql struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}
