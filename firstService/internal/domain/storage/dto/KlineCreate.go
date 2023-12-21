package dto

import "toscale-test-task/firstService/internal/domain/entity"

type KlineCreate struct {
	Symbol   string `json:"symbol" gorm:"<-:create, index" binding:"required"`
	Interval string `json:"interval" gorm:"<-create, indexß" binding:"required"`

	KlineOpenTime            int64   `json:"klineOpenTime" binding:"required"`
	OpenPrice                float32 `json:"openPrice" binding:"required"`
	HighPrice                float32 `json:"highPrice" binding:"required"`
	LowPrice                 float32 `json:"lowPrice" binding:"required"`
	ClosePrice               float32 `json:"closePrice" binding:"required"`
	Volume                   float32 `json:"volume" binding:"required"`
	KlineCloseTime           int64   `json:"klineCloseTime" binding:"required"`
	QuoteAssetVolume         float32 `json:"quoteAssetVolume" binding:"required"`
	NumberOfTrades           int64   `json:"numberOfTrades" binding:"required"`
	TakerBuyBaseAssetVolume  float32 `json:"takerBuyBaseAssetVolume" binding:"required"`
	TakerBuyQuoteAssetVolume float32 `json:"takerBuyQuoteAssetVolume" binding:"required"`
}

func NewKlineCreate(symbol string, interval string, klineOpenTime int64, openPrice float32, highPrice float32, lowPrice float32, closePrice float32, volume float32, klineCloseTime int64, quoteAssetVolume float32, numberOfTrades int64, takerBuyBaseAssetVolume float32, takerBuyQuoteAssetVolume float32) *KlineCreate {
	return &KlineCreate{Symbol: symbol, Interval: interval, KlineOpenTime: klineOpenTime, OpenPrice: openPrice, HighPrice: highPrice, LowPrice: lowPrice, ClosePrice: closePrice, Volume: volume, KlineCloseTime: klineCloseTime, QuoteAssetVolume: quoteAssetVolume, NumberOfTrades: numberOfTrades, TakerBuyBaseAssetVolume: takerBuyBaseAssetVolume, TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume}
}

func (c *KlineCreate) ToEntity() *entity.Kline {
	return &entity.Kline{
		Symbol:                   c.Symbol,
		Interval:                 c.Interval,
		KlineOpenTime:            c.KlineOpenTime,
		OpenPrice:                c.OpenPrice,
		HighPrice:                c.HighPrice,
		LowPrice:                 c.LowPrice,
		ClosePrice:               c.ClosePrice,
		Volume:                   c.Volume,
		KlineCloseTime:           c.KlineCloseTime,
		QuoteAssetVolume:         c.QuoteAssetVolume,
		NumberOfTrades:           c.NumberOfTrades,
		TakerBuyBaseAssetVolume:  c.TakerBuyBaseAssetVolume,
		TakerBuyQuoteAssetVolume: c.TakerBuyQuoteAssetVolume,
	}
}
