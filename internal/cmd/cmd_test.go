package cmd

import (
	"errors"
	"testing"
)

func Test_yf(t *testing.T) {
	tests := []struct {
		name    string
		wantErr error
	}{
		{
			name:    "not implemented",
			wantErr: errNotImplemented,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yf := New()
			err := yf.Execute()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("yf.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
