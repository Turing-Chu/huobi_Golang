package market

import (
	"github.com/Turing-Chu/huobi_golang/pkg/model/base"
	"github.com/shopspring/decimal"
)

type SubscribeMarketByPriceResponse struct {
	base.WebSocketResponseBase
	Tick *MarketByPrice
	Data *MarketByPrice
}

type MarketByPrice struct {
	SeqNum     int64               `json:"seqNum"`
	PrevSeqNum int64               `json:"prevSeqNum"`
	Bids       [][]decimal.Decimal `json:"bids"`
	Asks       [][]decimal.Decimal `json:"asks"`
}
