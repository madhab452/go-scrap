// cwriter console writer, writes to stdout
package cwriter

import (
	"fmt"

	"github.com/madhab452/go-scrap/internal/colly"
)

type CWriter struct {
}

func (cw *CWriter) Write(rows []*colly.Row) error {
	for _, row := range rows {
		fmt.Println(row)
	}
	return nil
}

func New() (*CWriter, error) {
	return &CWriter{}, nil
}
