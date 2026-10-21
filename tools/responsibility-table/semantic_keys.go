package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func loadSemanticKeyLookup(sourceRoot, explicitPath string) (map[string]string, error) {
	if explicitPath == "" {
		defaultPath := filepath.Join(sourceRoot, ".json-like-comment-outs", "semantic-keys.json")
		data, err := os.ReadFile(defaultPath)
		if os.IsNotExist(err) {
			return buildSemanticKeyLookup(defaultSemanticKeys())
		}
		if err != nil {
			return nil, err
		}
		config, err := decodeSemanticKeys(data)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", defaultPath, err)
		}
		return buildSemanticKeyLookup(config)
	}

	configPath, err := filepath.Abs(explicitPath)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", configPath, err)
	}
	config, err := decodeSemanticKeys(data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", configPath, err)
	}
	return buildSemanticKeyLookup(config)
}

func defaultSemanticKeys() map[string][]string {
	return map[string][]string{
		"Responsibility": {"責務", "役割", "Role", "Responsible"},
	}
}

func decodeSemanticKeys(data []byte) (map[string][]string, error) {
	var root map[string]json.RawMessage
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	if root == nil {
		return nil, fmt.Errorf("configuration must be a JSON object")
	}

	var config map[string][]string
	for key, value := range root {
		if strings.EqualFold(key, "SemanticKeys") {
			if err := json.Unmarshal(value, &config); err != nil {
				return nil, fmt.Errorf("SemanticKeys must map canonical keys to alias arrays: %w", err)
			}
			break
		}
	}
	if config == nil {
		if err := json.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("configuration must map canonical keys to alias arrays: %w", err)
		}
	}
	if config == nil {
		return nil, fmt.Errorf("configuration must not be null")
	}
	return config, nil
}

func buildSemanticKeyLookup(config map[string][]string) (map[string]string, error) {
	keys := make([]string, 0, len(config))
	for canonical := range config {
		keys = append(keys, canonical)
	}
	sort.Strings(keys)

	lookup := make(map[string]string)
	for _, canonical := range keys {
		canonicalKey := normalizeSemanticKey(canonical)
		if canonicalKey == "" {
			return nil, fmt.Errorf("canonical keys must not be empty")
		}
		if previous, exists := lookup[canonicalKey]; exists && !strings.EqualFold(previous, canonical) {
			return nil, fmt.Errorf("canonical key %q conflicts with %q", canonical, previous)
		}
		lookup[canonicalKey] = canonical

		aliases := append([]string(nil), config[canonical]...)
		sort.Strings(aliases)
		for _, alias := range aliases {
			normalized := normalizeSemanticKey(alias)
			if normalized == "" {
				return nil, fmt.Errorf("aliases for %q must not be empty", canonical)
			}
			if previous, exists := lookup[normalized]; exists && !strings.EqualFold(previous, canonical) {
				return nil, fmt.Errorf("alias %q maps to both %q and %q", alias, previous, canonical)
			}
			lookup[normalized] = canonical
		}
	}
	return lookup, nil
}

func normalizeSemanticKey(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func extractResponsibility(comment string, semanticLookup map[string]string) ([]tableRow, bool, error) {
	body := strings.TrimSpace(comment)
	if !strings.HasPrefix(body, "{") {
		return nil, false, nil
	}

	var object map[string]json.RawMessage
	if err := json.Unmarshal([]byte(body), &object); err != nil {
		candidate := mentionsResponsibilityKey(body, semanticLookup)
		if !candidate {
			return nil, false, nil
		}
		return nil, true, fmt.Errorf("invalid JSON comment: %w", err)
	}
	if object == nil {
		return nil, false, nil
	}

	keys := make([]string, 0, len(object))
	for key := range object {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var value json.RawMessage
	found := false
	for _, key := range keys {
		if strings.EqualFold(semanticLookup[normalizeSemanticKey(key)], "Responsibility") {
			if found {
				return nil, true, fmt.Errorf("multiple keys resolve to Responsibility")
			}
			value = object[key]
			found = true
		}
	}
	if !found {
		return nil, false, nil
	}

	trimmedValue := bytes.TrimSpace(value)
	if len(trimmedValue) == 0 || trimmedValue[0] != '{' {
		return nil, true, fmt.Errorf("Responsibility value must be an object")
	}
	var responsibilities map[string]string
	if err := json.Unmarshal(value, &responsibilities); err != nil {
		return nil, true, fmt.Errorf("Responsibility values must be strings: %w", err)
	}

	names := make([]string, 0, len(responsibilities))
	for name := range responsibilities {
		names = append(names, name)
	}
	sort.Strings(names)
	rows := make([]tableRow, 0, len(names))
	for _, name := range names {
		rows = append(rows, tableRow{name: name, responsibility: responsibilities[name]})
	}
	return rows, true, nil
}

func mentionsResponsibilityKey(body string, lookup map[string]string) bool {
	lowerBody := strings.ToLower(body)
	for key, canonical := range lookup {
		if strings.EqualFold(canonical, "Responsibility") && strings.Contains(lowerBody, key) {
			return true
		}
	}
	return false
}
