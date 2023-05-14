package server

import (
	"context"
	"encoding/json"
	"go-learinng/simpleHttp/repository"
	"go-learinng/simpleHttp/utils"
	"io"
	"log"
	"net/http"
)

func New(ctx context.Context) {
	stockRepository := ctx.Value("repo").(*repository.StockRepository)
	handler := stockHandler{stockRepository}
	http.Handle("/stock", utils.NewLoggingHandler(&handler))

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Panic(err)
	}
}

type stockHandler struct {
	repo *repository.StockRepository
}

func (h *stockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getStocks(w)
	case http.MethodPost:
		h.addStock(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *stockHandler) getStocks(w http.ResponseWriter) {
	stocks, err := h.repo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	jsonString, err := json.Marshal(stocks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(jsonString)
	if err != nil {
		log.Println("Error on sending response!", err)
	}
}

func (h *stockHandler) addStock(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error on reading request!")
		return
	}
	if len(requestBody) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	var stock repository.Stock
	if err := json.Unmarshal(requestBody, &stock); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error on unmarshalling request from %s!\n", requestBody)
		return
	}
	if err = h.repo.Add(stock); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error on additing stock %+v!\n%v\n", stock, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
