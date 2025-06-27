package tools

type Tool interface {
	Name() string
	Description() string
	Run() (string, error)
}
