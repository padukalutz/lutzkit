package config

const (
	ConfigFileName = ".lutzkit.json"
	ConfigVersion  = 1
)

type Config struct {
	Version int `json:"version"`
}

func Default() Config {
	return Config{
		Version: ConfigVersion,
	}
}
