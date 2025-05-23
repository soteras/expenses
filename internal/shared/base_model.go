package shared

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" valid:"-"`
	CreatedAt time.Time      `valid:"-"`
	UpdatedAt time.Time      `valid:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" valid:"-"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}
