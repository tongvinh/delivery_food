package restaurantmodel

import (
	"errors"
	"myapp/common"
	"strings"
)

type RestaurantType string

const TypeNormal RestaurantType = "normal"
const TypePremium RestaurantType = "premium"
const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	Name       string             `json:"name" gorm:"column:name;"`
	Addr       string             `json:"address" gorm:"column:addr;"`
	Type       RestaurantType     `json:"type" gorm:"column:type;"`
	Logo       *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover      *common.Images     `json:"cover" gorm:"column:cover;"`
	UserId     int                `json:"-" gorm:"column:user_id;"`
	User       *common.SimpleUser `json:"user" gorm:"preload:false;"`
	LikedCount int                `json:"liked_count" gorm:"column:liked_count;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

type RestaurantCreate struct {
	common.SQLModel
	Name   string         `json:"name" gorm:"column:name;"`
	Addr   string         `json:"address" gorm:"column:addr;"`
	UserId int            `json:"-" gorm:"column:user_id;"`
	Logo   *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover  *common.Images `json:"cover" gorm:"column:cover;"`
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"addr" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
