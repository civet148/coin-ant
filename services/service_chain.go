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
		if err = m.runWithOptionAsync(opt); err != nil {
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

func (m *ServiceChain) runWithOptionAsync(opt *types.ChainOption) (err error) {
	var instance api.ChainApi
	switch opt.ChainId {
	case chains.ChainIdBitcoin:
		instance = chains.NewBitcoin(m.db, opt)

	case chains.ChainIdEthereum:
		instance = chains.NewEthereum(m.db, opt)
	default:
		return log.Errorf("unknown chain id [%v] name [%v]", opt.ChainId, opt.ChainName)
	}
	go func() {
		err = instance.Run()
		if err != nil {
			log.Errorf("run chain id [%v] name [%v] error [%s]", opt.ChainId, opt.ChainName, err.Error())
		}
	}()
	return nil
}
