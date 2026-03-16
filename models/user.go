package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID          uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Email       string    `bun:"email,notnull,unique" json:"email"`
	DisplayName string    `bun:"display_name,notnull" json:"display_name"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`

	Projects []*Project `bun:"rel:has-many,join:id=user_id" json:"-"`
}
