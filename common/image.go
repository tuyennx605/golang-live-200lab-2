package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"` // aws, google...
	Extension string `json:"extension,omitempty" gorm:"-"`
}

// ten table
func (Image) TableName() string { return "images" }

func (j *Image) Fulfill(doumain string) {
	j.Url = fmt.Sprintf("%s/%s", doumain, j.Url)
}

// khi doc tu db se tu dong vao ham nay de chuyen doi data
// chuyển đổi từ database sang struct
func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte) // lấy mảng byte từ db lên
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil { // chuyển đổi từ mảng byte sang struct
		return err
	}

	*j = img // gan', sua  value cho con tro
	return nil
}

// Value return json value, impement driver.Valuer interface
// luu vao db
// chuyển đổi từ struct vào db
func (j *Image) Value() (driver.Value, error) { // chi dan cho sql biet lam gi voi struct image nay de luu vao db
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j) // bien object thanh json string dang mang byte
}

/////////////////////// images (nhiều ảnh)

type Images []Image

// khi doc tu db se tu dong vao ham nay de chuyen doi data
func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img []Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img // gan', sua  value cho con tro
	return nil
}

// Value return json value, impement driver.Valuer interface
// luu vao db
func (j *Images) Value() (driver.Value, error) { // chi dan cho sql biet lam gi voi struct image nay de luu vao db
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j) // bien object thanh json string dang mang byte
}
