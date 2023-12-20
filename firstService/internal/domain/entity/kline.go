package entity

type Kline struct {
	Symbol   string `json:"symbol" binding:"required"`
	Interval string `json:"interval" binding:"required"`

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

func NewKline(symbol string, interval string, klineOpenTime int64, openPrice float32, highPrice float32, lowPrice float32, closePrice float32, volume float32, klineCloseTime int64, quoteAssetVolume float32, numberOfTrades int64, takerBuyBaseAssetVolume float32, takerBuyQuoteAssetVolume float32) *Kline {
	return &Kline{Symbol: symbol, Interval: interval, KlineOpenTime: klineOpenTime, OpenPrice: openPrice, HighPrice: highPrice, LowPrice: lowPrice, ClosePrice: closePrice, Volume: volume, KlineCloseTime: klineCloseTime, QuoteAssetVolume: quoteAssetVolume, NumberOfTrades: numberOfTrades, TakerBuyBaseAssetVolume: takerBuyBaseAssetVolume, TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume}
}
