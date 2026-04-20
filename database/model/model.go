package model

import "encoding/json"

type Setting struct {
	Id    uint   `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Key   string `json:"key" form:"key" gorm:"unique;type:varchar(64);not null;"`
	Value string `json:"value" form:"value"`
}

type Tls struct {
	Id     uint            `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Name   string          `json:"name" form:"name" gorm:"type:varchar(64);not null;"`
	Server json.RawMessage `json:"server" form:"server" gorm:"type:JSON;"`
	Client json.RawMessage `json:"client" form:"client" gorm:"type:JSON;"`
}

type User struct {
	Id         uint   `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Username   string `json:"username" form:"username" gorm:"unique;type:varchar(32);not null;"`
	Password   string `json:"password" form:"password" gorm:"type:varchar(64);not null;"`
	LastLogins string `json:"lastLogin" gorm:"type:varchar(128);"`
}

type Client struct {
	Id       uint            `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Enable   bool            `json:"enable" form:"enable"`
	Name     string          `json:"name" form:"name" gorm:"unique;type:varchar(64);not null;"`
	Config   json.RawMessage `json:"config,omitempty" form:"config" gorm:"type:JSON;"`
	Inbounds json.RawMessage `json:"inbounds" form:"inbounds" gorm:"type:JSON;"`
	Links    json.RawMessage `json:"links,omitempty" form:"links" gorm:"type:JSON;"`
	Volume   int64           `json:"volume" form:"volume"`
	Expiry   int64           `json:"expiry" form:"expiry"`
	Down     int64           `json:"down" form:"down"`
	Up       int64           `json:"up" form:"up"`
	Desc     string          `json:"desc" form:"desc" gorm:"type:varchar(64);"`
	Group    string          `json:"group" form:"group" gorm:"type:varchar(64);"`

	// Delay start and periodic reset
	DelayStart bool  `json:"delayStart" form:"delayStart" gorm:"default:false;not null"`
	AutoReset  bool  `json:"autoReset" form:"autoReset" gorm:"default:false;not null"`
	ResetDays  int   `json:"resetDays" form:"resetDays" gorm:"default:0;not null"`
	NextReset  int64 `json:"nextReset" form:"nextReset" gorm:"default:0;not null"`
	TotalUp    int64 `json:"totalUp" form:"totalUp" gorm:"default:0;not null"`
	TotalDown  int64 `json:"totalDown" form:"totalDown" gorm:"default:0;not null"`
}

type Stats struct {
	Id        uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	DateTime  int64  `json:"dateTime"`
	Resource  string `json:"resource" gorm:"type:varchar(64);"`
	Tag       string `json:"tag" gorm:"type:varchar(64);"`
	Direction bool   `json:"direction"`
	Traffic   int64  `json:"traffic"`
}

type Changes struct {
	Id       uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	DateTime int64           `json:"dateTime"`
	Actor    string          `json:"actor" gorm:"type:varchar(32);not null;"`
	Key      string          `json:"key" gorm:"type:varchar(64);"`
	Action   string          `json:"action" gorm:"type:varchar(64);"`
	Obj      json.RawMessage `json:"obj" gorm:"type:JSON;"`
}

type Tokens struct {
	Id     uint   `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Desc   string `json:"desc" form:"desc" gorm:"type:varchar(64);not null;"`
	Token  string `json:"token" form:"token" gorm:"type:varchar(128);not null;"`
	Expiry int64  `json:"expiry" form:"expiry"`
	UserId uint   `json:"userId" form:"userId"`
	User   *User  `json:"user" gorm:"foreignKey:UserId;references:Id"`
}
