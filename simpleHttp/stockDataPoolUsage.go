package simpleHttp

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-learinng/simpleHttp/connectionPool"
	"go-learinng/simpleHttp/repository"
	"go-learinng/simpleHttp/twelveDataClient"
	"go-learinng/simpleHttp/updatePool"
	"log"
	"time"
)

const (
	TIMEOUT       = 1 * time.Minute
	WORKERS_COUNT = 8
)

var stocks = map[string]bool{
	"AAPL":  true,
	"MSFT":  true,
	"GOOG":  true,
	"AMZN":  true,
	"BRK.A": true,
	"NVDA":  true,
	"META":  true,
	"TSLA":  true,
	"V":     true,
	"UNH":   true,
	"XOM":   true,
	"JNJ":   true,
	"WMT":   true,
	"LLY":   true,
	"JPM":   true,
	"PG":    true,
	"MA":    true,
	"MRK":   true,
	"CVX":   true,
	"HD":    true,
	"KO":    true,
	"PEP":   true,
	"ORCL":  true,
	"AVGO":  true,
	"ABBV":  true,
	"COST":  true,
	"MCD":   true,
	"BAC":   true,
	"PFE":   true,
	"TMO":   true,
	"CRM":   true,
	"ABT":   true,
	"CSCO":  true,
	"NKE":   true,
	"ACN":   true,
	"LIN":   true,
	"TMUS":  true,
	"DIS":   true,
	"DHR":   true,
	"CMCSA": true,
	"VZ":    true,
	"NEE":   true,
	"ADBE":  true,
	"AMD":   true,
	"NFLX":  true,
	"PM":    true,
	"TXN":   true,
	"UPS":   true,
	"BMY":   true,
	"WFC":   true,
	"RTX":   true,
	"MS":    true,
	"HON":   true,
	"AMGN":  true,
	"T":     true,
	"SBUX":  true,
	"UNP":   true,
	"LOW":   true,
	"INTC":  true,
	"BA":    true,
	"COP":   true,
	"INTU":  true,
	"MDT":   true,
	"QCOM":  true,
	"SPGI":  true,
	"LMT":   true,
	"IBM":   true,
	"DE":    true,
	"AXP":   true,
	"ELV":   true,
	"SYK":   true,
	"CAT":   true,
	"GE":    true,
	"ISRG":  true,
	"GS":    true,
	"MDLZ":  true,
	"BX":    true,
	"AMAT":  true,
	"GILD":  true,
	"BKNG":  true,
	"BLK":   true,
	"NOW":   true,
	"ADI":   true,
	"TJX":   true,
	"MMC":   true,
	"VRTX":  true,
	"SCHW":  true,
	"C":     true,
	"CVS":   true,
	"ADP":   true,
	"ZTS":   true,
	"CB":    true,
	"MO":    true,
	"REGN":  true,
	"SO":    true,
	"SHOP":  true,
	"PGR":   true,
	"UBER":  true,
	"BSX":   true,
	"HCA":   true,
}

func StockDataPoolUsage(ctx context.Context) context.Context {
	cp := connectionPool.New(ctx, WORKERS_COUNT, TIMEOUT, 10*time.Second)
	jobs := make(chan updatePool.Job, WORKERS_COUNT)
	results := make(chan updatePool.Result, WORKERS_COUNT)
	workerPool := updatePool.New(ctx, WORKERS_COUNT, TIMEOUT, results)
	repo := repository.New(ctx.Value("db").(*sqlx.DB))
	ctxWithRepo := context.WithValue(ctx, "repo", repo)

	go processResults(repo, results)

	go func() {
		removeStoredSymbols(repo)
		generateJobs(cp, workerPool)
		close(jobs)
		workerPool.Stop()
		close(results)
	}()
	return ctxWithRepo
}

func removeStoredSymbols(repo *repository.StockRepository) {
	stored, err := repo.GetAll()
	if err != nil {
		log.Panic(err)
	}
	for _, stock := range stored {
		if _, ok := stocks[stock.Stock]; ok {
			delete(stocks, stock.Stock)
		}
	}
}

func processResults(repo *repository.StockRepository, results chan updatePool.Result) {
	go func() {
		for result := range results {
			if result.Error == nil {
				stock := result.Result.(*twelveDataClient.TimeSeries)
				lastValue := stock.Values[len(stock.Values)-1]
				log.Printf("%s | %s | %s | %s | %s | %s\n", stock.Meta.Symbol, stock.Meta.Currency, lastValue.Datetime, lastValue.Close, lastValue.Low, lastValue.High)
				err := repo.Add(repository.Stock{
					Stock:    stock.Meta.Symbol,
					Currency: stock.Meta.Currency,
					MicCode:  stock.Meta.MicCode,
					Exchange: stock.Meta.Exchange,
					Price:    lastValue.Close,
				})
				if err != nil {
					log.Println("Error at saving!", err)
				}
			} else {
				log.Println("Error!", result.Error)
			}
		}
	}()
}

func generateJobs(cp *connectionPool.ConnectionPool, pool *updatePool.Pool) {
	for stock := range stocks {
		stock := stock
		pool.Push(updatePool.Job{Execute: func() (interface{}, error) {
			return cp.GetHistoricDataForSymbol(stock)
		}})
	}
}
