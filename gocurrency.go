// http://play.golang.org/p/ebVIqodGPd

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
)

const (
	API = "http://api.fixer.io/latest"
)

var (
	val   string
	base  string
)

func help() {
	fmt.Println("GoCurrency is used to quickly check the conversion rates")
	flag.PrintDefaults()
}

func exitHelp(code int) {
	help()
	os.Exit(code)
}

func init() {
	flag.Usage = help

	flag.StringVar(&base, "base", "USD", "The base currency to quote a currency against (default USD)")
	flag.Parse()

	if base == "" {
		exitHelp(3)
	}

	val = flag.Arg(0)
}

func sendRequest() {
	v := url.Values{}
	if len(base) > 0 {
		v.Add("base", base)
	}
	if len(val) > 0 {
		v.Add("symbols", val)
	}
	params := v.Encode()

	resp, err := http.Get(API + "?" + params)
	if err != nil {
		fmt.Println("Failed to get data", err)
		os.Exit(3)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Not 200 found")
		os.Exit(3)
	}

	handleResponse(resp.Body)
}

func sortDataKeys(m map[string]float64) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

type DataType struct {
	Base  string
	Date  string
	Rates map[string]float64
}

func handleResponse(r io.ReadCloser) {
	var data DataType

	err := json.NewDecoder(r).Decode(&data)
	if err != nil {
		fmt.Println("Error reading data", err)
		os.Exit(3)
	}

	fmt.Printf("Base currency: %s, Date: %s\n", data.Base, data.Date)

	fmt.Printf("%-8s| %-8s\n", "Symbol", "Rate")
	fmt.Printf("%-8s|-%-8.3s\n", "--------", "--------")

	keys := sortDataKeys(data.Rates)
	for _, k := range keys {
		fmt.Printf("%-8s| %-8.3f\n", k, data.Rates[k])
	}
}

func main() {
	sendRequest()
}