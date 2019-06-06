package model

import (
	"time"
)
type ActionType string

const (
	ACTIVE ActionType = "Active"
	DELETE = "Delete"
	CREATED = "Created"
)
type AdminActivityEntity struct {
	tableName struct{} `sql:"admin_activity"`
	
	Id			  string         `json:"id"`
	Action 	      ActionType     `json:"action"`
	Date	      time.Time      `json:"date"`
	Description   string		 `hson:"description"`
}
