package logic

import (
	"time"
)

type ClosingID string

type Closings = []Closing

type Closing struct {
	ID        ClosingID
	Start     time.Time
	End       time.Time
	Sales     Sales
	Costs     float64
	NetProfit float64
}

func (c *Closing) GrossProfit() (gross float64) {
	for _, sale := range c.Sales {
		gross += sale.Value
	}

	return
}

func (c *Closing) CalculateNetProfit() {
	c.NetProfit = c.GrossProfit() - c.Costs
}

type Sales = []Sale

type Sale struct {
	Date  time.Time
	Value float64
}

const (
	Undefined ClosingStatus = iota
	Open
	Closed
)

type ClosingStatus uint8
