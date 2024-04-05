package core

import "github.com/HsiaoCz/code-beast/probcx/types"

type Header struct {
	Version   uint32
	PreBlock  types.Hash
	Timestamp uint64
	Height    uint32
	Nonce     uint32
}

type Block struct {
	Header
	Transactions []Transaction
}
