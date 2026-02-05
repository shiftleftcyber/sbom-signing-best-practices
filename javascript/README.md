# JavaScript Reference Implementation

This directory contains the JavaScript reference implementation for computing canonical
SHA-256 hashes of SBOMs used for signing and verification.

## Quick links

- Implementation: [hash.js](hash.js)
- Test Harness: [hash.test.js](hash.test.js)
- NPM Project Build Configuration: [package.json](package.json)
- Test Vectors (source of truth): [test-vectors/](../test-vectors/)
- Test Manifest: [test-vectors/manifest.json](../test-vectors/manifest.json)
- CI: [.github/workflows/javascript.yml](../.github/workflows/javascript.yml)

## Requirements

- Node 25

## Quick Start

1. Install Dependencies
   - `npm install`

2. Run Tests
   - `npm test`

## Notes

- Tests read SBOM files from the project's [/test-vectors/](../test-vectors/)
  directory. Do not modify those files for test runs.
- The implementation should produce deterministic hashes across varied formats
  (original / minified / pretty / canonical).

## Contributing

- Follow repository guidelines in the root [README.md](../README.md) and
  [CONTRIBUTING.md](../CONTRIBUTING.md).
