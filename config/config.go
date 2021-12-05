package config

import "github.com/spf13/viper"

type Config struct {
	HostPort          string `mapstructure:"HOST_PORT"`
	UrlRedirect       string `mapstructure:"URL_REDIRECT"`
	FirstUrlRedirect  string `mapstructure:"FIRST_URL_REDIRECT"`
	SecondUrlRedirect string `mapstructure:"SECOND_URL_REDIRECT"`
	FolderContents    string `mapstructure:"FOLDER_CONTENT"`
	Protocol          string `mapstructure:"PROTOCOL"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
