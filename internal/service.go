package internal

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/madhab452/go-scrap/internal/colly"
	"github.com/madhab452/go-scrap/internal/writer/cwriter"
	"github.com/madhab452/go-scrap/internal/writer/tcpwriter"
)

// Conf represents configuration options for scrapper
type Conf struct {
	PROVIDER   string
	DSRC       string
	URL        string
	FILE_PATH  string
	TARGET_URL string
}
type Reader interface {
	Read() ([]colly.Row, error)
}
type Writer interface {
	Write([]colly.Row) error
}

type Service struct {
	log  *logrus.Entry
	conf *Conf

	reader Reader
	writer Writer
}

func (s *Service) ReadAndWrite() error {
	rows, err := s.reader.Read()
	if err != nil {
		return fmt.Errorf("s.reader.Read(): %w", err)
	}

	if err := s.writer.Write(rows); err != nil {
		return fmt.Errorf("s.writer.Write: %w", err)
	}
	return nil
}

// New creates an instance of Service
func New(ctx context.Context, log *logrus.Entry, conf *Conf) (*Service, error) {
	opt := colly.ReaderOpt{
		PROVIDER:  conf.PROVIDER,
		DSRC:      conf.DSRC,
		URL:       conf.URL,
		FILE_PATH: conf.FILE_PATH,
	}

	if conf.DSRC == "" {
		return nil, fmt.Errorf("required env var DSRC is missing")
	} else {
		if conf.DSRC == "INTERNET" && conf.URL == "" {
			return nil, fmt.Errorf("remote url for data source is missing")
		} else if conf.DSRC == "FILE" && conf.FILE_PATH == "" {
			return nil, fmt.Errorf("file path to read the data is missing")
		}
	}

	var w Writer
	if conf.TARGET_URL != "" {
		var err error
		w, err = tcpwriter.New(ctx, log, conf.TARGET_URL)
		if err != nil {
			return nil, fmt.Errorf("tcpwriter.New(): %w", err)
		}
	} else {
		w, _ = cwriter.New()
	}

	reader, err := colly.NewColly(opt)
	if err != nil {
		return nil, fmt.Errorf("colly.NewColly(): %w", err)
	}

	svc := &Service{
		log:    log,
		reader: reader,
		conf:   conf,
		writer: w,
	}
	return svc, nil
}
