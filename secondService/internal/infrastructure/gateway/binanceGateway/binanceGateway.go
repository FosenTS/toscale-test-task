package binanceGateway

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"toscale-test-task/secondService/internal/application/config"
	"toscale-test-task/secondService/internal/domain/entity"
)

type BinanceGateway interface {
	GetKlines(ctx context.Context, symbol string, interval string) ([]*entity.Kline, error)
}

type binanceGateway struct {
	config *config.BinanceConfig

	log *logrus.Entry
}

func NewBinanceGateway(config *config.BinanceConfig, log *logrus.Entry) BinanceGateway {
	return &binanceGateway{config: config, log: log}
}

func (bG *binanceGateway) GetKlines(ctx context.Context, symbol string, interval string) ([]*entity.Kline, error) {
	uString, err := url.JoinPath(bG.config.Domain, bG.config.KlinePath)
	if err != nil {
		bG.log.Errorln(err)
		return nil, err
	}

	u, err := url.Parse(uString)
	if err != nil {
		bG.log.Errorln(err)
		return nil, err
	}

	queries := url.Values{}
	queries.Set("symbol", symbol)
	queries.Set("interval", interval)

	u.RawQuery = queries.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		err := fmt.Errorf("don't create request context: %w", err)
		bG.log.Errorln(err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		err := fmt.Errorf("couldn't request binance klines: %w", err)
		bG.log.Errorln(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err := fmt.Errorf("couldn't get body: %w", err)
		bG.log.Errorln(err)
		return nil, err
	}

	externalArray := make([][]string, 0)
	internalArray := ""
	work := false
	lenBody := len(body) - 1
	for i, b := range body {
		if i == 0 || i == lenBody {
			continue
		}
		elem := string(b)
		if elem == `"` {
			continue
		}
		if elem == "[" {
			internalArray = internalArray[:0]
			work = true
			continue
		}
		if elem == "]" && string(body[i+1]) == "," {
			externalArray = append(externalArray, strings.Split(internalArray, ","))
			work = false
			continue
		}
		if work {
			internalArray += elem
		}
	}

	klines := make([]*entity.Kline, 0)
	for _, kline := range externalArray {
		klineOpenTime, err := strconv.ParseInt(kline[0], 10, 64)
		if err != nil {
			bG.log.Errorln(err)
			return nil, err
		}

		openPrice, err := strconv.ParseFloat(kline[1], 32)
		if err != nil {
			bG.log.Errorln(err)
			return nil, err
		}

		highPrice, err := strconv.ParseFloat(kline[2], 32)
		if err != nil {
			bG.log.Errorln(err)
			return nil, err
		}

		lowPrice, err := strconv.ParseFloat(kline[3], 32)
		if err != nil {
			bG.log.Errorln(err)
			return nil, err
		}

		closePrice, err := strconv.ParseFloat(kline[4], 32)
		if err != nil {
			bG.log.Errorln(err)
			return nil, err
		}

		volume, err := strconv.ParseFloat(kline[5], 32)
		if err != nil {
			bG.log.Errorln(err)
			return nil, err
		}

		klineCloseTime, err := strconv.ParseInt(kline[6], 10, 64)
		if err != nil {
			bG.log.Errorln(err)
			return nil, err
		}

		quoteAssetVolume, err := strconv.ParseFloat(kline[7], 32)
		if err != nil {
			bG.log.Errorln(err)
			return nil, err
		}

		numberOfTrades, err := strconv.ParseInt(kline[8], 10, 64)
		if err != nil {
			bG.log.Errorln(err)
			return nil, err
		}

		takerBuyBaseAssetVolume, err := strconv.ParseFloat(kline[9], 32)
		if err != nil {
			bG.log.Errorln(err)
			return nil, err
		}

		takerBuyQuoteAssetVolume, err := strconv.ParseFloat(kline[10], 32)
		if err != nil {
			bG.log.Errorln(err)
			return nil, err
		}

		klines = append(
			klines,
			entity.NewKline(
				klineOpenTime,
				float32(openPrice),
				float32(highPrice),
				float32(lowPrice),
				float32(closePrice),
				float32(volume),
				klineCloseTime,
				float32(quoteAssetVolume),
				numberOfTrades,
				float32(takerBuyBaseAssetVolume),
				float32(takerBuyQuoteAssetVolume),
			),
		)
	}

	return klines, nil
}
