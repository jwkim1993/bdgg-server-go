package dto

import (
	"github.com/jwkim1993/bdgg-server/models"
)

type Member struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
}

type MemberList struct {
	Items []Member `json:"items"`
}

func (m *Member) SetMember(member models.Member) {
	m.ID = member.ID
	m.Name = member.Name
	m.ImageUrl = member.ImageUrl
}

func (m *MemberList) SetMembers(members []models.Member) {
	m.Items = make([]Member, len(members))

	for i := range members {
		m.Items[i].SetMember(members[i])
	}
}
