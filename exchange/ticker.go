package exchange

import (
	"fmt"
	"time"
)

/*
{
    "instrument": "$PAC-BTC",
    "last": 0.0000005,
    "percentChange": 8.6957,
    "low": 0.00000046,
    "high": 0.0000005,
    "baseVolume": 144.855144855145,
    "quoteVolume": 0.0000724275724275725,
    "volumeInBtc": 0.0000724275724275725,
    "volumeInUsd": 0.538195529470530008725,
    "ask": 0.0000005,
    "bid": 0.00000046,
    "timestamp": "2018-05-31T12:48:56Z"
  }
*/

type Ticker struct {
	Instrument    string    `json:"instrument"`
	Last          float64   `json:"last"`
	PercentChange float64   `json:"percentChange"`
	Low           float64   `json:"low"`
	High          float64   `json:"high"`
	BaseVolume    float64   `json:"baseVolume"`
	QuoteVolume   float64   `json:"quoteVolume"`
	VolumeInBtc   float64   `json:"volumeInBtc"`
	VolumeInUsd   float64   `json:"volumeInUsd"`
	Ask           float64   `json:"ask"`
	Bid           float64   `json:"bid"`
	Timestamp     time.Time `json:"timestamp"`
}

func (t *Ticker) String() (s string) {
	s = fmt.Sprintf(
		"(Ticker) %s = Last: %.8f, High: %.8f, Low: %.8f, Ask: %.8f, Bid: %.8f, Volume (BTC): %.8f",
		t.Instrument, t.Last, t.High, t.Low, t.Ask, t.Bid, t.VolumeInBtc,
	)
	return
}

type Tickers []*Ticker

func (e *Exchange) Ticker(instrument string) (i *Ticker, err error) {
	params := EmptyParams()

	var is Tickers
	err = e.getJSON("/v2/public/tickers?instrument="+instrument, params, &is, false)
	if err == nil && len(is) > 0 {
		i = is[0]
	}
	return
}

func (e *Exchange) Tickers() (is Tickers, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/public/tickers", params, &is, false)
	return
}