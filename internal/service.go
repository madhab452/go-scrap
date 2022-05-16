package internal

import (
	"context"

	"github.com/madhab452/go-scrap/internal/colly"
	"github.com/sirupsen/logrus"
)

// Conf represents configuration options for scrapper
type Conf struct {
	PROVIDER  string
	DSRC      string
	URL       string
	FILE_PATH string
}

type Service struct {
}

func (s *Service) Read() {

}

// New creates an instance of Service
func New(context context.Context, log *logrus.Entry, conf *Conf) (*Service, error) {
	opt := colly.Opt{
		PROVIDER:  conf.PROVIDER,
		DSRC:      conf.DSRC,
		URL:       conf.URL,
		FILE_PATH: conf.FILE_PATH,
	}
	colly.NewColly(context, log, &opt)

	return &Service{}, nil
}
