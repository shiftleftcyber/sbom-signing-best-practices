package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/json/jsontext"
	"fmt"
)

// ComputeJSONHash calculates the SHA-256 hash of a JSON SBOM
func ComputeJSONHash(input []byte) (string, error) {
	// Parse JSON
	var data map[string]interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		return "", fmt.Errorf("Failed to parse JSON: %w", err)
	}

	// 1. Apply Exclusions (only for CycloneDX embedded signatures)
	if f, ok := data["bomFormat"].(string); ok && f == "CycloneDX" {
		// Check if there's a signature block
		if signature := data["signature"]; signature != nil {
			if m, ok := signature.(map[string]interface{}); ok {
				// Extract any additional exclusions from the signature block
				var exclusions []string
				if rawExcludes, ok := m["excludes"].([]interface{}); ok {
					for _, e := range rawExcludes {
						if s, ok := e.(string); ok {
							exclusions = append(exclusions, s)
						}
					}
				}
				// Always remove signature property before hashing
				exclusions = append(exclusions, "signature")

				// Apply all exclusions
				for _, property := range exclusions {
					delete(data, property)
				}
			}
		}
	}

	// 2. Re-marshal to JSON so we can canonicalize the raw bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal pruned JSON: %w", err)
	}

	// 3. Apply RFC 8785 (JCS)
	canonicalBytes := jsontext.Value(jsonBytes)
	canonicalBytes.Canonicalize()

	// 4. Hash the canonical bytes
	hash := sha256.Sum256(canonicalBytes)

	return hex.EncodeToString(hash[:]), nil
}
