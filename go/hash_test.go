package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type TestManifest []struct {
	Name          string `json:"name"`
	File          string `json:"file"`
	FileMin       string `json:"file-min"`
	FilePretty    string `json:"file-pretty"`
	FileCanonical string `json:"file-canonical"`
	Format        string `json:"format"`
	SHA256        string `json:"sha256"`
	Source        string `json:"source"`
}

func TestHash(t *testing.T) {
	// Path to the shared vectors at the root of the repo
	manifestPath := filepath.Join("..", "test-vectors", "manifest.json")

	content, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to load manifest: %v", err)
	}

	var manifest TestManifest
	json.Unmarshal(content, &manifest)

	for _, testCase := range manifest {
		//t.Run(testCase.Name, func(t *testing.T) {
		if testCase.Format != "json" {
			t.Skip("XML not yet implemented")
		}

		files := []struct {
			name string
			path string
		}{
			{"original", testCase.File},
			{"min", testCase.FileMin},
			{"pretty", testCase.FilePretty},
			{"canonical", testCase.FileCanonical},
		}

		for _, f := range files {
			t.Run(testCase.Name+"_"+f.name, func(t *testing.T) {
				path := filepath.Join("..", "test-vectors", f.path)
				bytes, _ := os.ReadFile(path)

				computedHash, err := ComputeJSONHash(bytes)
				if err != nil {
					t.Fatalf("Hash error: %v", err)
				}

				if computedHash != testCase.SHA256 {
					t.Errorf("Hash Mismatch!\nExpected: %s\nGot: %s\nFile: %s", testCase.SHA256, computedHash, f.path)
				}
			})
		}
		//})
	}
}
