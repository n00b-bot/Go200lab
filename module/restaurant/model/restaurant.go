package restaurantmodel

import "food/common"

const Entity string = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Addr            string        `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image `json:"logo" gorm:"column:logo"`
	Cover			*common.Images `json:"cover" gorm:"column:cover"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Addr            string        `json:"addr" gorm:"column:addr;"`
	Logo            *common.Image `json:"logo" gorm:"column:logo"`
	Cover			*common.Images `json:"cover" gorm:"column:cover"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }
