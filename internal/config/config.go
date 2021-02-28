package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultConfigType    = "yml"
	defaultDBFileName    = "db/default.db"
	defaultDBHost        = "localhost"
	defaultDBPort        = 3306
	defaultWebServerHost = "localhost"
	defaulWebServerPort  = 8080
	prefixEnvironmet     = "APP"
)

// Config is used for config.
type Config struct {
	DB     DBConfig
	WebSrv WebServerConfig
}

type (
	// DBConfig is used for data base config.
	DBConfig struct {
		FileName string
		Host     string
		Port     int32
	}

	// WebServerConfig for web server config.
	WebServerConfig struct {
		Host string
		Port int32
	}
)

// Init is used as constructor for config.
func Init(configFile string) (conf *Config, err error) {
	dir, fileName, err := parseFilePath(configFile)
	if err != nil {
		return nil, err
	}

	var runtimeViper = viper.New()

	// Set default values and init settings.
	populateDefaults(runtimeViper)

	// Load from file.
	runtimeViper.AddConfigPath(dir)
	runtimeViper.SetConfigName(fileName)
	runtimeViper.SetConfigType(defaultConfigType)

	// Load from evironment if exists.
	runtimeViper.SetEnvPrefix(prefixEnvironmet)
	runtimeViper.AutomaticEnv()

	// Fill viper map.
	if err = runtimeViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
	}

	// Fill config from viper.
	err = runtimeViper.Unmarshal(&conf)
	return
}

func populateDefaults(viperRuntime *viper.Viper) {
	viperRuntime.SetDefault("DB.FileNameDB", defaultDBFileName)
	viperRuntime.SetDefault("DB.Host", defaultDBHost)
	viperRuntime.SetDefault("DB.Port", defaultDBPort)
	viperRuntime.SetDefault("Server.Port", defaulWebServerPort)
	viperRuntime.SetDefault("Server.Host", defaultWebServerHost)
}

func parseFilePath(filePath string) (dir string, fileName string, err error) {
	path := strings.Split(filePath, "/")
	if len(path) < 2 {
		return "", "", ErrFileNotFound
	}
	dir = path[0]
	fileName = path[1]

	return
}
