package utils

import "github.com/joho/godotenv"

type Config struct {
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// func LoadConfig(path string) (config Config, err error) {
func LoadConfig() {
	// log.Println(path)
	// viper.AddConfigPath(path)
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")

	// viper.AutomaticEnv()

	// err = viper.ReadInConfig()
	// if err != nil {
	// 	return
	// }
	// log.Println(" === > ", err)

	// err = viper.Unmarshal(&config)
	// return
	godotenv.Load(".env")

}
