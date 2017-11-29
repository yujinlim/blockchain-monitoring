package coin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// EtherchainBlockResponse etherscan response object
type EtherchainBlockResponse struct {
	Data []struct {
		Count int64 `json:"count"`
	} `json:"data"`
}

//NewCompare create new compare func
func NewCompare(coinType Type, network NetworkType) func(int64) (float64, error) {
	if coinType == BitcoinType || ((coinType == DogecoinType || coinType == LitecoinType) && network == Mainnet) {
		return compareBlockCypher(coinType, network)
	} else if coinType == DogecoinType || coinType == LitecoinType || coinType == DashType {
		compare, err := NewChainSoCompareFunc(coinType, network)
		if err != nil {
			log.Fatal(err)
		}

		return compare
	}

	return compareEthBlockCount
}

func compareEthBlockCount(current int64) (float64, error) {
	var etherResponseBody EtherchainBlockResponse
	response, err := http.Get("https://etherchain.org/api/blocks/count")
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return 0, errors.New(response.Status)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(responseBody, &etherResponseBody); err != nil {
		return 0, err
	}

	diff := etherResponseBody.Data[0].Count - current
	count := float64(diff)

	return count, nil
}

// Only uses block cypher for btc, btct, dogecoin, litecoin
func compareBlockCypher(coinType Type, network NetworkType) func(int64) (float64, error) {
	client := NewBlockCypherClient(coinType, network)

	return func(current int64) (float64, error) {
		chain, err := client.GetChain()
		if err != nil {
			log.Panic(err)
		}

		diff := int64(chain.Height) - current
		return float64(diff), nil
	}
}
