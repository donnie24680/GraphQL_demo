// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Block struct {
	Number       string   `json:"number"`
	Hash         string   `json:"hash"`
	ParentHash   string   `json:"parentHash"`
	Nonce        string   `json:"nonce"`
	Transactions []string `json:"transactions"`
}

type Query struct {
}
