package main

import (
	"os"    
	"fmt"
    "encoding/json"
	"github.com/headzoo/surf"
    "github.com/joho/godotenv"
    "github.com/PuerkitoBio/goquery"
)

type CafeDeCrieCard struct {
    CardId string `json:"card_id"`
    Balance string `json:"balance"`
    ExpirationDate string `json:"expiration_date"`
}

const ENDPOINT = "https://rapico.jp/R0TW1TPRC/PcCard/login"

func init() {
    
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        os.Exit(1)
    }
}

func main() {

	bow := surf.NewBrowser()
	err := bow.Open(ENDPOINT)

	if err != nil {
		fmt.Println(err.Error())
	}

	fm, err := bow.Form("form")
	if err != nil {
		fmt.Println(err.Error())
	}

	fm.Input("card_id", os.Getenv("CARD_ID"))
	fm.Input("pin_id", os.Getenv("PIN_ID"))

	if fm.Submit() != nil {
		fmt.Println(err.Error())
	}

	dom := bow.Dom()
    trs := dom.Find("table.list1 tr td")
    results := []string{}
    trs.Each(func(_ int, s *goquery.Selection) {
        results = append(results, s.Text())
    })

    var card CafeDeCrieCard
    card.CardId = results[0]
    card.Balance = results[1]
    card.ExpirationDate = results[2]

    bytes, err := json.Marshal(card)
    if err != nil {
        fmt.Println("Error: ", err)
    }
    
    fmt.Println(string(bytes))
    
}
