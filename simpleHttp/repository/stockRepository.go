package repository

import (
	"github.com/jmoiron/sqlx"
)

type StockRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) Add(stock Stock) error {
	const insertStock = `INSERT INTO stocks (stock, currency, exchange, mic_code, price) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(insertStock, stock.Stock, stock.Currency, stock.Exchange, stock.MicCode, stock.Price)
	return err
}

func (r *StockRepository) GetById(id int64) (Stock, error) {
	const selectById = `SELECT * FROM stocks WHERE id = $1`
	var result Stock
	err := r.db.Select(&result, selectById, id)
	return result, err
}

func (r *StockRepository) GetBySymbol(symbol string) (Stock, error) {
	const selectById = `SELECT * FROM stocks WHERE stock = $1`
	var result Stock
	err := r.db.Select(&result, selectById, symbol)
	return result, err
}

func (r *StockRepository) GetAll() ([]Stock, error) {
	const selectAll = `SELECT * FROM stocks`
	var result []Stock
	err := r.db.Select(&result, selectAll)
	return result, err
}
