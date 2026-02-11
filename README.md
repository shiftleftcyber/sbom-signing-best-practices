# SBOM Signing Best Practices

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Linter](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/lint.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/lint.yml)
[![Go](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/go.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/go.yml)
[![Java](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/java.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/java.yml)
[![JavaScript](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/javascript.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/javascript.yml)
[![Python](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/python.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/python.yml)
[![Rust](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/rust.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/rust.yml)

A set of reference implementations and test suites for computing the **canonical
hashes** of SBOMs (Software Bill of Materials). This project demonstrates
**interoperability** across multiple languages to support reliable **SBOM signing**
and **integrity verification**.

## Supported Standards

| Standard | Version | Formats | Canonicalization Method |
| :--- | :--- | :--- | :--- |
| **CycloneDX** | 1.7 | JSON | RFC 8785 (JCS) |
| **CycloneDX** | 1.7 | XML | W3C Exclusive C14N * |
| **SPDX** | 2.2 | JSON | RFC 8785 (JCS) |
| **SPDX** | 2.2 | XML | W3C Exclusive C14N * |
| **SPDX** | 3.0 | JSON-LD | RFC 8785 (JCS) |
| **SPDX** | 3.0 | XML | W3C Exclusive C14N * |

\* XML Work is still TODO.

## Goal

The goal is **Deterministic Hashing**: to compute a hash of the SBOM content
itself, not its formatting.

Deterministic hashing is accomplished by using **canonicalization** prior to
computing the hash. This ensures that the same logical content always produces
the same hash, regardless of formatting differences.

This approach enables SBOMs to be transmitted over the wire in any form (for
example, either "minified" or "pretty-printed" JSON) while the computed hash
remains identical as long as the actual content has not changed. This allows for
reliable **SBOM signing and integrity verification** across different
implementations, representations, and platforms:

> `Hash(Go_Impl(sbom-pretty.json))` == `Hash(Python_Impl(sbom-min.json))`

## Directory Structure

* `go/`: Go reference implementation.
* `java/`: Java reference implementation.
* `javascript/`: JavaScript reference implementation.
* `python/`: Python reference implementation.
* `rust/`: Rust reference implementation.
* `specs/`: Technical details on the canonicalization rules used.
* `test-vectors/`: The shared "golden" set of SBOMs and their expected hashes.

## Getting Started

See the `README.md` in each language directory for build instructions.
