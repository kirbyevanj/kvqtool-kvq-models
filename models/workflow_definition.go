package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type WorkflowDefinition struct {
	bun.BaseModel `bun:"table:workflow_definitions,alias:wd"`

	ID          uuid.UUID       `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	ProjectID   uuid.UUID       `bun:"project_id,notnull,type:uuid" json:"project_id"`
	Name        string          `bun:"name,notnull" json:"name"`
	DAGJson     json.RawMessage `bun:"dag_json,notnull,type:jsonb" json:"dag_json"`
	InputSchema json.RawMessage `bun:"input_schema,type:jsonb" json:"input_schema,omitempty"`
	CreatedAt   time.Time       `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time       `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`

	Project *Project `bun:"rel:belongs-to,join:project_id=id" json:"-"`
	Jobs    []*Job   `bun:"rel:has-many,join:id=workflow_id" json:"-"`
}
