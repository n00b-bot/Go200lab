package common

import "time"

const (
	RestaurantType = 1
)

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeID    UID        `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (m *SQLModel) GenUID(DbType int) {
	uid := NewUID(uint32(m.Id), DbType, 1)
	m.FakeID = uid
}
