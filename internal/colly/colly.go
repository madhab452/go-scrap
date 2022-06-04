package colly

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

// ReaderOpt reader options to
type ReaderOpt struct {
	PROVIDER string
	SRC      string
	TARGET   string
}

// Colly holds reading and read behavior
type Colly struct {
	opt ReaderOpt
}

// NewColly new instance of colly.
func NewColly(opt ReaderOpt) (*Colly, error) {
	return &Colly{
		opt: opt,
	}, nil
}

// Row a row of a transaction table.
type Row struct {
	SN              string `json:"-"`
	TradedCompany   string `json:"company"`
	NumOfTxn        string `json:"txn_num"`
	MaxPrice        string `json:"max_price"`
	MinPrice        string `json:"min_price"`
	ClosingPrice    string `json:"closing_price"`
	TradedShares    string `json:"traded_shares"`
	Amount          string `json:"amount"`
	PreviousClosing string `json:"prev_closing"`
	Diff            string `json:"diff"`
}

func mapPosition(row *Row, x int, data string) *Row {
	if x == 0 {
		row.SN = data
	}
	if x == 1 {
		row.TradedCompany = data
	}
	if x == 2 {
		row.NumOfTxn = data
	}
	if x == 3 {
		row.MaxPrice = data
	}
	if x == 4 {
		row.MinPrice = data
	}
	if x == 5 {
		row.ClosingPrice = data
	}
	if x == 6 {
		row.TradedShares = data
	}
	if x == 7 {
		row.Amount = data
	}
	if x == 8 {
		row.PreviousClosing = data
	}
	if x == 9 {
		var diff = data
		if strings.Contains(data, "Difference") {
			diff = data
		} else {
			for i := len(data); i > -1; i-- {
				pData := diff[0:i]
				_, err := strconv.ParseFloat(pData, 32)
				if err != nil {
					continue
				} else {
					diff = pData
				}
			}
		}
		row.Diff = diff
	}
	return row
}

func (co *Colly) Read() ([]*Row, error) {
	c := colly.NewCollector()

	var dsrc string

	if !strings.HasPrefix(co.opt.SRC, "http") { // file source
		fp, err := filepath.Abs(co.opt.SRC)
		if err != nil {
			return nil, fmt.Errorf("filepath.Abs(): %w", err)
		}
		t := &http.Transport{}
		t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

		c.WithTransport(t)
		dsrc = "file://" + fp
	} else {
		dsrc = co.opt.SRC
	}

	var result []*Row

	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		table := e.DOM.Find("table").First()
		table.Find("tr").Each(func(i int, tr *goquery.Selection) {
			if tds := tr.Find("td"); tds.Length() == 10 {
				row := &Row{}
				tr.Find("td").Each(func(j int, s *goquery.Selection) {
					row = mapPosition(row, j, s.Text()) // inefficient. multiple copies
				})
				result = append(result, row)
			}
		})
	})

	err := c.Visit(dsrc)
	if err != nil {
		return nil, fmt.Errorf("c.Visit(%v): %w", dsrc, err)
	}

	return result, nil
}
