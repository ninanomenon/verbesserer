package internal

import (
	"github.com/ninanomenon/verbesserer/internal/report"
	"github.com/ninanomenon/verbesserer/internal/tools"
)

// Run function takes a slice of `Tool` and executes the `Run` function of each.
func Run(tools []tools.Tool) (*report.Reports, []error) {
	var errors = []error{}
	var reports = report.Reports{}

	for _, tool := range tools {
		result, err := tool.Run()
		if err != nil {
			errors = append(errors, err)
			continue
		}

		for _, outputResult := range *result {
			issueReport := getReport(&reports, outputResult.Location.Path)
			issues := issueReport.Issues

			*issues = append(*issues, report.Issue{
				Message: outputResult.Description,
				Hash:    "", // Calculate hash
				Lines: report.Lines{
					Begin: outputResult.Location.Lines.Begin,
					End:   outputResult.Location.Lines.End,
				},
			})
		}
	}

	return &reports, errors
}

func getReport(reports *report.Reports, filePath string) *report.Report {
	for _, report := range *reports {
		if report.FilePath == filePath {
			return &report
		}
	}

	report := report.Report{
		FilePath: filePath,
		FileHash: "", // TODO: Calculate filehash
		Issues:   &[]report.Issue{},
	}

	*reports = append(*reports, report)

	return &report
}
