package coin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Data etherscan data object
type Data struct {
	Count int64 `json:"count"`
}

// EtherchainBlockResponse etherscan response object
type EtherchainBlockResponse struct {
	Data []Data `json:"data"`
}

func compareDogecoinBlockCount(current int64) (float64, error) {
	var blockCount int64
	response, err := http.Get("https://dogechain.info/chain/Dogecoin/q/getblockcount")
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

	if err := json.Unmarshal(responseBody, &blockCount); err != nil {
		return 0, err
	}

	diff := blockCount - current
	return float64(diff), nil
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
