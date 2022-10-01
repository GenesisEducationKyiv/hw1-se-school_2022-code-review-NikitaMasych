package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	ServerUrl         string
	CoinApiKey        string
	GenesisProvider   string
	EmailAddress      string
	EmailPassword     string
	StorageFile       string
	LoggerFile        string
	SMTPHost          string
	SMTPPort          int
	CacheHost         string
	CacheDb           int
	CacheDurationMins int
)

func init() {
	loadEnv()
	setupVariables()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err == nil {
		return
	}
	// In case we test from inner directories:
	max_directory_deep := 6
	for i := 1; i != max_directory_deep; i++ {
		escape_sequence := strings.Repeat("../", i)
		err = godotenv.Load("./" + escape_sequence + ".env")
		if err == nil {
			return
		}
	}
	log.Fatal(err)
}

func setupVariables() {
	var err error
	ServerUrl = os.Getenv("SERVER_URL")
	CoinApiKey = os.Getenv("COIN_API_KEY")
	GenesisProvider = os.Getenv("GENESIS_PROVIDER")
	EmailAddress = os.Getenv("EMAIL_ADDRESS")
	EmailPassword = os.Getenv("EMAIL_PASSWORD")
	StorageFile = os.Getenv("STORAGE_FILE")
	LoggerFile = os.Getenv("LOGGER_FILE")
	SMTPHost = os.Getenv("SMTP_HOST")
	SMTPPort, err = strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	CacheHost = os.Getenv("CACHE_HOST")
	CacheDb, err = strconv.Atoi(os.Getenv("CACHE_DB"))
	if err != nil {
		log.Fatal(err)
	}
	CacheDurationMins, err = strconv.Atoi(os.Getenv("CACHE_DURATION_MINS"))
	if err != nil {
		log.Fatal(err)
	}
}
