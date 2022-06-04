// tcpwriter send http.Post request
package tcpwriter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/madhab452/go-scrap/internal/colly"
)

type TcpWriter struct {
	ctx       context.Context
	log       *logrus.Entry
	targetURL string
}

func (tw *TcpWriter) Write(rows []colly.Row) error {
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

func New(ctx context.Context, log *logrus.Entry, targetURL string) (*TcpWriter, error) {
	return &TcpWriter{
		ctx:       ctx,
		log:       log,
		targetURL: targetURL,
	}, nil
}
