package plugins

import (
	"fmt"
	"plugin"
)

// LoadPolicyPlugin loads a plugin from the specified path and returns the Policy interface.
func LoadPolicyPlugin(path string) (Policy, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, fmt.Errorf("plugin open failed: %w", err)
	}

	sym, err := p.Lookup("Plugin")
	if err != nil {
		return nil, fmt.Errorf("plugin lookup failed: %w", err)
	}

	policy, ok := sym.(Policy)
	if !ok {
		return nil, fmt.Errorf("invalid plugin type")
	}

	return policy, nil
}
