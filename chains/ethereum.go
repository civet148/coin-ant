package chains

import (
	"coin-ant/api"
	"coin-ant/types"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

const (
	ChainIdEthereum   = 1
	ChainNameEthereum = "Ethereum"
)

type Ethereum struct {
	db  *sqlca.Engine
	opt *types.ChainOption
}

func NewEthereum(db *sqlca.Engine, opt *types.ChainOption) api.ChainApi {
	return &Ethereum{
		db:  db,
		opt: opt,
	}
}

// run service loop
func (m *Ethereum) Run() (err error) {
	log.Infof("chain name [%s] symbol [%s] running...", m.opt.ChainName, m.opt.Symbol)
	return nil
}

func (m *Ethereum) Close() {

}
