package coin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	dash            string = "DASH"
	dashTestnet     string = "DASHTEST"
	dogecoinTestnet string = "DOGETEST"
	litecoinTestnet string = "LTCTEST"
)

const baseURL string = "https://chain.so/api/v2/get_info/"

// ChainSONetworkInfoResponse chain.so network info api response body
type ChainSONetworkInfoResponse struct {
	Data struct {
		Blocks int64 `json:"blocks"`
	} `json:"data"`
}

// NewChainSoCompareFunc get compare function from chain.so
// only for dash, dogecoin testnet, litecoin testnet
func NewChainSoCompareFunc(coinType Type, network NetworkType) (func(int64) (float64, error), error) {
	var blockchainType string

	if coinType == LitecoinType {
		blockchainType = litecoinTestnet
	} else if coinType == DogecoinType {
		blockchainType = dogecoinTestnet
	} else if coinType == DashType && network == Testnet {
		blockchainType = dashTestnet
	} else if coinType == DashType && network == Mainnet {
		blockchainType = dash
	} else {
		return nil, errors.New("Does not support type specified")
	}

	url := baseURL + blockchainType

	return func(current int64) (float64, error) {
		var chainResponseBody ChainSONetworkInfoResponse
		response, err := http.Get(url)
		if err != nil {
			return 0, err
		}

		if response.StatusCode != 200 {
			return 0, errors.New(response.Status)
		}

		responseBody, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(responseBody, &chainResponseBody); err != nil {
			return 0, err
		}

		diff := chainResponseBody.Data.Blocks - current
		return float64(diff), nil
	}, nil
}
