package colly

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

type Opt struct {
	PROVIDER  string
	DSRC      string
	URL       string
	FILE_PATH string
}

type Colly struct {
}

func NewFileColly() *Colly {
	return &Colly{}

}

func NewInternetColly() *Colly {
	return &Colly{}
}

func NewColly(ctx context.Context, log *logrus.Entry, opt *Opt) (*Colly, error) {

	var target = ""
	fmt.Println(opt)
	c := colly.NewCollector()

	if opt.DSRC == "file" {
		filepath, err := filepath.Abs(opt.FILE_PATH)
		if err != nil {
			panic(err)
		}
		t := &http.Transport{}
		t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

		c.WithTransport(t)
		target = "file://" + filepath
	} else {
		target = opt.URL
	}

	err := c.Visit(target)
	if err != nil {
		panic(err)
	}
	return &Colly{}, nil
}
