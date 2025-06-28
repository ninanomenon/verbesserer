package tools

import "fmt"

type Tool interface {
	Name() string
	Description() string
	Run() (*[]ReportFormat, error)
}

// Run function takes a slice of `Tool` and executes the
// `Run` function of each.
func Run(tools []Tool) {
	for _, tool := range tools {
		output, err := tool.Run()
		if err != nil {
			fmt.Printf("Error while running %s: %s\n", tool.Name(), err)
			continue
		}

		fmt.Printf("Result of %s: %v\n", tool.Name(), output)
	}
}
