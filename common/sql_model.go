package common

import "time"

type SQLModel struct {
	Id int `json:"id" gorm:"column:id;"`
	//FakeId *UID `json:"fake_id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"created_at"`
	UpdateAt  *time.Time `json:",omitempty" gorm:"update_at"`
}
