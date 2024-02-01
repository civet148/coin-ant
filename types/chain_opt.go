package types

import "github.com/civet148/sqlca/v2"

type ChainOption struct {
	ChainId      int64
	ChainName    string
	Symbol       string
	ContractAddr string
	Decimals     int32
	MinBalance   sqlca.Decimal
}
