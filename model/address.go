package model

import (
	"wallet-srv/lib/db"

	"github.com/golang/glog"
)

type IfAddress interface {
	AddData(table string, addr *CoinAddress)
	GetData(table string, addr *CoinAddress)
	FindData(table string, limit uint32) *[]CoinAddress
	UpdateData(table string, Id int64) bool
	CountData(table string, isUsed uint8) int64
}

type CoinAddress struct {
	ID      int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Address string
	PrivKey string
	IsUsed  int
	Multi   int
}

func (a *CoinAddress) AddData(table string, addr *CoinAddress) {

	a.ID = addr.ID
	a.Address = addr.Address
	a.PrivKey = addr.PrivKey
	a.IsUsed = 0
	a.Multi = addr.Multi

	db.DB().Table(table).Create(&a)
}

func (a *CoinAddress) GetData(table string, addr *CoinAddress) {
	db.DB().Table(table).Where("address = ?", addr.Address).Take(&a)
	if a.Address == "" || a.PrivKey == "" {
		return
	}

	addr.PrivKey = a.PrivKey
	addr.ID = a.ID
	addr.IsUsed = a.IsUsed
	addr.Multi = a.Multi

	//decrypt privkey
	err := DBDecrypt(addr)

	if err != nil {
		glog.Fatal(err)
		return
	}
}

func (a *CoinAddress) FindData(table string, limit uint32) *[]CoinAddress {
	var addr []CoinAddress
	ret := db.DB().Table(table).Where("is_used = ?", 0).Limit(limit).Find(&addr)
	if ret.Error != nil {
		return nil
	}
	return &addr
}

func (a *CoinAddress) UpdateData(table string, Id int64) bool {
	ret := db.DB().Table(table).Where("Id = ?", Id).Update("is_used", 1)
	if ret.Error != nil {
		glog.Info("UpdateData ", table, " ID=", Id, " error: ", ret.Error)
		return false
	}
	return true
}

func (a *CoinAddress) CountData(table string, isUsed uint8) int64 {
	var count int64
	count = 0
	db.DB().Table(table).Where("is_used = ?", isUsed).Count(&count)

	return count
}
