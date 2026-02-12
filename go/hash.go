package hash

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

	// Apply JSF Signature Exclusions
	if bomFormat, ok := data["bomFormat"].(string); ok && bomFormat == "CycloneDX" {
		if signature, ok := data["signature"]; ok && signature != nil {
			if signatureMap, ok := signature.(map[string]interface{}); ok {

				// Handle Dynamic Exclusions (from the 'excludes' property)
				if excludes, ok := signatureMap["excludes"].([]interface{}); ok {
					for _, e := range excludes {
						if propName, ok := e.(string); ok {
							// These are excluded from the root object
							delete(data, propName)
						}
					}
				}

				// Handle JSF Core Requirement: Delete ONLY the "value" property
				// This leaves 'algorithm', 'publicKey', etc., in the hash.
				delete(signatureMap, "value")
			}
		}
	}

	// Re-marshal to JSON so we can canonicalize the raw bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal pruned JSON: %w", err)
	}

	// Apply RFC 8785 (JCS)
	canonicalBytes := jsontext.Value(jsonBytes)
	canonicalBytes.Canonicalize()

	// Hash the canonical bytes
	hash := sha256.Sum256(canonicalBytes)

	return hex.EncodeToString(hash[:]), nil
}
