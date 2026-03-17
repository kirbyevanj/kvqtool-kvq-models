package messages

// Legacy job message types retained for compatibility.
// New workflows use Temporal activities via types.WorkflowDAG.

import "github.com/google/uuid"

type JobProgress struct {
	JobID        uuid.UUID `json:"job_id"`
	Status       string    `json:"status"`
	ProgressPct  int32     `json:"progress_pct"`
	CurrentFrame int64     `json:"current_frame"`
	TotalFrames  int64     `json:"total_frames"`
	FPS          float64   `json:"fps"`
	Message      string    `json:"message,omitempty"`
}
