package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"todo-list/common"
	"todo-list/module/user/model"
)

type ItemStatus int

const (
	ItemStatusDOing ItemStatus = iota // 0
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatus = [3]string{"Doing", "Done", "Deleted"}

func (item *ItemStatus) StatusString() string {
	return allItemStatus[*item]
}

func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatus {
		if allItemStatus[i] == s {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New("invalid status string")
}

// lay tu db len map vao struct
// override function scan data  tu db
func (item *ItemStatus) Scan(value interface{}) error { // impliment khi 2 du lieu duoi db va struct khac nhau (duoi db len)
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	v, err := parseStr2ItemStatus(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	*item = v

	return nil
}

// override function map du lieu tu struct sang 1 dang khac
// thay doi gia tri len json
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.StatusString())), nil
}

/////////// chieu day du lieu len sever

// override function Value de chuyen du lieu tu struct xuong db
// int -> string de luu vao db
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.StatusString(), nil
}

// override function UnmarshalJSON map du lieu tu json -> struct sang 1 dang khac
// thay doi gia tri tu json len struct (Doing -> 0)
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "") // xoa dau "

	itemValue, err := parseStr2ItemStatus(str)
	if err != nil {
		return err
	}

	*item = itemValue // doi data
	return nil
}

// /// err
var (
	ErrTitleCannotBeEmpty = errors.New("Title cannot be empty")
	ErrItemIsDeleted      = errors.New("Item is Deleted")
)

const (
	EntityName = "Item"
)

// struct
type TodoItem struct {
	common.SQLModel
	Title       string        `json:"title" gorm:"column:title;"`
	Description string        `json:"description" gorm:"column:description;"`
	Status      *ItemStatus   `json:"status" gorm:"column:status;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
	LikedCount  int           `json:"liked_count" gorm: "-"`
	// image
	IntExample    int               `json:"int_example" gorm:"column:int_example;"`
	DoubleExample float64           `json:"double_example" gorm:"column:double_example;"`
	UserId        int               `json:"user_id" gorm:"column:user_id;"`
	Owner         *model.SimpleUser `json:"owner" gorm:"foreignKey:UserId;"` // foreignkey de ket noi
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreate struct {
	Id          int           `json:"-" gorm:"column:id;"`
	Title       string        `json:"title" gorm:"column:title;"`
	Description string        `json:"description" gorm:"column:description;"`
	Image       *common.Image `json:"image" gorm:"column:image;"`
	UserId      int           `json:"-" gorm:"column:user_id;"`
}

// validate
func (i *TodoItemCreate) Validate() error {
	i.Title = strings.TrimSpace(i.Title)
	if i.Title == "" {
		return ErrTitleCannotBeEmpty
	}
	return nil
}

func (TodoItemCreate) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	// Id          int    `json:"-" gorm:"column:id;"`
	Title       *string     `json:"title" gorm:"column:title;"`
	Description *string     `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
