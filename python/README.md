# Python Reference Implementation

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Python](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/python.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/python.yml)

This directory contains the Python reference implementation for computing canonical
SHA-256 hashes of SBOMs used for signing and verification.

## Quick links

- Implementation: [hash.py](hash.py)
- Test Harness: [test_hash.py](test_hash.py)
- Test Vectors (source of truth): [test-vectors/](../test-vectors/)
- Test Manifest: [test-vectors/manifest.json](../test-vectors/manifest.json)
- CI: [.github/workflows/python.yml](../.github/workflows/python.yml)

## Requirements

- Python 3

## Quick Start

1. Install Dependencies
   - `pip install -r requirements.txt`

2. Run Tests
   - `python test -v`

## Notes

- Tests read SBOM files from the project's [/test-vectors/](../test-vectors/)
  directory. Do not modify those files for test runs.
- The implementation should produce deterministic hashes across varied formats
  (original / minified / pretty / canonical).

## Contributing

- Follow repository guidelines in the root [README.md](../README.md) and
  [CONTRIBUTING.md](../CONTRIBUTING.md).
