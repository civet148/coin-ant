package types

type ChainOption struct {
	ChainId      int64
	ChainName    string
	Symbol       string
	ContractAddr string
	Decimals     int32
	MinBalance   int64
}
