package tools

type ReportFormat struct {
	Description string `json:"description"`
	CheckName   string `json:"check_name"`
	Fingerprint string `json:"fingerprint"`
	Location    struct {
		Lines struct {
			Begin int `json:"begin"`
			End   int `json:"end"`
		}
		Path string `json:"path"`
	} `json:"location"`
	Severity string `json:"severity,omitempty"`
}
