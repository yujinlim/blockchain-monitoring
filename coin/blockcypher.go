package coin

import "github.com/blockcypher/gobcy"

//NewBlockCypherClient create new blockcypher client
func NewBlockCypherClient(coinType Type, network NetworkType) *gobcy.API {
	bc := gobcy.API{}

	if coinType == BitcoinType {
		bc.Coin = "btc"
	} else if coinType == DogecoinType {
		bc.Coin = "doge"
	} else if coinType == LitecoinType {
		bc.Coin = "ltc"
	}

	bc.Chain = string(network)

	return &bc
}
