package repository

type Stock struct {
	ID       string `db:"id" json:"-"`
	Stock    string `db:"stock" json:"stock,omitempty"`
	Currency string `db:"currency" json:"currency,omitempty"`
	Exchange string `db:"exchange" json:"exchange,omitempty"`
	MicCode  string `db:"mic_code" json:"mic_code,omitempty"`
	Price    string `db:"price" json:"price,omitempty"`
}
