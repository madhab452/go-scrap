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
	PROVIDER string
	SRC      string
	TARGET   string
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
		PROVIDER: conf.PROVIDER,
		SRC:      conf.SRC,
		TARGET:   conf.TARGET,
	}

	if conf.SRC == "" {
		return nil, fmt.Errorf("required env var SRC missing")
	}

	var w Writer
	if conf.TARGET != "" {
		var err error
		w, err = tcpwriter.New(ctx, log, conf.TARGET)
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
