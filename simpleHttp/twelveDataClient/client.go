package twelveDataClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-learinng/simpleHttp/utils"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type TwelveDataClient struct {
	c      *http.Client
	apiKey string
}

func NewClient(timeout time.Duration) *TwelveDataClient {
	return &TwelveDataClient{
		c: &http.Client{
			Timeout:   timeout,
			Transport: utils.NewRequestLoggingTransport(http.DefaultTransport),
		},
		apiKey: os.Getenv("TWELVEDATA_API_KEY"),
	}
}

func (c *TwelveDataClient) getUrl(resource string, params *url.Values) string {
	baseURL := "https://api.twelvedata.com"
	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	result := u.String()
	return result
}

func (c *TwelveDataClient) GetHistoricDataForSymbol(symbol string) (*TimeSeries, error) {
	params := url.Values{}
	params.Add("apikey", c.apiKey)
	params.Add("symbol", symbol)
	params.Add("interval", "1day")
	params.Add("start_date", "2023-01-01")
	response, err := c.c.Get(c.getUrl("time_series", &params))
	if err != nil {
		return nil, err
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		var result ErrorResponse
		if err = json.Unmarshal(responseBody, &result); err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("error within request: %d code, message: %s", result.Code, result.Message))
	}
	var results TimeSeries
	if err = json.Unmarshal(responseBody, &results); err != nil {
		return nil, err
	}
	if results.Status == "error" {
		return nil, errors.New("unknown error in response: " + string(responseBody))
	}
	return &results, nil
}
