package ruff_test

import (
	"reflect"
	"testing"

	"github.com/ninanomenon/verbesserer/internal/tools"
	"github.com/ninanomenon/verbesserer/internal/tools/ruff"
)

func TestRuff_Name(t *testing.T) {
	t.Run("check name", func(t *testing.T) {
		r := ruff.Ruff{CheckPath: ""}
		if got := r.Name(); got != "Ruff" {
			t.Errorf("Ruff.Name() = %v, want Ruff", got)
		}
	})
}

func TestRuff_Description(t *testing.T) {
	t.Run("check description", func(t *testing.T) {
		r := ruff.Ruff{CheckPath: ""}
		if got := r.Description(); got != "An extremely fast Python linter and code formatter." {
			t.Errorf("Ruff.Description() = %v, want %v", got, "An extremely fast Python linter and code formatter.")
		}
	})
}

func TestRuff_Preflight(t *testing.T) {
	t.Run("FIXME", func(t *testing.T) {
		_ = ruff.Ruff{CheckPath: ""}
	})
}

func TestRuff_Run(t *testing.T) {
	type fields struct {
		CheckPath string
	}

	tests := []struct {
		name    string
		fields  fields
		want    *[]tools.Result
		wantErr bool
	}{
		{
			name: "Test with error report from ruff",
			fields: fields{
				CheckPath: "test_data/",
			},
			want: &[]tools.Result{
				{
					Description: "Undefined name `Float`",
					CheckName:   "F821",
					Fingerprint: "ba4223675cfa63d0",
					Location: tools.Location{
						Lines: tools.Lines{
							Begin: 4,
							End:   4,
						},
						Path: "test_data/main.py",
					},
					Severity: "major",
				},
				{
					Description: "Undefined name `Any`",
					CheckName:   "F821",
					Fingerprint: "6f3be4c40f1cb394",
					Location: tools.Location{
						Lines: tools.Lines{
							Begin: 5,
							End:   5,
						},
						Path: "test_data/main.py",
					},
					Severity: "major",
				},
			},
			wantErr: false,
		},
		{
			name: "Wrong check path",
			fields: fields{
				CheckPath: "lorem impsum/",
			},
			want: &[]tools.Result{
				{
					Description: "No such file or directory (os error 2)",
					CheckName:   "E902",
					Fingerprint: "675ff1e5c93b7e1e",
					Location: tools.Location{
						Lines: tools.Lines{
							Begin: 1,
							End:   1,
						},
						Path: "lorem impsum",
					},
					Severity: "major",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ruff.Ruff{
				CheckPath: tt.fields.CheckPath,
			}
			got, err := r.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("Ruff.Run() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ruff.Run() = %#v, want %v", got, tt.want)
			}
		})
	}
}
