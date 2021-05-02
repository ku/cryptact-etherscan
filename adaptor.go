package main

import (
	"github.com/ku/cryptact-etherscan/csv"
	"github.com/nanmu42/etherscan-api"
	"math/big"
)

type Transaction struct {
	tx etherscan.NormalTx
}

var prec uint = 256
var wei *big.Float = func() *big.Float {
	wei := big.NewInt(10)
	wei.Exp(wei, big.NewInt(int64(18)), nil)
	fwei := big.Float{}
	fwei.SetPrec(prec)
	fwei.SetInt(wei)
	return &fwei
}()

func NewCryptactLog(source string, tx *etherscan.NormalTx) *csv.CryptactLog {

	gasUsed := big.NewInt(int64(tx.GasUsed))
	price := tx.GasPrice.Int()

	gasUsed.Mul(gasUsed, price)
	txFee := big.Float{}
	txFee.SetPrec(prec)
	txFee.SetInt(gasUsed)
	txFee.Quo(&txFee, wei)

	var action string
	volume := big.NewFloat(0)
	fee := big.NewFloat(0)

	if tx.Value.Int().Int64() == 0 {
		action = "PAY"
		volume.Set(&txFee)
	} else {
		action = "SENDFEE"

		volume.SetInt(tx.Value.Int())
		volume.Quo(volume, wei)

		fee.Set(&txFee)
	}

	return &csv.CryptactLog{
		Timestamp:    tx.TimeStamp.Time(),
		Action:       action,
		Source:       source,
		Base:         "ETH",
		DerivType:    "",
		DerivDetails: "",
		Volume:       volume,
		Price:        big.NewFloat(0),
		Counter:      "",
		Fee:          fee,
		FeeCcy:       "ETH",
		Comment:      tx.Hash,
	}
}
