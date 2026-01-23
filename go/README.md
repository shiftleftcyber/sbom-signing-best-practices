# Go Reference Implementation

This directory contains the Go reference implementation for computing canonical
SHA-256 hashes of SBOMs used for signing and verification.

## Quick links

- Implementation: [hash.go](hash.go)
- Test Harness: [hash_test.go](hash_test.go)
- Go Module: [go.mod](go.mod)
- Test Vectors (source of truth): [test-vectors/](../test-vectors/)
- Test Manifest: [test-vectors/manifest.json](../test-vectors/manifest.json)
- CI: [.github/workflows/go.yml](../.github/workflows/go.yml)

## Requirements

- Go toolchain matching the module: go 1.25.6 (see [go.mod](go.mod))
- GOEXPERIMENT=jsonv2 environment variable (CI enables this)

## Quick Start

1. Build
   - `go build`

2. Run Tests
   - `go test -v`

## Notes

- Tests read SBOM files from the project's [/test-vectors/](../test-vectors/)
directory. Do not modify those files for test runs.
- The implementation should produce deterministic hashes across varied formats
(original / minified / pretty / canonical).

## Contributing

- Follow repository guidelines in the root [README.md](../README.md) and
[CONTRIBUTING.md](../CONTRIBUTING.md).
