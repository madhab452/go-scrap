package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/madhab452/go-scrap/internal/colly"
	"github.com/sirupsen/logrus"
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
type Service struct {
	log    *logrus.Entry
	reader Reader
	conf   *Conf
}

func (s *Service) ReadAndWrite() error {
	rows, err := s.reader.Read()
	if err != nil {
		return fmt.Errorf("s.reader.Read(): %w", err)
	}
	if err := write(s.conf.TARGET_URL, rows); err != nil {
		return fmt.Errorf("write(): %w", err)
	}
	return nil
}

func write(targeturl string, rows []colly.Row) error {
	for i := 1; i < len(rows); i++ {
		r := rows[i]
		json_data, err := json.Marshal(r)
		if err != nil {
			return fmt.Errorf("json.Marshal(): %w", err)
		}

		//TODO: send rpc request
		resp, err := http.Post(targeturl, "application/json", strings.NewReader(string(json_data)))
		if err != nil {
			return fmt.Errorf("http.Post(): %w", err)
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ioutil.ReadAll(): %w", err)
		}
		fmt.Println(string(bodyBytes))

		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)
	}
	return nil
}

// New creates an instance of Service
func New(context context.Context, log *logrus.Entry, conf *Conf) (*Service, error) {
	opt := colly.ReaderOpt{
		PROVIDER:  conf.PROVIDER,
		DSRC:      conf.DSRC,
		URL:       conf.URL,
		FILE_PATH: conf.FILE_PATH,
	}

	reader, err := colly.NewColly(&opt)
	if err != nil {
		return nil, fmt.Errorf("colly.NewColly(): %w", err)
	}

	return &Service{
		log:    log,
		reader: reader,
		conf:   conf,
	}, nil
}
