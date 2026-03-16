package types

import (
	"encoding/json"

	"github.com/google/uuid"
)

// --- Projects ---

type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ProjectSummary struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	CreatedAt     string    `json:"created_at"`
	ResourceCount int       `json:"resource_count"`
	JobCount      int       `json:"job_count"`
}

type ProjectListResponse struct {
	Projects []ProjectSummary `json:"projects"`
}

// --- Virtual Folders ---

type CreateFolderRequest struct {
	Name     string     `json:"name"`
	ParentID *uuid.UUID `json:"parent_id"`
}

type UpdateFolderRequest struct {
	Name     string     `json:"name,omitempty"`
	ParentID *uuid.UUID `json:"parent_id,omitempty"`
}

type FolderNode struct {
	ID       uuid.UUID    `json:"id"`
	Name     string       `json:"name"`
	Path     string       `json:"path"`
	ParentID *uuid.UUID   `json:"parent_id"`
	Children []FolderNode `json:"children"`
}

type FolderTreeResponse struct {
	Folders []FolderNode `json:"folders"`
}

// --- Resources ---

type UploadURLRequest struct {
	Filename    string     `json:"filename"`
	ContentType string     `json:"content_type"`
	FolderID    *uuid.UUID `json:"folder_id"`
}

type UploadURLResponse struct {
	ResourceID       uuid.UUID `json:"resource_id"`
	UploadURL        string    `json:"upload_url"`
	S3Key            string    `json:"s3_key"`
	ExpiresInSeconds int       `json:"expires_in_seconds"`
}

type DownloadURLResponse struct {
	DownloadURL      string `json:"download_url"`
	ExpiresInSeconds int    `json:"expires_in_seconds"`
}

type UpdateResourceRequest struct {
	Name     string     `json:"name,omitempty"`
	FolderID *uuid.UUID `json:"folder_id,omitempty"`
}

// --- Workflows ---

type CreateWorkflowRequest struct {
	Name        string          `json:"name"`
	DAGJson     json.RawMessage `json:"dag_json"`
	InputSchema json.RawMessage `json:"input_schema,omitempty"`
}

type UpdateWorkflowRequest struct {
	Name        string          `json:"name,omitempty"`
	DAGJson     json.RawMessage `json:"dag_json,omitempty"`
	InputSchema json.RawMessage `json:"input_schema,omitempty"`
}

// --- Jobs ---

type CreateJobRequest struct {
	WorkflowID     uuid.UUID       `json:"workflow_id"`
	InputParams    json.RawMessage `json:"input_params,omitempty"`
	OutputFolderID *uuid.UUID      `json:"output_folder_id"`
}

type CreateJobResponse struct {
	JobID  uuid.UUID `json:"job_id"`
	Status string    `json:"status"`
}

// --- WebSocket ---

type WSProgressMessage struct {
	Type         string  `json:"type"` // "progress", "status", "error"
	ProgressPct  int32   `json:"progress_pct,omitempty"`
	CurrentFrame int64   `json:"current_frame,omitempty"`
	TotalFrames  int64   `json:"total_frames,omitempty"`
	FPS          float64 `json:"fps,omitempty"`
	Status       string  `json:"status,omitempty"`
	Message      string  `json:"message,omitempty"`
}

// --- Health ---

type HealthResponse struct {
	Status   string `json:"status"`
	Postgres string `json:"postgres"`
	Valkey   string `json:"valkey"`
}
