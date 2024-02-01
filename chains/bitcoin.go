package chains

import (
	"coin-ant/api"
	"coin-ant/types"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

const (
	ChainIdBitcoin   = 0
	ChainNameBitcoin = "Bitcoin"
)

type Bitcoin struct {
	db  *sqlca.Engine
	opt *types.ChainOption
}

func NewBitcoin(db *sqlca.Engine, opt *types.ChainOption) api.ChainApi {
	return &Bitcoin{
		db:  db,
		opt: opt,
	}
}

// run service loop
func (m *Bitcoin) Run() (err error) {
	log.Infof("chain name [%s] symbol [%s] running...", m.opt.ChainName, m.opt.Symbol)
	return nil
}

func (m *Bitcoin) Close() {

}
