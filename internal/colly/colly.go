package colly

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type ReaderOpt struct {
	PROVIDER  string
	DSRC      string
	URL       string
	FILE_PATH string
}

type Colly struct {
	opt *ReaderOpt
}

func NewColly(opt *ReaderOpt) (*Colly, error) {
	return &Colly{
		opt: opt,
	}, nil
}

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

func mapPosition(row Row, x int, data string) Row {
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

func (co *Colly) Read() ([]Row, error) {
	var target = ""
	c := colly.NewCollector()

	if co.opt.DSRC == "file" {
		filepath, err := filepath.Abs(co.opt.FILE_PATH)
		if err != nil {
			panic(err)
		}
		t := &http.Transport{}
		t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

		c.WithTransport(t)
		target = "file://" + filepath
	} else {
		target = co.opt.URL
	}

	var result []Row

	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		table := e.DOM.Find("table").First()
		table.Find("tr").Each(func(i int, tr *goquery.Selection) {

			if tds := tr.Find("td"); tds.Length() == 10 {
				row := Row{}
				tr.Find("td").Each(func(j int, s *goquery.Selection) {
					row = mapPosition(row, j, s.Text()) // inefficient. multiple copies
				})
				result = append(result, row)
			}
		})
	})

	err := c.Visit(target)
	if err != nil {
		return nil, err
	}
	return result, nil
}
