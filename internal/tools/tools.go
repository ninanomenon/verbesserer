package tools

type Result struct {
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

type Tool interface {
	Name() string
	Description() string
	Preflight() error
	Run() (*[]Result, error)
}
