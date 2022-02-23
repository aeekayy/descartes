package cron

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	cronV3 "github.com/robfig/cron/v3"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/aeekayy/descartes/pkg/config"
)

const (
	alphaVantageAPI = "https://www.alphavantage.co/query"
)

type CronServer struct {
	cron   *cronV3.Cron
	Config *config.CronConfig
}

type StockOverview struct {
	Symbol                     string  `json:"Symbol"`
	AssetType                  string  `json:"AssetType"`
	Name                       string  `json:"Name"`
	Description                string  `json:"Description"`
	CIK                        string  `json:"CIK"`
	Exchange                   string  `json:"Exchange"`
	Currency                   string  `json:"Currency"`
	Country                    string  `json:"Country"`
	Sector                     string  `json:"Sector"`
	Industry                   string  `json:"Industry"`
	Address                    string  `json:"Address"`
	FiscalYearEnd              string  `json:"FiscalYearEnd"`
	LatestQuarter              string  `json:"LatestQuarter"`
	MarketCapitalization       int64   `json:"MarketCapitalization,string"`
	EBITDA                     string  `json:"EBITDA"`
	PERatio                    string  `json:"PERatio"`
	PEGRatio                   float32 `json:"PEGRatio,string"`
	BookValue                  float32 `json:"BookValue,string"`
	DividendPerShare           string  `json:"DividendPerShare"`
	DividendYield              string  `json:"DividendYield"`
	EPS                        float32 `json:"EPS,string"`
	RevenuePerShareTTM         float32 `json:"RevenuePerShareTTM,string"`
	ProfitMargin               float32 `json:"ProfitMargin,string"`
	OperatingMarginTTM         float32 `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM          float32 `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM          float32 `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                 int64   `json:"RevenueTTM,string"`
	GrossProfitTTM             int64   `json:"GrossProfitTTM,string"`
	DilutedEPSTTM              float32 `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY float32 `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY  float32 `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice         float32 `json:"AnalystTargetPrice,string"`
	TrailingPE                 string  `json:"TrailingPE"`
	ForwardPE                  string  `json:"ForwardPE"`
	PriceToSalesRatioTTM       float32 `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio           float32 `json:"PriceToBookRatio,string"`
	EVToRevenue                float32 `json:"EVToRevenue,string"`
	EVToEBITDA                 string  `json:"EVToEBITDA"`
	Beta                       string  `json:"Beta"`
	FiftyTwoWeekHigh           float32 `json:"52WeekHigh,string"`
	FiftyTwoWeekLow            float32 `json:"52WeekLow,string"`
	FiftyDayMovingAverage      float32 `json:"50DayMovingAverage,string"`
	TwoHundredDayMovingAverage float32 `json:"200DayMovingAverage,string"`
	SharesOutstanding          int64   `json:"SharesOutstanding,string"`
	DividendDate               string  `json:"DividendDate"`
	ExDividendDate             string  `json:"ExDividendDate"`
}

func New(cronConfig *config.CronConfig) *CronServer {
	newCron := cronV3.New(cronV3.WithSeconds())

	return &CronServer{
		cron:   newCron,
		Config: cronConfig,
	}
}

func (c *CronServer) Start() error {
	c.cron.Start()

	return nil
}

func (c *CronServer) Init() error {
	c.cron.AddFunc("@every 6h", c.GetStockOverview)

	return nil
}

func (c *CronServer) GetStockOverview() {
	symbols := []string{
		"VERI",
		"PLTR",
		"IBM",
		"FB",
	}
	client := http.Client{}

	for _, symbol := range symbols {
		var stockOverview StockOverview

		url := fmt.Sprintf("%s?function=OVERVIEW&symbol=%s&apikey=%s", alphaVantageAPI, symbol, c.Config.API.Keys.AlphaVantage)
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			fmt.Printf("error creating request: %s", err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("error sending request: %s", err)
			return
		}

		// parse the result
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("error reading the body: %s", err)
			return
		}

		fmt.Printf("%+v", string(body))

		if err := json.Unmarshal(body, &stockOverview); err != nil {
			fmt.Printf("error unmarshaling stock overview: %s, %s", err, string(body))
			return
		}

		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Config.DB.Host, c.Config.DB.Port, c.Config.DB.Username, c.Config.DB.Password, c.Config.DB.Name)

		// open database
		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			fmt.Printf("error connecting to the database: %s", err)
			return
		}

		// close database
		defer db.Close()

		// check db
		err = db.Ping()
		if err != nil {
			fmt.Printf("error with database ping: %s", err)
			return
		}

		q := fmt.Sprintf("INSERT INTO stock_overview_rawdata(symbol,assettype,name,description,cik,exchange,currency,country,sector,industry,address,fiscalyearend,latestquarter,marketcapitalization,ebitda,peratio,pegratio,bookvalue,dividendpershare,dividendyield,eps,revenuepersharettm,profitmargin,operatingmarginttm,returnonassetsttm,returnonequityttm,revenueTTM,grossprofitttm,dilutedepsttm,quarterlyearningsgrowthyoy,quarterlyrevenuegrowthyoy,analysttargetprice,trailingpe,forwardpe,pricetosalesratiottm,pricetobookratio,evtorevenue,evtoebitda,beta,fiftytwoweekhigh,fiftytwoweeklow,fiftydaymovingaverage,twohundreddaymovingaverage,sharesoutstanding,dividenddate,exdividenddate) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d, '%s', '%s', %.2f, %.2f, '%.s', '%s', %.2f, %.2f, %.5f, %.5f, %.5f, %.5f, %d, %d, %.2f, %.5f, %.5f, %.2f, '%s', '%s', %.5f, %.2f, %.5f, '%s', '%s', %.2f, %.2f, %.2f, %.2f, %d, '%s', '%s')", stockOverview.Symbol, stockOverview.AssetType, stockOverview.Name, stockOverview.Description, stockOverview.CIK, stockOverview.Exchange, stockOverview.Currency, stockOverview.Country, stockOverview.Sector, stockOverview.Industry, stockOverview.Address, stockOverview.FiscalYearEnd, stockOverview.LatestQuarter, stockOverview.MarketCapitalization, stockOverview.EBITDA, stockOverview.PERatio, stockOverview.PEGRatio, stockOverview.BookValue, stockOverview.DividendPerShare, stockOverview.DividendYield, stockOverview.EPS, stockOverview.RevenuePerShareTTM, stockOverview.ProfitMargin, stockOverview.OperatingMarginTTM, stockOverview.ReturnOnAssetsTTM, stockOverview.ReturnOnEquityTTM, stockOverview.RevenueTTM, stockOverview.GrossProfitTTM, stockOverview.DilutedEPSTTM, stockOverview.QuarterlyEarningsGrowthYOY, stockOverview.QuarterlyRevenueGrowthYOY, stockOverview.AnalystTargetPrice, stockOverview.TrailingPE, stockOverview.ForwardPE, stockOverview.PriceToSalesRatioTTM, stockOverview.PriceToBookRatio, stockOverview.EVToRevenue, stockOverview.EVToEBITDA, stockOverview.Beta, stockOverview.FiftyTwoWeekHigh, stockOverview.FiftyTwoWeekLow, stockOverview.FiftyDayMovingAverage, stockOverview.TwoHundredDayMovingAverage, stockOverview.SharesOutstanding, stockOverview.DividendDate, stockOverview.ExDividendDate)
		_, err = db.Exec(q)
		if err != nil {
			fmt.Printf("error inserting stock overview record: %s \n %+v", err, stockOverview)
			return
		}
	}
}
