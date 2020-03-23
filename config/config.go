package config

type Config struct {
	MongoConfig *MongoConfig
	Port        string
}

type MongoConfig struct {
	ConnectionUrl string
	//	Host     string
	//	Port     int
	//	Username string
	//	Password string
}

func GetConfig() *Config {
	return &Config{
		MongoConfig: &MongoConfig{
			ConnectionUrl: "mongodb://localhost:27017",
		},
		Port: ":6000",
	}
}
