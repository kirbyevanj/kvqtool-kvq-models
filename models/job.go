package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Job struct {
	bun.BaseModel `bun:"table:jobs,alias:j"`

	ID           uuid.UUID       `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	ProjectID    uuid.UUID       `bun:"project_id,notnull,type:uuid" json:"project_id"`
	WorkflowID   uuid.UUID       `bun:"workflow_id,notnull,type:uuid" json:"workflow_id"`
	Status       string          `bun:"status,notnull,default:'pending'" json:"status"`
	InputParams  json.RawMessage `bun:"input_params,type:jsonb" json:"input_params,omitempty"`
	Result       json.RawMessage `bun:"result,type:jsonb" json:"result,omitempty"`
	ProgressPct  int32           `bun:"progress_pct,notnull,default:0" json:"progress_pct"`
	CurrentFrame int64           `bun:"current_frame,notnull,default:0" json:"current_frame"`
	TotalFrames  int64           `bun:"total_frames,notnull,default:0" json:"total_frames"`
	StartedAt    *time.Time      `bun:"started_at" json:"started_at,omitempty"`
	CompletedAt  *time.Time      `bun:"completed_at" json:"completed_at,omitempty"`
	CreatedAt    time.Time       `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time       `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`

	Project  *Project            `bun:"rel:belongs-to,join:project_id=id" json:"-"`
	Workflow *WorkflowDefinition `bun:"rel:belongs-to,join:workflow_id=id" json:"-"`
}

const (
	JobStatusPending   = "pending"
	JobStatusQueued    = "queued"
	JobStatusRunning   = "running"
	JobStatusCompleted = "completed"
	JobStatusFailed    = "failed"
	JobStatusCancelled = "cancelled"
)
