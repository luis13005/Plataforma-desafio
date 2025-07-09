package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (base *Base) BeforeCreate(scopo *gorm.Scope) error {

	err := scopo.SetColumn("CreatedAt", time.Now())
	if err != nil {
		log.Fatalf("Erro during setting time created: %v", err)
	}

	err = scopo.SetColumn("ID", uuid.NewV4().String())
	if err != nil {
		log.Fatalf("Erro during setting time created: %v", err)
	}
	return nil
}
