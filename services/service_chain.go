package services

import (
	"coin-ant/api"
	"coin-ant/chains"
	"coin-ant/config"
	"coin-ant/dao"
	"coin-ant/models"
	"coin-ant/types"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

type ServiceChain struct {
	db *sqlca.Engine
}

func NewServiceChain(cfg *config.Config) api.ManagerApi {
	db, err := sqlca.NewEngine(cfg.DSN)
	if err != nil {
		panic(err.Error())
	}
	db.Debug(cfg.Debug)
	return &ServiceChain{
		db: db,
	}
}

func (m *ServiceChain) Run() error {
	tokenDAO := dao.NewTokenListDAO(m.db)
	dos, err := tokenDAO.QueryAll()
	if err != nil {
		return log.Errorf(err.Error())
	}
	for _, do := range dos {
		opt := m.makeOptionByModel(do)
		if err = m.runWithOption(opt); err != nil {
			log.Errorf(err.Error())
			continue
		}
	}
	var cancel = make(chan bool)
	<-cancel
	return nil
}

func (m *ServiceChain) Close() {
}

func (m *ServiceChain) makeOptionByModel(do *models.TokenListDO) (opt *types.ChainOption) {
	return &types.ChainOption{
		ChainId:      do.ChainId,
		ChainName:    do.ChainName,
		Symbol:       do.Symbol,
		ContractAddr: do.ContractAddr,
		Decimals:     do.Decimals,
		MinBalance:   do.MinBalance,
	}
}

func (m *ServiceChain) runWithOption(opt *types.ChainOption) (err error) {
	switch opt.ChainId {
	case chains.ChainIdBitcoin:
		inst := chains.NewBitcoin(m.db, opt)
		err = inst.Run()
		if err != nil {
			return log.Errorf(err.Error())
		}
	case chains.ChainIdEthereum:
		inst := chains.NewEthereum(m.db, opt)
		err = inst.Run()
		if err != nil {
			return log.Errorf(err.Error())
		}
	default:
		return log.Errorf("unknown chain id %v name %v", opt.ChainId, opt.ChainName)
	}
	return nil
}
