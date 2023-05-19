package simpleHttp

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-learinng/simpleHttp/server"
	"log"
	"os"
	"sync"
)

func Examples() {
	db, err := sqlx.Connect("postgres", "host=127.0.0.1 port=5432 dbname=postgres user=stocks_app sslmode=disable password="+os.Getenv("DB_PASSWORD"))
	if err != nil {
		log.Panic(err)
	}
	stopWg := &sync.WaitGroup{}
	ctx := context.WithValue(context.Background(), "stopWg", stopWg)
	ctx = context.WithValue(ctx, "db", db)
	ctx = StockDataPoolUsage(ctx)
	server.New(ctx)
	stopWg.Wait()
}
