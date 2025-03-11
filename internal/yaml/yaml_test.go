package yaml

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Format(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		spaces  int
		want    string
		wantErr error
	}{
		{
			name:    "error if invalid yaml",
			in:      "invalid: yaml: text",
			wantErr: errUnmarshal,
		},
		{
			name: "format yaml with 2 spaces",
			in: `key:
    value: x
`,
			spaces: 2,
			want: `key:
  value: x
`,
		},
		{
			name: "append newline at the end",
			in:   `key: value`,
			want: `key: value
`,
		},
		{
			name: "preserve comments",
			in: `# comment
key: value
`,
			want: `# comment
key: value
`,
		},
		{
			name: "empty lines are removed",
			in: `key: value

key2: value2
`,
			want: `key: value
key2: value2
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Format([]byte(tt.in), tt.spaces)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(string(got), tt.want); diff != "" {
				t.Errorf("Format() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
