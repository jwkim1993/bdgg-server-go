package models

import (
	"gorm.io/gorm"
	"time"
)

type Match struct {
	gorm.Model
	GameId     uint      `json:"gameId"`
	PlayedTime time.Time `json:"playedTime"`
	Place      string    `json:"place,omitempty"`
	WinnerId   uint      `json:"winnerId,omitempty"`

	GroupID uint `json:"groupId"`
	Group   Group
}
