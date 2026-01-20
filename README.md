# SBOM Signing Best Practices

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

A reference implementation and test suite for computing **canonical hashes** of
SBOMs (Software Bill of Materials). This project demonstrates
**interoperability** across languages to support reliable **SBOM signing**
and integrity verification.

## Supported Standards

| Standard | Version | Formats | Canonicalization Method |
| :--- | :--- | :--- | :--- |
| **CycloneDX** | 1.7 | JSON | RFC 8785 (JCS) |
| **CycloneDX** | 1.7 | XML | W3C Exclusive C14N |
| **SPDX** | 3.0 | JSON-LD | RFC 8785 (JCS) * |
| **SPDX** | 3.0 | XML | W3C Exclusive C14N |

## Goal

The goal is **Deterministic Hashing**: to compute a hash of the SBOM content
itself, not its formatting.

Deterministic hashing is accomplished by **normalizing the data using
canonicalization** prior to computing the hash. This ensures that the same logical
content always produces the same hash, regardless of formatting differences.

This approach enables SBOMs to be transmitted over the wire in any form (for
example, either "minified" or "pretty-printed" JSON) while the computed hash
remains identical as long as the actual content has not changed. This allows for
reliable **SBOM signing and integrity verification** across different
implementations and representations:

> `Hash(Go_Impl(sbom.json))` == `Hash(Python_Impl(sbom.json))`

## Directory Structure

* `test-vectors/`: The shared "golden" set of SBOMs and their expected hashes.
* `specs/`: Technical details on the normalization rules used.
* `go/`: Go implementation.
* `python/`: Python implementation.
* ... (other languages)

## Getting Started

See the `README.md` in each language directory for build instructions.
