package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name              string `json:"name"`
	ImageUrl          string `json:"imageUrl,omitempty"`
	Provider          string `json:"provider,omitempty"`
	SnsId             string `json:"snsId,omitempty"`
	NotificationToken string `json:"token,omitempty"`
}

type MemberMatch struct {
	gorm.Model
	MemberID uint `json:"groupId"`
	Member   Member
	MatchID  uint `json:"matchId"`
	Match    Match
	Score    int64 `json:"score,omitempty"`
}
