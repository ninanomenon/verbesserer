package tools

type Tool interface {
	Name() string
	Description() string
	Run() (*[]ReportFormat, error)
}
