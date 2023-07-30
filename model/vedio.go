package model

import (
	"errors"
	"log"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Vedio struct {
	gorm.Model
	UserId   uint   `gorm:"column:user_id"`
	PlayUrl  string `gorm:"column:play_url"`
	CoverUrl string `gorm:"column:cover_url"`
	Title    string `gorm:"column:title"`
}

func (Vedio) TableName() string {
	return "vedio"
}

type VedioDao struct {
}

var vedioDao *VedioDao
var vedioOnce sync.Once

func NewVedioDaoInstance() *VedioDao {
	vedioOnce.Do(func() {
		vedioDao = &VedioDao{}
	})
	return vedioDao
}

func (*VedioDao) QueryVedioListByTime(limit int, lastest time.Time) ([]*Vedio, error) {
	var vedioList []*Vedio = make([]*Vedio, 0, 30)
	err := db.Model(&Vedio{}).Where("created_at < ?", lastest).Limit(limit).Order("created_at desc").Find(&vedioList).Error
	log.Println("Query vedio nums:", len(vedioList))
	if err != nil {
		return nil, err
	}
	return vedioList, nil
}

func (*VedioDao) AddVedio(vedio *Vedio) (*Vedio, error) {
	if vedio == nil {
		log.Println("Vedio is nil")
		return nil, errors.New("Vedio is nil")
	}
	err := db.Create(vedio).Error
	if err != nil {
		log.Println("Add Vedio failed:", err.Error())
		return nil, err
	}
	return vedio, nil
}
