package main

import (
	"airline-voucher-seat-assignment/app/http"
	"airline-voucher-seat-assignment/app/http/handler"
	"airline-voucher-seat-assignment/internal/config/database"
	"airline-voucher-seat-assignment/internal/entity"
	voucherRepo "airline-voucher-seat-assignment/internal/repository/voucher"
	voucherService "airline-voucher-seat-assignment/internal/service/voucher"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	logSetup()

	db, err := database.New()

	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	db.AutoMigrate(
		&entity.Voucher{},
	)

	vr := voucherRepo.New(db)

	vs := voucherService.New(vr)

	h := handler.New(vs)

	s := http.NewServer(h)
	s.Run()
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
