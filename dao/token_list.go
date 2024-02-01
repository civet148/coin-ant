package dao

import (
	"coin-ant/models"
	"fmt"
	"github.com/civet148/sqlca/v2"
)

type TokenListDAO struct {
	db *sqlca.Engine
}

func NewTokenListDAO(db *sqlca.Engine) *TokenListDAO {
	return &TokenListDAO{
		db: db,
	}
}

// insert into table by data model
func (dao *TokenListDAO) Insert(do *models.TokenListDO) (lastInsertId int64, err error) {
	return dao.db.Model(&do).Table(models.TableNameTokenList).Insert()
}

// insert if not exist or update columns on duplicate key...
func (dao *TokenListDAO) Upsert(do *models.TokenListDO, columns ...string) (lastInsertId int64, err error) {
	return dao.db.Model(&do).Table(models.TableNameTokenList).Select(columns...).Upsert()
}

// update table set columns where id=xxx
func (dao *TokenListDAO) Update(do *models.TokenListDO, columns ...string) (rows int64, err error) {
	return dao.db.Model(&do).Table(models.TableNameTokenList).Select(columns...).Update()
}

// query records by id
func (dao *TokenListDAO) QueryById(id interface{}, columns ...string) (do *models.TokenListDO, err error) {
	if _, err = dao.db.Model(&do).Table(models.TableNameTokenList).Id(id).Select(columns...).Query(); err != nil {
		return nil, err
	}
	return
}

// query records by conditions
func (dao *TokenListDAO) QueryByCondition(conditions map[string]interface{}, columns ...string) (dos []*models.TokenListDO, err error) {
	if len(conditions) == 0 {
		return nil, fmt.Errorf("condition must not be empty")
	}
	e := dao.db.Model(&dos).Table(models.TableNameTokenList).Select(columns...)
	for k, v := range conditions {
		e.Eq(k, v)
	}
	if _, err = e.Query(); err != nil {
		return nil, err
	}
	return
}

// query records by id
func (dao *TokenListDAO) QueryAll(columns ...string) (dos []*models.TokenListDO, err error) {
	if _, err = dao.db.Model(&dos).Table(models.TableNameTokenList).Select(columns...).Query(); err != nil {
		return nil, err
	}
	return
}
