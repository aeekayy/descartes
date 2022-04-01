package dal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/aeekayy/descartes/server/pkg/config"
)

// DBConn represent database connection
type DBConn struct {
	Config     *config.DBConfig `yaml:"config",json:"config"`
	Connection *sql.DB          `yaml:"connection",json:"connection"`
}

// New create a new database connection
func New(dbConfig *config.DBConfig) (*DBConn, error) {
	dbcstr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name)

	// open database
	db, err := sql.Open("postgres", dbcstr)
	if err != nil {
		fmt.Printf("error connecting to the database: %s", err)
		return nil, err
	}

	dbConn := DBConn{
		Config:     dbConfig,
		Connection: db,
	}

	return &dbConn, nil
}

func (d *DBConn) CreateStockOverview() {
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
