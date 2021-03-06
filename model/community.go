package model

import "time"

const (
	COMMUNITY_CATE_COM = 0x01
)

type Community struct {
	Id       int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Name     string    `xorm:"varchar(30)" form:"name" json:"name"`
	OwnerId  int64     `xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	Icon     string    `xorm:"varchar(250)" form:"icon" json:"icon"`
	Cate     int       `xorm:"int(11)" form:"cate" json:"cate"`
	Memo     string    `xorm:"varchar(120)" form:"memo" json:"memo"`
	CreateAt time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
