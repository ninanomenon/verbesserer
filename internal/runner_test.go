package internal

import (
	"reflect"
	"testing"

	"github.com/ninanomenon/verbesserer/internal/report"
	"github.com/ninanomenon/verbesserer/internal/tools"
)

type mockedRunner struct{}

func (mockedrunner mockedRunner) Name() string {
	panic("not implemented") // TODO: Implement
}
func (mockedrunner mockedRunner) Description() string {
	panic("not implemented") // TODO: Implement
}
func (mockedrunner mockedRunner) Preflight() error {
	panic("not implemented") // TODO: Implement
}
func (mockedrunner mockedRunner) Run() (*[]tools.Result, error) {
	return &[]tools.Result{
		{
			Description: "Test Error",
			CheckName:   "T1337",
			Fingerprint: "ABCDEFG",
			Location: tools.Location{
				Lines: tools.Lines{
					Begin: 1,
					End:   100,
				},
				Path: "/dev/random",
			},
			Severity: "Yes",
		},
	}, nil
}

func TestRun(t *testing.T) {
	var runner = mockedRunner{}

	var toolSlice []tools.Tool
	toolSlice = append(toolSlice, runner)

	var reports = report.Reports{}
	reports["/dev/random"] = report.Report{
		FileHash: "",
		Issues: []report.Issue{
			{
				Message: "Test Error",
				Hash:    "",
				Lines: report.Lines{
					Begin: 1,
					End:   100,
				},
			},
		},
	}

	type args struct {
		tools []tools.Tool
	}

	tests := []struct {
		name      string
		args      args
		want      report.Reports
		errorWant []error
	}{
		{
			name: "",
			args: args{
				tools: toolSlice,
			},
			want:      reports,
			errorWant: []error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Execute(tt.args.tools)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.errorWant) {
				t.Errorf("Run() got1 = %v, want %v", got1, tt.errorWant)
			}
		})
	}
}
