package repository

import (
	"time"
)

type ClosingID string

type Closings = []Closing

type Closing struct {
	ID     ClosingID
	Start  time.Time
	End    time.Time
	Sales  Sales
	Costs  float64
	Profit float64
}

func (c *Closing) CalculateProfit() {
	for _, sale := range c.Sales {
		c.Profit += sale.Value
	}

	c.Profit -= c.Costs
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
