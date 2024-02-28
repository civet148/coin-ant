package chains

import (
	"coin-ant/api"
	"coin-ant/dao"
	"coin-ant/models"
	"coin-ant/types"
	"encoding/hex"
	"fmt"
	"github.com/civet148/btckeys"
	"github.com/civet148/httpc"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"github.com/go-co-op/gocron"
	"net/url"
	"time"
)

const (
	PageMax                  = 100
	ChainIdBitcoin           = 0
	ChainNameBitcoin         = "Bitcoin"
	BitcoinContractAddr      = ""
	SymbolNameBitcoin        = "BTC"
	CronIntervalSyncRichList = "60s"
	BitcoinRichListUri       = "https://api.blockchair.com/bitcoin/addresses"
)

type wallet struct {
	Address string `json:"address"`
	Balance int64  `json:"balance"`
}

type BlockChairResponse struct {
	Data    []wallet `json:"data"`
	Context struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	} `json:"context"`
}

type Bitcoin struct {
	page        int
	client      *httpc.Client
	db          *sqlca.Engine
	opt         *types.ChainOption
	scheduler   *gocron.Scheduler
	richListDAO *dao.RichListDAO
}

func NewBitcoin(db *sqlca.Engine, opt *types.ChainOption) api.ChainApi {
	return &Bitcoin{
		page:        1,
		db:          db,
		opt:         opt,
		richListDAO: dao.NewRichListDAO(db),
		client: httpc.NewHttpClient(&httpc.Option{
			Timeout: 30,
		}),
		scheduler: gocron.NewScheduler(time.UTC),
	}
}

// run service loop
func (m *Bitcoin) Run() (err error) {
	log.Infof("chain name [%s] symbol [%s] running...", m.opt.ChainName, m.opt.Symbol)
	_, err = m.scheduler.Every(CronIntervalSyncRichList).Do(m.syncRichList)
	if err != nil {
		return log.Errorf(err.Error())
	}
	m.scheduler.StartAsync()

	m.runKeyCompare()
	return nil
}

func (m *Bitcoin) Close() {

}

func (m *Bitcoin) syncRichList() {
	//GET https://api.blockchair.com/bitcoin/addresses?limit=100&offset=200&page=2
	if m.page > PageMax {
		m.page = 1
	}
	limit := 100
	offset := m.page * limit
	log.Infof("rich list sync page %v limit %v offset %v...", m.page, limit, offset)
	r, err := m.client.Get(BitcoinRichListUri, url.Values{
		"limit":  []string{fmt.Sprintf("%v", limit)},
		"offset": []string{fmt.Sprintf("%v", offset)},
		"page":   []string{fmt.Sprintf("%v", m.page)},
	})
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	if r.StatusCode != 200 {
		log.Errorf("GET request return http status code [%v] message [%s]", r.StatusCode, r.Body)
		return
	}
	var response BlockChairResponse
	if err = r.Unmarshal(&response); err != nil {
		log.Errorf("unmarshal error: %v", err.Error())
		return
	}
	if response.Context.Code != 200 {
		log.Errorf("response error: %v", response.Context.Error)
		return
	}
	for _, w := range response.Data {
		_, err = m.richListDAO.Insert(&models.RichListDO{
			ChainId:   ChainIdBitcoin,
			ChainName: ChainNameBitcoin,
			Symbol:    SymbolNameBitcoin,
			Address:   w.Address,
			Balance:   sqlca.NewDecimal(w.Balance),
		})
		if err != nil {
			log.Errorf(err.Error())
			return
		}
	}
	m.page++
}

func (m *Bitcoin) runKeyCompare() {
	for {
		key, err := btckeys.GenBitcoinKey("", "", 128, 0)
		if err != nil {
			log.Errorf(err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		btcAddr := key.Address()
		bech32Addr := key.Bech32()
		dos, err := m.richListDAO.QueryByAddresses(ChainIdBitcoin, SymbolNameBitcoin, BitcoinContractAddr, btcAddr, bech32Addr)
		if err != nil {
			log.Errorf(err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		log.Debugf("gen key address [%v] bech32 [%s] to compare...", btcAddr, bech32Addr)
		for _, do := range dos {
			do.IsOk = true
			do.PrivateKey = hex.EncodeToString(key.PrivateKeyBytes())
			_, err = m.richListDAO.Update(do, models.RICH_LIST_COLUMN_IS_OK, models.RICH_LIST_COLUMN_PRIVATE_KEY)
			if err != nil {
				log.Errorf(err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
		}
		time.Sleep(50 * time.Nanosecond)
	}
}
