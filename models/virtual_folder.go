package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VirtualFolder struct {
	bun.BaseModel `bun:"table:virtual_folders,alias:vf"`

	ID        uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	ProjectID uuid.UUID  `bun:"project_id,notnull,type:uuid" json:"project_id"`
	ParentID  *uuid.UUID `bun:"parent_id,type:uuid" json:"parent_id"`
	Name      string     `bun:"name,notnull" json:"name"`
	Path      string     `bun:"path,notnull" json:"path"`
	CreatedAt time.Time  `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`

	Project   *Project         `bun:"rel:belongs-to,join:project_id=id" json:"-"`
	Parent    *VirtualFolder   `bun:"rel:belongs-to,join:parent_id=id" json:"-"`
	Children  []*VirtualFolder `bun:"rel:has-many,join:id=parent_id" json:"children,omitempty"`
	Resources []*Resource      `bun:"rel:has-many,join:id=folder_id" json:"-"`
}
