package tools

type ReportFormat struct {
	Description string   `json:"description"`
	CheckName   string   `json:"check_name"`
	Fingerprint string   `json:"fingerprint"`
	Location    Location `json:"location"`
	Severity    string   `json:"severity"`
}

type Location struct {
	Lines Lines  `json:"lines"`
	Path  string `json:"path"`
}

type Lines struct {
	Begin int `json:"begin"`
	End   int `json:"end"`
}
