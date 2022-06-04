package main

import (
	"context"
	"os"

	"github.com/madhab452/go-scrap/internal"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func main() {
	log = logrus.NewEntry(logrus.New())

	conf := internal.Conf{
		PROVIDER: os.Getenv("PROVIDER"),
		SRC:      os.Getenv("SRC"),
		TARGET:   os.Getenv("TARGET"),
	}

	s, err := internal.New(context.Background(), log, &conf)
	if err != nil {
		log.WithError(err).Printf("internal.New()")
		return
	}

	if err := s.ReadAndWrite(); err != nil {
		log.WithError(err).Printf("s.ReadAndWrite()")
		return
	}

	log.Printf("successfully scrapped the data")
}
