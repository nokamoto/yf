package yaml

import (
	"bytes"
	"fmt"

	"gopkg.in/yaml.v3"
)

var errUnmarshal = fmt.Errorf("failed to unmarshal yaml")

// Format formats the input yaml with the given number of spaces for indentation.
// It returns the formatted yaml.
func Format(in []byte, spaces int) ([]byte, error) {
	var n yaml.Node
	if err := yaml.Unmarshal(in, &n); err != nil {
		return nil, fmt.Errorf("%w: %w", errUnmarshal, err)
	}

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(spaces)
	if err := enc.Encode(&n); err != nil {
		return nil, fmt.Errorf("failed to encode yaml: %w", err)
	}
	return buf.Bytes(), nil
}
