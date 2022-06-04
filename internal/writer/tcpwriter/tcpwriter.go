// tcpwriter send http.Post request
package tcpwriter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/madhab452/go-scrap/internal/colly"
)

type TCPWriter struct {
	ctx       context.Context
	log       *logrus.Entry
	targetURL string
}

func (tw *TCPWriter) Write(rows []colly.Row) error {
	for i := 1; i < len(rows); i++ {
		r := rows[i]
		json_data, err := json.Marshal(r)
		if err != nil {
			return fmt.Errorf("json.Marshal(): %w", err)
		}

		resp, err := http.Post(tw.targetURL, "application/json", strings.NewReader(string(json_data)))
		if err != nil {
			return fmt.Errorf("http.Post(): %w", err)
		}
		defer resp.Body.Close()

		if _, err := io.ReadAll(resp.Body); err != nil {
			return fmt.Errorf("ioutil.ReadAll(): %w", err)
		}
	}
	return nil
}

func New(ctx context.Context, log *logrus.Entry, targetURL string) (*TCPWriter, error) {
	return &TCPWriter{
		ctx:       ctx,
		log:       log,
		targetURL: targetURL,
	}, nil
}
