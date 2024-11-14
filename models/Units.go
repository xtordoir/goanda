package models

import (
  "fmt"
  "math"
)

type Unit float64

type TradeUnits struct {
  Precision int64
  Quantity  int64
}

func NewTradeUnits(amount float64, precision int64) TradeUnits {
  return TradeUnits{
    Precision: precision,
    Quantity: int64(amount * math.Pow(10, float64(precision))),
  }
}

func (tradeUnits TradeUnits) String() string {
  return fmt.Sprintf("%.*f", tradeUnits.Precision)
}

func (tradeUnits TradeUnits) Float64() float64 {
  unitsDecimal := float64(tradeUnits.Quantity)
  return unitsDecimal / math.Pow(10.0, float64(tradeUnits.Precision))
}

func (tradeUnits TradeUnits) Int64() int64 {
  unitsDecimal := float64(tradeUnits.Quantity)
  return int64(unitsDecimal / math.Pow(10.0, float64(tradeUnits.Precision)))
}
