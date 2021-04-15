package model

type Debate struct {
	Id int32 `gorm:"type:int;not null;auto_increment" json:"id"`
	Yid int32 `gorm:"type:int;not null" json:"yid" label:"正方id"`
	Nid int32 `gorm:"type:int;not null" json:"nid" label:"反方id"`
	Title string `gorm:"type:varchar(33);not null" json:"title" label:"辩论标题"`
	NegativeContent string  `json:"negative_content" label:"正方发言"`
	PositiveContent string `json:"positive_content" label:"反方发言"`
}
