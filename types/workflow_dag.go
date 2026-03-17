package types

import "encoding/json"

// WorkflowDAG is the JSON contract between frontend (Drawflow) and backend (Temporal interpreter).
type WorkflowDAG struct {
	Version string              `json:"version"`
	Name    string              `json:"name"`
	Nodes   map[string]*DAGNode `json:"nodes"`
}

// DAGNode represents a single step in the workflow graph.
type DAGNode struct {
	ID      string            `json:"id"`
	Type    string            `json:"type"`
	Params  map[string]string `json:"params"`
	Inputs  []string          `json:"inputs"`
	Outputs []string          `json:"outputs"`
}

// Predefined activity type constants matching registered Temporal activities.
const (
	ActivityResDownload    = "ResourceDownload"
	ActivityResUpload      = "ResourceUpload"
	ActivityGstEncode      = "GStreamerEncode"
	ActivityGstMetrics     = "GStreamerMetrics"
	ActivitySplitVideo     = "SplitVideo"
	ActivityConcatVideo    = "ConcatVideo"
	ActivityGenerateReport = "GenerateReport"
	ActivityFMP4Repackage  = "FragmentedMP4Repackage"
)

// ActivityInput is the standard input passed to every Temporal activity.
type ActivityInput struct {
	NodeID    string            `json:"node_id"`
	NodeType  string            `json:"node_type"`
	Params    map[string]string `json:"params"`
	ProjectID string            `json:"project_id"`
	// Upstream results keyed by source node ID
	UpstreamResults map[string]json.RawMessage `json:"upstream_results,omitempty"`
}

// ActivityOutput is the standard output returned by every Temporal activity.
type ActivityOutput struct {
	NodeID  string          `json:"node_id"`
	Success bool            `json:"success"`
	S3Key   string          `json:"s3_key,omitempty"`
	Error   string          `json:"error,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

// RunWorkflowRequest is the API request to start a workflow execution.
type RunWorkflowRequest struct {
	WorkflowID string          `json:"workflow_id"`
	Params     json.RawMessage `json:"params,omitempty"`
}

// RunWorkflowResponse is returned when a workflow is started.
type RunWorkflowResponse struct {
	RunID      string `json:"run_id"`
	WorkflowID string `json:"workflow_id"`
	Status     string `json:"status"`
}

// WorkflowStatusResponse is returned when querying workflow status.
type WorkflowStatusResponse struct {
	RunID        string `json:"run_id"`
	Status       string `json:"status"`
	CurrentNode  string `json:"current_node,omitempty"`
	ProgressPct  int    `json:"progress_pct"`
	Error        string `json:"error,omitempty"`
}

const (
	TemporalTaskQueue = "kvqtool-workers"
	TemporalNamespace = "default"
)
