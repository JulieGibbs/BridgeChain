package diverclient

// AccountResp the response from thorclient
type AccountResp struct {
	Account struct {
		AccountNumber uint64 `json:"account_number,string"`
		Sequence      uint64 `json:"sequence,string"`
	} `json:"account"`
}

type AccountBalance struct {
	Balances []struct {
		Denom  string  `json:"denom"`
		Amount float64 `json:"amount"`
	} `json:"balances"`
}

// QueryResLastBlockHeights used to return the block height query
type QueryResLastBlockHeights struct {
	Block_id struct {
		Hash            string `json:"hash"`
		Part_set_header struct {
			Total uint64 `json:"total"`
			Hash  string `json:"hash"`
		} `json:"part_set_header"`
	} `json:"block_id"`
	Block struct {
		Header struct {
			Version struct {
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"version"`
			Chain_id      string `json:"chain_id"`
			Height        string `json:"height"`
			Time          string `json:"time"`
			Last_block_id struct {
				Hash            string `json:"hash"`
				Part_set_header struct {
					Total uint64 `json:"total"`
					Hash  string `json:"hash"`
				} `json:"part_set_header"`
			} `json:"last_block_id"`
			Last_commit_hash     string `json:"last_commit_hash"`
			Data_hash            string `json:"data_hash"`
			Validators_hash      string `json:"validators_hash"`
			Next_validators_hash string `json:"next_validators_hash"`
			Consensus_hash       string `json:"consensus_hash"`
			App_hash             string `json:"app_hash"`
			Last_results_hash    string `json:"last_results_hash"`
			Evidence_hash        string `json:"evidence_hash"`
			Proposer_address     string `json:"proposer_address"`
		} `json:"header"`
		Data struct {
			Txs []string `json:"txs"`
		} `json:"data"`
		Evidence struct {
			Evidence []string `json:"evidence"`
		} `json:"evidence"`
		Last_commit struct {
			Height   string `json:"height"`
			Round    uint64 `json:"round"`
			Block_id struct {
				Hash            string `json:"hash"`
				Part_set_header struct {
					Total uint64 `json:"total"`
					Hash  string `json:"hash"`
				} `json:"part_set_header"`
			} `json:"block_id"`
			Signatures []BlockSignature
		} `json:"last_commit"`
	}
}

//
type BlockSignature struct {
	Block_id_flag     string `json:"block_id_flag"`
	Validator_address string `json:"validator_address"`
	Timestamp         string `json:"timestamp"`
	Signature         string `json:"signature"`
}

//
type QueryTransactionData struct {
	Amount             string `json:"amount"`
	ConfirmedBlockHash string `json:"confirmedBlockHash"`
	Creator            string `json:"creator"`
	Index              string `json:"index"`
	OriginAddress      string `json:"originAddress"`
	OriginChain        string `json:"originChain"`
	Status             string `json:"status"`
	TargetAddress      string `json:"targetAddress"`
	TargetChain        string `json:"targetChain"`
	Time               string `json:"time"`
	Fee                string `json:"fee"`
}

//
type QueryTransactionDataList struct {
	TransactionData []QueryTransactionData
}
