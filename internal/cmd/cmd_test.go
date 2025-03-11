package cmd

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_yf(t *testing.T) {
	tests := []struct {
		name    string
		stdin   io.Reader
		stdout  string
		wantErr error
	}{
		{
			name:  "format yaml from stdin",
			stdin: bytes.NewBufferString(`key: value`),
			stdout: `key: value
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yf := New()
			var stdout bytes.Buffer
			yf.SetIn(tt.stdin)
			yf.SetOut(&stdout)
			err := yf.Execute()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("yf.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.stdout, stdout.String()); diff != "" {
				t.Errorf("yf.Execute() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
