package csv

import (
	"encoding/csv"
	"io"
	"math/big"
	"time"
)

type CSV struct {
	w *csv.Writer
}

type CryptactLog struct {
	Timestamp    time.Time
	Action       string
	Source       string
	Base         string
	DerivType    string
	DerivDetails string
	Volume       *big.Float
	Price        *big.Float
	Counter      string
	Fee          *big.Float
	FeeCcy       string
	Comment      string
}

var header = []string{
	"Timestamp",
	"Action",
	"Source",
	"Base",
	"DerivType",
	"DerivDetails",
	"Volume",
	"Price",
	"Counter",
	"Fee",
	"FeeCcy",
	"Comment",
}

func New(writer io.Writer) *CSV {

	w := csv.NewWriter(writer)
	w.Write(header)

	return &CSV{w}
}

func (c *CSV) Flush() {
	c.w.Flush()
}

func (c *CSV) Add(clog *CryptactLog) {
	t := clog.Timestamp.In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format("2006/1/2 15:04:05")

	c.w.Write([]string{
		t,
		clog.Action,
		clog.Source,
		clog.Base,
		clog.DerivType,
		clog.DerivDetails,
		clog.Volume.String(),
		clog.Price.String(),
		clog.Counter,
		clog.Fee.String(),
		clog.FeeCcy,
		clog.Comment,
	})
}
