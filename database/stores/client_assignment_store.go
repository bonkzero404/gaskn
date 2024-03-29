package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ClientAssignment /*
type ClientAssignment struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primary_key"`
	ClientId uuid.UUID `gorm:"type:char(36):index"`
	Client   Client    `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserId   uuid.UUID `gorm:"type:char(36):index"`
	User     User      `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsActive bool
}

// BeforeCreate /*
func (*ClientAssignment) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
