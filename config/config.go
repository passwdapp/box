package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds the global configuration of the app
type Config struct {
	ListenAddress   string
	SecretKey       string
	MaxUsers        int64
	UploadDirectory string
	DatabaseFile    string
	JWTSecret       string
}

// SetDefaults sets the config to default
func (c *Config) SetDefaults() {
	c.DatabaseFile = "data/db.sqlite"
	c.ListenAddress = "127.0.0.1:3000"
	c.MaxUsers = 6
	c.SecretKey = "qwerty" // Should be present in the env
	c.UploadDirectory = "data/upload"
}

// InitConfig initializes the configuration object
func (c *Config) InitConfig(dotenv bool) {
	c.SetDefaults()

	if dotenv {
		c.GetFromDotEnv()
	} else {
		c.GetFromEnv()
	}
}

// GetFromDotEnv gets the config from the .env file
func (c *Config) GetFromDotEnv() {
	godotenv.Load()
	c.GetFromEnv()
}

// GetFromEnv gets the config from environment variables
func (c *Config) GetFromEnv() {
	listenAddress, listenAddressPresent := os.LookupEnv("LISTEN_ADDRESS")
	secretKey, secretKeyPresent := os.LookupEnv("SECRET_KEY")
	maxUsers, maxUsersPresent := os.LookupEnv("MAX_USERS")
	jwtSecret, jwtSecretPresent := os.LookupEnv("JWT_SECRET")

	if !secretKeyPresent {
		log.Fatalln("No SECRET_KEY present in the environment")
	}

	if !jwtSecretPresent {
		log.Fatalln("No JWT_SECRET present in the environment")
	}

	c.SecretKey = secretKey
	c.JWTSecret = jwtSecret

	if listenAddressPresent {
		c.ListenAddress = listenAddress
	}

	if maxUsersPresent {
		parsedMaxUsers, err := strconv.ParseInt(maxUsers, 10, 64)

		if err != nil {
			log.Fatalln(err)
		}

		c.MaxUsers = parsedMaxUsers
	}
}

// Version is the box version
const Version = "1.0"

var config *Config

// GetConfig returns the configuration data.
func GetConfig() *Config {
	return config
}

// InitConfig initializes the configuration data
func InitConfig(dotenv bool) {
	conf := &Config{}
	conf.InitConfig(dotenv)

	config = conf
}

// SetConfig sets configuration data.
func SetConfig(conf *Config) {
	config = conf
}
