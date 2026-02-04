# Rust Reference Implementation

This directory contains the Rust reference implementation for computing canonical
SHA-256 hashes of SBOMs used for signing and verification.

## Quick links

- Implementation: [src/lib.rs](src/lib.rs)
- Test Harness: [src/tests/harness.rs](src/tests/harness.rs)
- Cargo Build Configuration: [Cargo.toml](Cargo.toml)
- Test Vectors (source of truth): [test-vectors/](../test-vectors/)
- Test Manifest: [test-vectors/manifest.json](../test-vectors/manifest.json)
- CI: [.github/workflows/rust.yml](../.github/workflows/rust.yml)

## Requirements

- Rust Version >= 1.93.0

## Quick Start

1. Install Dependencies and Build
   - `cargo build`

2. Run Tests
   - `cargo test`

## Notes

- Tests read SBOM files from the project's [/test-vectors/](../test-vectors/)
directory. Do not modify those files for test runs.
- The implementation should produce deterministic hashes across varied formats
(original / minified / pretty / canonical).

## Contributing

- Follow repository guidelines in the root [README.md](../README.md) and
[CONTRIBUTING.md](../CONTRIBUTING.md).
