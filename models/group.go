package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupImage string `json:"groupImage,omitempty"`
	GroupPlace string `json:"groupPlace,omitempty"`
}

type GroupMember struct {
	gorm.Model
	MemberID uint `json:"memberId"`
	Member   Member
	GroupID  uint `json:"groupId"`
	Group    Group
}
