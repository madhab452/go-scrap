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

// TCPWriter holds target url and behavior to send POST request to that remote url.
type TCPWriter struct {
	ctx       context.Context
	log       *logrus.Entry
	targetURL string
}

// Write posts to some remote location: targetURL
func (tw *TCPWriter) Write(rows []*colly.Row) error {
	for i := 1; i < len(rows); i++ {
		r := rows[i]
		jsonData, err := json.Marshal(r)
		if err != nil {
			return fmt.Errorf("json.Marshal(): %w", err)
		}

		resp, err := http.Post(tw.targetURL, "application/json", strings.NewReader(string(jsonData)))
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

// New returns an instance of TCPWriter
func New(ctx context.Context, log *logrus.Entry, targetURL string) (*TCPWriter, error) {
	return &TCPWriter{
		ctx:       ctx,
		log:       log,
		targetURL: targetURL,
	}, nil
}
