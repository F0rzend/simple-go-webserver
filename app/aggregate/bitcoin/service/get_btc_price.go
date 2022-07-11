package service

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
)

func (bs *BitcoinServiceImpl) GetBTCPrice() entity.BTCPrice {
	return bs.bitcoinRepository.GetPrice()
}
