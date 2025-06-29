package report

type Report struct {
	FilePath string   `json:"file_path"`
	FileHash string   `json:"file_hash"`
	Issues   *[]Issue `json:"issues"`
}

type Issue struct {
	Message string `json:"message"`
	Hash    string `json:"hash"`
	Lines   Lines  `json:"lines"`
}

type Lines struct {
	Begin int `json:"begin"`
	End   int `json:"end,omitempty"`
}

type Reports []Report

func (r Reports) GenerateReport() []byte {
	return make([]byte, 0)
}

func (r Reports) WriteToml() error {
	return nil
}
