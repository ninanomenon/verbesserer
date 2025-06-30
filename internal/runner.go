package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/ninanomenon/verbesserer/internal/report"
	"github.com/ninanomenon/verbesserer/internal/tools"
)

// Execute function takes a slice of `Tool` and executes the `Execute` function of each.
func Execute(tools []tools.Tool) (report.Reports, []error) {
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
				hash, err := generateFileFingerprint(outputResult.Location.Path)
				if err != nil {
					errors = append(errors, err)
				}

				reportData = report.Report{
					FileHash: hash,
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

			issue.Hash = generateIssueFingerprint(issue)
			reportData.Issues = append(reportData.Issues, issue)

			reports[outputResult.Location.Path] = reportData
		}
	}

	return reports, errors
}

func generateIssueFingerprint(issue report.Issue) string {
	hash := sha256.New()

	hash.Write([]byte(issue.Message))
	hash.Write(fmt.Appendf(nil, "%d-%d", issue.Lines.Begin, issue.Lines.End))

	return hex.EncodeToString(hash.Sum(nil))
}

func generateFileFingerprint(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Fatal error while calculating file hash: %s", err.Error()))
		}
	}()

	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
