// cwriter console writer, writes to stdout
package cwriter

import (
	"fmt"

	"github.com/madhab452/go-scrap/internal/colly"
)

// CWriter console writer
type CWriter struct {
}

// Write writes to console
func (cw *CWriter) Write(rows []*colly.Row) error {
	for _, row := range rows {
		fmt.Println(row)
	}
	return nil
}

// New returns new instance of console writer
func New() (*CWriter, error) {
	return &CWriter{}, nil
}
