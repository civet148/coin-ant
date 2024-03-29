// Code generated by db2go. DO NOT EDIT.
// https://github.com/civet148/sqlca

package models

import "github.com/civet148/sqlca/v2"

const TableNameTokenList = "token_list" //

const (
	TOKEN_LIST_COLUMN_ID            = "id"
	TOKEN_LIST_COLUMN_CHAIN_ID      = "chain_id"
	TOKEN_LIST_COLUMN_CHAIN_NAME    = "chain_name"
	TOKEN_LIST_COLUMN_SYMBOL        = "symbol"
	TOKEN_LIST_COLUMN_DECIMALS      = "decimals"
	TOKEN_LIST_COLUMN_CONTRACT_ADDR = "contract_addr"
	TOKEN_LIST_COLUMN_MIN_BALANCE   = "min_balance"
	TOKEN_LIST_COLUMN_DELETED       = "deleted"
	TOKEN_LIST_COLUMN_CREATED_TIME  = "created_time"
	TOKEN_LIST_COLUMN_UPDATED_TIME  = "updated_time"
)

type TokenListDO struct {
	Id           int32         `json:"id" db:"id" `                                     //auto-incr id
	ChainId      int64         `json:"chain_id" db:"chain_id" `                         //block chain id [0=Bitcoin 1=Ethereum]
	ChainName    string        `json:"chain_name" db:"chain_name" `                     //block chain name [Bitcoin/Ethereum]
	Symbol       string        `json:"symbol" db:"symbol" `                             //token symbol
	Decimals     int32         `json:"decimals" db:"decimals" `                         //token decimals
	ContractAddr string        `json:"contract_addr" db:"contract_addr" `               //contract address (empty means native token)
	MinBalance   sqlca.Decimal `json:"min_balance" db:"min_balance" `                   //account's minimal balance  to attack
	Deleted      bool          `json:"deleted" db:"deleted" `                           //is deleted (0=no 1=yes)
	CreatedTime  string        `json:"created_time" db:"created_time" sqlca:"readonly"` //create time
	UpdatedTime  string        `json:"updated_time" db:"updated_time" sqlca:"readonly"` //updte time
}

func (do *TokenListDO) GetId() int32                  { return do.Id }
func (do *TokenListDO) SetId(v int32)                 { do.Id = v }
func (do *TokenListDO) GetChainId() int64             { return do.ChainId }
func (do *TokenListDO) SetChainId(v int64)            { do.ChainId = v }
func (do *TokenListDO) GetChainName() string          { return do.ChainName }
func (do *TokenListDO) SetChainName(v string)         { do.ChainName = v }
func (do *TokenListDO) GetSymbol() string             { return do.Symbol }
func (do *TokenListDO) SetSymbol(v string)            { do.Symbol = v }
func (do *TokenListDO) GetDecimals() int32            { return do.Decimals }
func (do *TokenListDO) SetDecimals(v int32)           { do.Decimals = v }
func (do *TokenListDO) GetContractAddr() string       { return do.ContractAddr }
func (do *TokenListDO) SetContractAddr(v string)      { do.ContractAddr = v }
func (do *TokenListDO) GetMinBalance() sqlca.Decimal  { return do.MinBalance }
func (do *TokenListDO) SetMinBalance(v sqlca.Decimal) { do.MinBalance = v }
func (do *TokenListDO) GetDeleted() bool              { return do.Deleted }
func (do *TokenListDO) SetDeleted(v bool)             { do.Deleted = v }
func (do *TokenListDO) GetCreatedTime() string        { return do.CreatedTime }
func (do *TokenListDO) SetCreatedTime(v string)       { do.CreatedTime = v }
func (do *TokenListDO) GetUpdatedTime() string        { return do.UpdatedTime }
func (do *TokenListDO) SetUpdatedTime(v string)       { do.UpdatedTime = v }

/*
CREATE TABLE `token_list` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'auto-incr id',
  `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'block chain id [0=Bitcoin 1=Ethereum]',
  `chain_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'block chain name [Bitcoin/Ethereum]',
  `symbol` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'token symbol',
  `decimals` int NOT NULL DEFAULT '0' COMMENT 'token decimals',
  `contract_addr` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'contract address (empty means native token)',
  `min_balance` decimal(50,5) NOT NULL DEFAULT '0.00000' COMMENT 'account''s minimal balance  to attack',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'is deleted (0=no 1=yes)',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updte time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
*/
