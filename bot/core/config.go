package core

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	WebhookURL    string `env:"WEBHOOK_URL"`
	WebhookPath   string `env:"WEBHOOK_PATH"`
	WebhookListen string `env:"WEBHOOK_LISTEN"`
	BotAPIKey     string `env:"BOT_API_KEY,required"`
	DatabaseURL   string `env:"DATABASE_URL,required"`
	RedisAddress  string `env:"REDIS_ADDRESS,required"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	BotVer        string `env:"BOT_VERSION,required"`
	WebhookPort   int    `env:"WEBHOOK_PORT"`
	OwnerID       int    `env:"OWNER_ID,required"`
	SudoUsers     []int  `env:"SUDO_USERS,required" envSeparator:":"`
	LogEvent      int64  `env:"LOG_EVENT,required"`
	LogBan        int64  `env:"LOG_BAN,required"`
	MainGroup     int64  `env:"MAIN_GRP,required"`
	IsDebug       bool   `env:"IS_DEBUG"`
	CleanPolling  bool   `env:"CLEAN_POLLING,required"`
}

func NewConfig() *Config {
	conf := new(Config)

	err := godotenv.Load("data/.env")
	if err != nil {
		log.Println("Using declared Env vars!")

		conf.WebhookURL = os.Getenv("WEBHOOK_URL")
		conf.WebhookPath = os.Getenv("WEBHOOK_PATH")

		port, _ := strconv.Atoi(os.Getenv("PORT"))
		if port != 0 {
			conf.WebhookPort = port
		} else {
			conf.WebhookPort, _ = strconv.Atoi(os.Getenv("WEBHOOK_PORT"))
		}

		conf.WebhookListen = os.Getenv("WEBHOOK_LISTEN")
		conf.BotAPIKey = os.Getenv("BOT_API_KEY")
		_, cleanPolling := os.LookupEnv("CLEAN_POLLING")
		conf.CleanPolling = cleanPolling
		_, isDebug := os.LookupEnv("IS_DEBUG")
		conf.IsDebug = isDebug

		logEvent, _ := strconv.Atoi(os.Getenv("LOG_EVENT"))
		conf.LogEvent = int64(logEvent)
		logBan, _ := strconv.Atoi(os.Getenv("LOG_BAN"))
		conf.LogBan = int64(logBan)
		conf.OwnerID, _ = strconv.Atoi(os.Getenv("OWNER_ID"))
		conf.SudoUsers = strToIntSlice(strings.Split(os.Getenv("SUDO_USERS"), ":"))

		conf.DatabaseURL = os.Getenv("DATABASE_URL")
		conf.RedisAddress = os.Getenv("REDIS_ADDRESS")
		conf.RedisPassword = os.Getenv("REDIS_PASSWORD")

		return conf
	}

	err = env.Parse(conf)
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	log.Println("Configurations have been parsed succesfully!")
	return conf
}

func strToIntSlice(s []string) []int {
	var newIntSlice []int
	for _, val := range s {
		newInt, _ := strconv.Atoi(val)
		newIntSlice = append(newIntSlice, newInt)
	}
	return newIntSlice
}
