package postgres

import (
	"time"

	"github.com/google/uuid"
)

type DbUser struct {
	tableName   struct{}  `pg:"users"`
	ID          uuid.UUID `pg:"type:uuid"`
	Username    string
	Password    string
	Email       string
	DateOfBirth time.Time
	DateCreated time.Time `pg:"default:now()"`
}
