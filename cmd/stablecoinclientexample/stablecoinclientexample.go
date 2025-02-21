package stablecoinclientexample

import (
	"github.com/Turing-Chu/huobi_golang/config"
	"github.com/Turing-Chu/huobi_golang/logging/applogger"
	"github.com/Turing-Chu/huobi_golang/pkg/client"
)
func RunAllExamples() {
	getExchangeRate()
	exchangeStableCoin()
}

func getExchangeRate() {
	client := new(client.StableCoinClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.GetExchangeRate("pax", "1000", "buy")
	if err != nil {
		applogger.Error("Get stable coin error: %s", err)
	} else {
		applogger.Info("Get stable coin success: %+v", *resp.Data)
	}
}

func exchangeStableCoin() {
	client := new(client.StableCoinClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp, err := client.ExchangeStableCoin("123")
	if err != nil {
		applogger.Error("Exchange stable coin error: %s", err)
	} else {
		applogger.Info("Exchange stable coin success: %+v", *resp.Data)
	}
}
