package internal

import (
	"github.com/ninanomenon/verbesserer/internal/report"
	"github.com/ninanomenon/verbesserer/internal/tools"
)

// Run function takes a slice of `Tool` and executes the `Run` function of each.
func Run(tools []tools.Tool) (report.Reports, []error) {
	var errors = []error{}
	var reports = report.Reports{}

	for _, tool := range tools {
		result, err := tool.Run()
		if err != nil {
			errors = append(errors, err)
			continue
		}

		for _, outputResult := range *result {
			reportData, ok := reports[outputResult.Location.Path]
			if !ok {
				reportData = report.Report{
					FilePath: outputResult.Location.Path,
					FileHash: "", // TODO: calculate file hash
					Issues:   []report.Issue{},
				}
			}

			reportData.Issues = append(reportData.Issues, report.Issue{
				Message: outputResult.Description,
				Hash:    "", // TODO: Calculate hash
				Lines: report.Lines{
					Begin: outputResult.Location.Lines.Begin,
					End:   outputResult.Location.Lines.End,
				},
			})

			reports[outputResult.Location.Path] = reportData
		}
	}

	return reports, errors
}
