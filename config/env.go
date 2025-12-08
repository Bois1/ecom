package config


type Config struct {
	PublicHost string
	Port 	string
	DBUser 	string
	DBPassword 	string
	DBAddress string
	DBName	string
}


func initConfig() Config {
	return Config{
		PublicHost: getEnv(),
	}
}

func getEnv( key, fallback string) string 