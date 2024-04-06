package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/HsiaoCz/code-beast/probcx/types"
	"github.com/zeebo/assert"
)

func TestHeader_Encdoe_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PreBlock:  types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     989394,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}

func TestBlock_Decode_Encode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PreBlock:  types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     989394,
		},
		Transactions: nil,
	}
	buf := &bytes.Buffer{}
	assert.Nil(t, b.EecodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))

	assert.Equal(t, b, bDecode)
}
