package main

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	logSetup()

	if err := godotenv.Load(); err != nil {
		log.Error(err.Error())
		panic(err)
	}
}

func logSetup() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
