package configuration

var EnvConfigs *Config
type Config struct {
	DBUsername string `mapstructure:"DB_USERNAME`
	DBPassword string `mapstructure: "DB_PASSWORD`
	DBName string	`mapstructure: "DB_NAME"`
	DBPort string `mapstructure: "DB_PORT"`
	DBHost string `mapstructure: "DB_HOST"`
	SecretToken string `mapstructure: "JWT_SECRET"`
}