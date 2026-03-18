package types

// SceneCutFile is the standard output from scene detection and segmentation nodes.
type SceneCutFile struct {
	Version  string    `json:"version"`
	Source   string    `json:"source"`
	Segments []Segment `json:"segments"`
}

// Segment represents a single time range within a media file.
type Segment struct {
	Index     int    `json:"index"`
	StartTime string `json:"start_time"` // "HH:MM:SS.mmm"
	EndTime   string `json:"end_time"`
}

// MetricReports is the standardized output from metric analysis nodes.
// Outer map key = distorted file index (string), inner map key = frame index (string), value = metric value.
type MetricReports struct {
	Header MetricHeader                 `json:"header"`
	Vmaf   map[string]map[string]string `json:"vmaf,omitempty"`
	Ssim   map[string]map[string]string `json:"ssim,omitempty"`
	Psnr   map[string]map[string]string `json:"psnr,omitempty"`
}

// MetricHeader contains metadata about a metric analysis run.
type MetricHeader struct {
	Version   string            `json:"version"`
	Metrics   []string          `json:"metrics"`
	Reference string            `json:"reference"`
	Dist      map[string]string `json:"dist"`
}
