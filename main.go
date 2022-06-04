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
		PROVIDER: os.Getenv("PROVIDER"),
		SRC:      os.Getenv("SRC"),
		TARGET:   os.Getenv("TARGET"),
	}

	s, err := internal.New(context.Background(), &log, &conf)
	if err != nil {
		log.Panic(err)
	}

	if err := s.ReadAndWrite(); err != nil {
		fmt.Println(err)
	}
}
