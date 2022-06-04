package main

import (
	"context"
	"fmt"
	"os"

	"github.com/madhab452/go-scrap/internal"
	"github.com/sirupsen/logrus"
)

var log logrus.Entry

func main() {
	conf := internal.Conf{
		PROVIDER:   os.Getenv("PROVIDER"),
		DSRC:       os.Getenv("DSRC"),
		URL:        os.Getenv("URL"),
		FILE_PATH:  os.Getenv("FILE_PATH"),
		TARGET_URL: os.Getenv("TARGET_URL"),
	}

	s, err := internal.New(context.Background(), &log, &conf)
	if err != nil {
		log.Panic(err)
	}

	if err := s.ReadAndWrite(); err != nil {
		fmt.Println(err)
	}
}
