package config

type Config struct {
	EthereumNodeURL string
	SolanaNodeURL   string
}

func LoadConfig() *Config {
	// TODO:Load configuration values from environment variables or a configuration file
	return &Config{
		EthereumNodeURL: "TODO",
		SolanaNodeURL:   "TODO",
	}
}
