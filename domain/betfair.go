package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/rhiadc/betcontrol/config"
	"github.com/rhiadc/betcontrol/infra"
)

type legacyData struct {
	AvgPrice        float64   `json:"avgPrice"`
	BetSize         float64   `json:"betSize"`
	BetType         string    `json:"betType"`
	BetCategoryType string    `json:"betCategoryType"`
	CommissionRate  string    `json:"commissionRate"`
	EventID         int       `json:"eventId"`
	EventTypeID     int       `json:"eventTypeId"`
	FullMarketName  string    `json:"fullMarketName"`
	GrossBetAmount  float64   `json:"grossBetAmount"`
	MarketName      string    `json:"marketName"`
	MarketType      string    `json:"marketType"`
	PlacedDate      time.Time `json:"placedDate"`
	SelectionID     int       `json:"selectionId"`
	StartDate       time.Time `json:"startDate"`
	TransactionType string    `json:"transactionType"`
	TransactionID   int       `json:"transactionId"`
	WinLose         string    `json:"winLose"`
	AvgPriceRaw     float64   `json:"avgPriceRaw"`
}

type AccountStatement struct {
	RefID string `json:"refId"`
	//ItemDate      time.Time     `json:"itemDate"`
	Amount     float64    `json:"amount"`
	Balance    float64    `json:"balance"`
	LegacyData legacyData `json:"legacyData"`
	ItemClass  string     `json:"itemClass"`
}

type Response struct {
	AccountStatement []AccountStatement `json:"accountStatement"`
	MoreAvailable    bool               `json:"moreAvailable"`
}

type Betfair struct {
	env *config.Env
}

func NewBetfairService(env *config.Env) *Betfair {
	return &Betfair{env: env}
}

func (b *Betfair) CallBetfairAPI() string {
	appKey := b.env.AppKey
	sessionToken := b.env.ExternalApiToken
	host := b.env.Host
	fmt.Println(b.env)

	jsonBody := []byte(`{ "itemDateRange": { "from": "2023-11-22T03:00:00Z" }, "includeItem": "ALL" }`)
	bodyReader := bytes.NewReader(jsonBody)

	req, _ := http.NewRequest(http.MethodPost, host, bodyReader)

	req.Header.Add("X-Authentication", sessionToken)
	req.Header.Add("X-Application", appKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var con Response

	err = json.Unmarshal(body, &con)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(con)

	for _, v := range con.AccountStatement {
		if result := v.LegacyData.WinLose; result != "RESULT_NOT_APPLICABLE" {
			market := strings.Split(v.LegacyData.FullMarketName, "/")
			//dataToBePlaces := []string{market[0], market[1], market[2], fmt.Sprintf("%f", v.LegacyData.AvgPrice), result}
			fmt.Println(market[0], "|", market[1], "|", market[2], "|", v.LegacyData.AvgPrice, "|", result)
		}

	}

	names := map[string]string{"A1": "Date", "B1": "Event", "C1": "Market", "D1": "Odd", "E1": "Result"}
	columnName := infra.ColumnName{Column: names}
	sheet := infra.NewSheet("teste", "planilha.xlsx")
	sheet.CreateSpreadSheet(columnName)
	res.Body.Close()
	return "it's working"
}
