package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"

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
					FileHash: "", // TODO: calculate file hash
					Issues:   []report.Issue{},
				}
			}

			issue := report.Issue{
				Message: outputResult.Description,
				Lines: report.Lines{
					Begin: outputResult.Location.Lines.Begin,
					End:   outputResult.Location.Lines.End,
				},
			}

			issue.Hash = generateFingerprint(issue)
			reportData.Issues = append(reportData.Issues, issue)

			reports[outputResult.Location.Path] = reportData
		}
	}

	return reports, errors
}

func generateFingerprint(issue report.Issue) string {
	hash := sha256.New()

	values := reflect.ValueOf(issue)
	for i := range values.NumField() {
		field := values.Field(i)
		hash.Write([]byte(field.String()))
	}

	return hex.EncodeToString(hash.Sum(nil))
}
