package entity

type Kline struct {
	KlineOpenTime            int64
	OpenPrice                float32
	HighPrice                float32
	LowPrice                 float32
	ClosePrice               float32
	Volume                   float32
	KlineCloseTime           int64
	QuoteAssetVolume         float32
	NumberOfTrades           int64
	TakerBuyBaseAssetVolume  float32
	TakerBuyQuoteAssetVolume float32
}

func NewKline(klineOpenTime int64, openPrice float32, highPrice float32, lowPrice float32, closePrice float32, volume float32, klineCloseTime int64, quoteAssetVolume float32, numberOfTrades int64, takerBuyBaseAssetVolume float32, takerBuyQuoteAssetVolume float32) *Kline {
	return &Kline{KlineOpenTime: klineOpenTime, OpenPrice: openPrice, HighPrice: highPrice, LowPrice: lowPrice, ClosePrice: closePrice, Volume: volume, KlineCloseTime: klineCloseTime, QuoteAssetVolume: quoteAssetVolume, NumberOfTrades: numberOfTrades, TakerBuyBaseAssetVolume: takerBuyBaseAssetVolume, TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume}
}
