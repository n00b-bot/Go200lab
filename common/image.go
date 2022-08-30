package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty"`
	Extension string `json:"extension"`
}

func (Image) TableName() string { return "images" }

func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to Unmarsal JSONB value")
	}
	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}
	*i = img
	return nil
}

func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}

type Images []Image

func (i *Images) Scan(data interface{}) error {
	bytes, ok := data.([]byte)
	if !ok {
		return errors.New("Cannot Unmarsal Images")
	}
	var imgs Images
	if err := json.Unmarshal(bytes, &imgs); err != nil {
		return err
	}
	*i = imgs
	return nil
}

func (i *Images) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}
