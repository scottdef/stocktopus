package stock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Pulls a stock quote from google finance
// Assumes the format is passed back in json
func GetQuoteGoogle(symbol string) (string, error) {

	symbol = strings.ToUpper(symbol)

	url := fmt.Sprintf("http://finance.google.com/finance/info?client=ig&q=%v", symbol)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// Read the quote into the slice
	defer resp.Body.Close()
	jsonQuote, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Google quotes start with '//' as the response
	// as well as surrounding the json with '[]'
	jsonQuote = jsonQuote[6 : len(jsonQuote)-2]

	var q interface{}
	err = json.Unmarshal(jsonQuote, &q)
	if err != nil {
		return "", err
	}

	// Type assertion
	quote, ok := q.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf(fmt.Sprintf("Quote was in unexpected format"))
	}

	// Pull the current price and the change
	l_cur := quote["l_cur"]
	c := quote["c"]

	return fmt.Sprintf("*%v*\tCurrent Price: %v\tTodays Change: %v", symbol, l_cur, c), nil
}

// Pulls a stock quote from markit on demand
// markitondeman.com
func GetQuoteMOD(symbol string) (string, error) {

	symbol = strings.ToUpper(symbol)

	url := fmt.Sprintf("http://dev.markitondemand.com/Api/v2/Quote/json?symbol=%v", symbol)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// Read the quote into the slice
	defer resp.Body.Close()
	jsonQuote, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var q interface{}
	err = json.Unmarshal(jsonQuote, &q)
	if err != nil {
		return "", err
	}

	// Type assertion
	quote, ok := q.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf(fmt.Sprintf("Quote was in unexpected format"))
	}

	// Pull the current price and the change
	l_cur := quote["LastPrice"]
	c := quote["Change"]

	return fmt.Sprintf("*%v*\tCurrent Price: %v\tTodays Change: %v", symbol, l_cur, c), nil
}
