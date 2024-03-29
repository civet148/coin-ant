package dao

import (
	"coin-ant/models"
	"fmt"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

type RichListDAO struct {
	db *sqlca.Engine
}

func NewRichListDAO(db *sqlca.Engine) *RichListDAO {
	return &RichListDAO{
		db: db,
	}
}

// insert into table by data model
func (dao *RichListDAO) Insert(do *models.RichListDO) (lastInsertId int64, err error) {
	_, err = dao.db.Model(&do).
		Table(models.TableNameRichList).
		Select(models.RICH_LIST_COLUMN_ID).
		Eq(models.RICH_LIST_COLUMN_CHAIN_ID, do.ChainId).
		Eq(models.RICH_LIST_COLUMN_SYMBOL, do.Symbol).
		Eq(models.RICH_LIST_COLUMN_ADDRESS, do.Address).
		Query()
	if err != nil {
		return 0, log.Errorf(err.Error())
	}
	if do.Id != 0 {
		return 0, nil
	}
	return dao.db.Model(&do).Table(models.TableNameRichList).Insert()
}

// insert if not exist or update columns on duplicate key...
func (dao *RichListDAO) Upsert(do *models.RichListDO, columns ...string) (lastInsertId int64, err error) {
	return dao.db.Model(&do).Table(models.TableNameRichList).Select(columns...).Upsert()
}

// update table set columns where id=xxx
func (dao *RichListDAO) Update(do *models.RichListDO, columns ...string) (rows int64, err error) {
	return dao.db.Model(&do).Table(models.TableNameRichList).Select(columns...).Update()
}

// query records by id
func (dao *RichListDAO) QueryById(id interface{}, columns ...string) (do *models.RichListDO, err error) {
	if _, err = dao.db.Model(&do).Table(models.TableNameRichList).Id(id).Select(columns...).Query(); err != nil {
		return nil, err
	}
	return
}

// query records by conditions
func (dao *RichListDAO) QueryByCondition(conditions map[string]interface{}, columns ...string) (dos []*models.RichListDO, err error) {
	if len(conditions) == 0 {
		return nil, fmt.Errorf("condition must not be empty")
	}
	e := dao.db.Model(&dos).Table(models.TableNameRichList).Select(columns...)
	for k, v := range conditions {
		e.Eq(k, v)
	}
	if _, err = e.Query(); err != nil {
		return nil, err
	}
	return
}

func (dao *RichListDAO) QueryByAddresses(chainId int, symbol, contract string, addresses ...string) (dos []*models.RichListDO, err error) {
	for _, addr := range addresses {
		e := dao.db.Model(&dos).
			Table(models.TableNameRichList).
			Eq(models.RICH_LIST_COLUMN_CHAIN_ID, chainId).
			Eq(models.RICH_LIST_COLUMN_ADDRESS, addr).
			Eq(models.RICH_LIST_COLUMN_CONTRACT_ADDR, contract).
			Eq(models.RICH_LIST_COLUMN_SYMBOL, symbol)
		if _, err = e.Query(); err != nil {
			return nil, err
		}
		if len(dos) != 0 {
			return dos, nil
		}
	}
	return
}
