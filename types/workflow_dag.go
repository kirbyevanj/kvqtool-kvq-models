package types

import "encoding/json"

// WorkflowDAG is the JSON contract between frontend (Drawflow) and backend (Temporal interpreter).
type WorkflowDAG struct {
	Version       string                   `json:"version"`
	Name          string                   `json:"name"`
	Nodes         map[string]*DAGNode      `json:"nodes"`
	GlobalInputs  map[string]*GlobalInput  `json:"global_inputs,omitempty"`
	SessionGroups map[string]*SessionGroup `json:"session_groups,omitempty"`
}

// GlobalInput defines a typed workflow-level parameter that can be mapped into individual node params.
type GlobalInput struct {
	Name    string `json:"name"`
	Type    string `json:"type"`              // "int", "float", "bool", "string"
	Default string `json:"default,omitempty"`
}

// SessionGroup groups nodes that must execute on the same physical worker (Temporal session).
type SessionGroup struct {
	ID    string   `json:"id"`
	Label string   `json:"label,omitempty"`
	Nodes []string `json:"nodes"`
}

// PortType constants for node connection ports.
const (
	PortTypeVideo    = "video"     // raw/encoded video file
	PortTypeSceneCut = "scenecut"  // SceneCutFile JSON
	PortTypeMetrics  = "metrics"   // MetricReports JSON
	PortTypeReport   = "report"    // generated report artifact
	PortTypeAny      = "any"       // untyped / passthrough
)

// DAGNode represents a single step in the workflow graph.
type DAGNode struct {
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	Params       map[string]string `json:"params"`
	Inputs       []string          `json:"inputs"`
	Outputs      []string          `json:"outputs"`
	// InputTypes / OutputTypes map port-index ("0","1",...) → PortType constant.
	InputTypes   map[string]string `json:"input_types,omitempty"`
	OutputTypes  map[string]string `json:"output_types,omitempty"`
	InputMap     map[string]string `json:"input_map,omitempty"`
	SessionGroup string            `json:"session_group,omitempty"`
	WorkflowRef  string            `json:"workflow_ref,omitempty"`
}

// NodePortSpec describes the expected input/output port types for a node type.
type NodePortSpec struct {
	InputTypes  []string // per-port type, indexed 0..n
	OutputTypes []string
}

// NodePortSpecs is the canonical per-node-type port definition used by the interpreter and frontend.
var NodePortSpecs = map[string]NodePortSpec{
	ActivityResDownload:         {InputTypes: []string{}, OutputTypes: []string{PortTypeVideo}},
	ActivityResUpload:           {InputTypes: []string{PortTypeVideo}, OutputTypes: []string{}},
	ActivityX264Transcode:       {InputTypes: []string{PortTypeVideo}, OutputTypes: []string{PortTypeVideo}},
	ActivityFileMetricAnalysis:  {InputTypes: []string{PortTypeVideo, PortTypeVideo}, OutputTypes: []string{PortTypeMetrics}},
	ActivityX264RemoteTranscode: {InputTypes: []string{}, OutputTypes: []string{PortTypeVideo}},
	ActivityRemoteFileMetric:    {InputTypes: []string{}, OutputTypes: []string{PortTypeMetrics}},
	ActivitySceneCut:            {InputTypes: []string{PortTypeVideo}, OutputTypes: []string{PortTypeSceneCut}},
	ActivityRemoteSceneCut:      {InputTypes: []string{}, OutputTypes: []string{PortTypeSceneCut}},
	ActivityTransnetV2SceneCut:  {InputTypes: []string{PortTypeVideo}, OutputTypes: []string{PortTypeSceneCut}},
	ActivitySegmentMedia:        {InputTypes: []string{PortTypeVideo}, OutputTypes: []string{PortTypeSceneCut}},
	ActivityRemoteSegmentMedia:  {InputTypes: []string{}, OutputTypes: []string{PortTypeSceneCut}},
	ActivityGenerateReport:      {InputTypes: []string{PortTypeAny}, OutputTypes: []string{PortTypeReport}},
	ActivityFMP4Repackage:       {InputTypes: []string{PortTypeVideo}, OutputTypes: []string{PortTypeVideo}},
	ActivityFetchWorkflowDAG:    {InputTypes: []string{}, OutputTypes: []string{PortTypeAny}},
	MetaSceneCutDispatch:        {InputTypes: []string{PortTypeSceneCut}, OutputTypes: []string{}},
	MetaCompositeWorkflow:       {InputTypes: []string{PortTypeAny}, OutputTypes: []string{PortTypeAny}},
}

// Activity type constants matching registered Temporal activities.
const (
	ActivityResDownload        = "ResourceDownload"
	ActivityResUpload          = "ResourceUpload"
	ActivityX264Transcode      = "x264Transcode"
	ActivityFileMetricAnalysis = "FileMetricAnalysis"
	ActivityX264RemoteTranscode = "x264RemoteTranscode"
	ActivityRemoteFileMetric   = "RemoteFileMetricAnalysis"
	ActivitySceneCut           = "SceneCut"
	ActivityRemoteSceneCut     = "RemoteSceneCut"
	ActivityTransnetV2SceneCut = "TransnetV2SceneCut"
	ActivitySegmentMedia       = "SegmentMedia"
	ActivityRemoteSegmentMedia = "RemoteSegmentMedia"
	ActivityGenerateReport     = "GenerateReport"
	ActivityFMP4Repackage      = "FragmentedMP4Repackage"
	ActivityFetchWorkflowDAG   = "FetchWorkflowDAG"
)

// Meta-node types handled by the interpreter, not registered as activities.
const (
	MetaSceneCutDispatch  = "SceneCutDispatch"
	MetaCompositeWorkflow = "CompositeWorkflow"
)

// ActivityInput is the standard input passed to every Temporal activity.
type ActivityInput struct {
	NodeID       string            `json:"node_id"`
	NodeType     string            `json:"node_type"`
	Params       map[string]string `json:"params"`
	ProjectID    string            `json:"project_id"`
	SessionGroup string            `json:"session_group,omitempty"`
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
	RunID       string `json:"run_id"`
	Status      string `json:"status"`
	CurrentNode string `json:"current_node,omitempty"`
	ProgressPct int    `json:"progress_pct"`
	Error       string `json:"error,omitempty"`
}

const (
	TemporalTaskQueue = "kvqtool-workers"
	TemporalNamespace = "default"
)
