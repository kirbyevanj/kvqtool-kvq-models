package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Project struct {
	bun.BaseModel `bun:"table:projects,alias:p"`

	ID          uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID `bun:"user_id,notnull,type:uuid" json:"user_id"`
	Name        string    `bun:"name,notnull" json:"name"`
	Description string    `bun:"description" json:"description"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`

	User      *User                 `bun:"rel:belongs-to,join:user_id=id" json:"-"`
	Resources []*Resource           `bun:"rel:has-many,join:id=project_id" json:"-"`
	Folders   []*VirtualFolder      `bun:"rel:has-many,join:id=project_id" json:"-"`
	Workflows []*WorkflowDefinition `bun:"rel:has-many,join:id=project_id" json:"-"`
	Jobs      []*Job                `bun:"rel:has-many,join:id=project_id" json:"-"`
}
