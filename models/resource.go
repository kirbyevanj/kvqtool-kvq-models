package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Resource struct {
	bun.BaseModel `bun:"table:resources,alias:r"`

	ID           uuid.UUID       `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	ProjectID    uuid.UUID       `bun:"project_id,notnull,type:uuid" json:"project_id"`
	FolderID     *uuid.UUID      `bun:"folder_id,type:uuid" json:"folder_id"`
	ResourceType string          `bun:"resource_type,notnull" json:"resource_type"`
	Name         string          `bun:"name,notnull" json:"name"`
	S3Key        string          `bun:"s3_key,notnull,unique" json:"s3_key"`
	SizeBytes    int64           `bun:"size_bytes" json:"size_bytes"`
	Metadata     json.RawMessage `bun:"metadata,type:jsonb" json:"metadata,omitempty"`
	CreatedAt    time.Time       `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time       `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`

	Project *Project       `bun:"rel:belongs-to,join:project_id=id" json:"-"`
	Folder  *VirtualFolder `bun:"rel:belongs-to,join:folder_id=id" json:"-"`
}
