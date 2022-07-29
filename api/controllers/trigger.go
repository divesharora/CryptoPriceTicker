package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/divesharora/KryptoBackendTask/api/db"
	"github.com/divesharora/KryptoBackendTask/api/models"
	"github.com/gorilla/websocket"
)

func InitializeTrigger() {
	alertDB := db.GetDB()
	alertDB.AutoMigrate(&models.Alert{})
	c, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws/!ticker@arr", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// create the input channel
	updateChannel := make(chan []models.PriceData)
	tickerFinder := make(chan models.Alert)

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				break
			}
			var trade []models.PriceData
			json.Unmarshal(message, &trade)

			updateChannel <- trade
		}
		close(updateChannel)
	}()
	go func() {
		for trade := range updateChannel {
			alerts := []models.Alert{}
			alertDB.Raw("SELECT * FROM alerts WHERE status = ?", 0).Scan(&alerts)

			for _, alert := range alerts {
				for _, trade := range trade {
					if trade.Symbol == alert.Symbol && trade.LastPrice > alert.Price {
						tickerFinder <- alert
					}
				}
			}
		}
		close(tickerFinder)
	}()

	for trade := range tickerFinder {
		SendAlert(trade.Email)
		alertDB.Model(&trade).Update("status", 1)
	}
}
