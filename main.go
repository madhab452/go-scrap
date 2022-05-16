package main

import (
	"context"
	"os"

	"github.com/madhab452/go-scrap/internal"
	"github.com/sirupsen/logrus"
)

var log logrus.Entry

func main() {
	conf := internal.Conf{
		PROVIDER:  os.Getenv("PROVIDER"),
		DSRC:      os.Getenv("DSRC"),
		URL:       os.Getenv("SOURCE"),
		FILE_PATH: os.Getenv("FILE_PATH"),
	}
	s, err := internal.New(context.Background(), &log, &conf)
	if err != nil {
		log.Panic(err)
	}
	s.Read()
}
