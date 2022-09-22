package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerURL              string
	BinanceApiFormatUrl    string
	CoinbaseApiFormatUrl   string
	CoinApiFormatURL       string
	CoinApiKey             string
	BaseCurrency           string
	QuotedCurrency         string
	CryptoCurrencyProvider string
	EmailAddress           string
	EmailPassword          string
	StorageFile            string
	LoggerFile             string
	SMTPHost               string
	SMTPPort               int
	CacheDurationMins      int
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {

	once.Do(func() {
		LoadEnv()
		cfg = Config{
			ServerURL:              os.Getenv(ServerUrl),
			BinanceApiFormatUrl:    os.Getenv(BinanceApiFormatUrl),
			CoinbaseApiFormatUrl:   os.Getenv(CoinbaseApiFormatUrl),
			CoinApiFormatURL:       os.Getenv(CoinApiFormatURL),
			CoinApiKey:             os.Getenv(CoinApiKey),
			BaseCurrency:           os.Getenv(BaseCurrency),
			QuotedCurrency:         os.Getenv(QuotedCurrency),
			CryptoCurrencyProvider: os.Getenv(CryptoCurrencyProvider),
			EmailAddress:           os.Getenv(EmailAddress),
			EmailPassword:          os.Getenv(EmailPassword),
			StorageFile:            os.Getenv(StorageFile),
			LoggerFile:             os.Getenv(LoggerFile),
			SMTPHost:               os.Getenv(SMTPHost),
		}
		cfg.SMTPPort, _ = strconv.Atoi(os.Getenv(SMTPPort))
		cfg.CacheDurationMins, _ = strconv.Atoi(os.Getenv(CacheDurationMins))
	})
	return &cfg
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		// in case we test from inner directories;
		// sequence to go to the upper ones
		err = godotenv.Load("./../../.env")
		if err != nil {
			err = godotenv.Load("./../.env")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
