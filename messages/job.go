package messages

import (
	"encoding/json"

	"github.com/google/uuid"
)

// JobMessage is pushed to the Valkey job queue for worker consumption.
type JobMessage struct {
	JobID      uuid.UUID       `json:"job_id"`
	ProjectID  uuid.UUID       `json:"project_id"`
	WorkflowID uuid.UUID       `json:"workflow_id"`
	DAGJson    json.RawMessage `json:"dag_json"`
	Params     json.RawMessage `json:"input_params"`
}

// JobProgress is published to the Valkey pub/sub channel for real-time updates.
type JobProgress struct {
	JobID        uuid.UUID `json:"job_id"`
	Status       string    `json:"status"`
	ProgressPct  int32     `json:"progress_pct"`
	CurrentFrame int64     `json:"current_frame"`
	TotalFrames  int64     `json:"total_frames"`
	FPS          float64   `json:"fps"`
	Message      string    `json:"message,omitempty"`
}

// JobResult is published when a job completes or fails.
type JobResult struct {
	JobID       uuid.UUID       `json:"job_id"`
	Status      string          `json:"status"`
	OutputS3Key string          `json:"output_s3_key,omitempty"`
	ReportS3Key string          `json:"report_s3_key,omitempty"`
	Error       string          `json:"error,omitempty"`
	Result      json.RawMessage `json:"result,omitempty"`
}

const (
	JobQueueKey       = "kvq:jobs:pending"
	JobProgressPrefix = "job:%s:progress"
)
