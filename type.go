package virclerpc

// Block ...
type Block struct {
	Hash              string   `json:"hash"`
	Confirmations     int32    `json:"confirmations"`
	Strippedsize      int32    `json:"strippedsize"`
	Size              int32    `json:"size"`
	Weight            int32    `json:"weight"`
	Height            int32    `json:"height"`
	Version           int32    `json:"version"`
	VersionHex        string   `json:"versionHex"`
	Merkleroot        string   `json:"merkleroot"`
	Witnessmerkleroot string   `json:"witnessmerkleroot"`
	Txs               []string `json:"tx"`
	Time              int32    `json:"time"`
	Mediantime        int32    `json:"mediantime"`
	Nonce             int32    `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float32  `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	Txn               int32    `json:"nTx"`
	PreviousBlockhash string   `json:"previousblockhash"`
	NextBlockhash     string   `json:"nextblockhash"`
}

// Transaction ...
type Transaction struct {
	Txid     string `json:"txid"`
	Hash     string `json:"hash"`
	Size     int32  `json:"size"`
	Vsize    int32  `json:"vsize"`
	Version  int32  `json:"version"`
	Locktime int32  `json:"locktime"`
	Vins     Vin
	Vouts    []Vout
}

// Vin ...
type Vin map[int]interface{}

// Vout ...
type Vout struct {
	Value        float32 `json:"value"`
	N            int32   `json:"n"`
	ScriptPubKey ScriptPubKey
}

// ScriptPubKey ...
type ScriptPubKey struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int32    `json:"reqSigs"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses"`
}

// VinTransaction is a struct of Vin(inputs) with normal transaction.
type VinTransaction struct {
	Txid        string `json:"txid"`
	Vout        int32  `json:"vout"`
	ScriptSig   ScriptSig
	TxinWitness []string `json:"txinWitness"`
	Sequence    int64    `json:"Sequence"`
}

// VinCoinbaseTransaction is a struct of Vins(inputs) with mining.
type VinCoinbaseTransaction struct {
	Coinbase string `json:"coinbase"`
	Sequence int64  `json:"sequence"`
}

// ScriptSig ...
type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type MasterNodeStatus struct {
	Outpoint string `json:"outpoint"`
	Service  string `json:"service"`
	State    string `json:"state"`
	Status   string `json:"status"`
}
