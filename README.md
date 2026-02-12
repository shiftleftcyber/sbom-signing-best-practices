# SBOM Signing Best Practices

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Linter](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/lint.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/lint.yml)
[![Go](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/go.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/go.yml)
[![Java](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/java.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/java.yml)
[![JavaScript](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/javascript.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/javascript.yml)
[![Python](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/python.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/python.yml)
[![Rust](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/rust.yml/badge.svg)](https://github.com/shiftleftcyber/sbom-signing-best-practices/actions/workflows/rust.yml)

A set of reference implementations and test suites for computing canonical
hashes of Software Bill of Materials (SBOMs). This project demonstrates
multi-language interoperability to support reliable signing and integrity verification.

## Supported Standards

| Standard | Version | Format | Canonicalization | Signature Spec |
| :--- | :--- | :--- | :--- | :--- |
| **CycloneDX** | 1.7 | JSON | JCS (RFC 8785) | JSF |
| **SPDX** | 2.2 | JSON | JCS (RFC 8785) | Undefined (detached) |
| **SPDX** | 3.0 | JSON-LD | JCS (RFC 8785) | Undefined (detached) |

\* Canonical hashing for XML formats (CycloneDX and SPDX) is currently under
   evaluation and is slated for a future release.

## The Goal: Deterministic Hashing

The objective is to compute a hash of the SBOM **content**, independent of its **formatting**.

By applying **canonicalization** (JCS) and **pruning** (JSF) prior to hashing, we
ensure that the same logical data always produces the same hash. This allows
SBOMs to be transmitted in any form — whether "minified" or "pretty-printed" — while
maintaining a stable cryptographic identity across different platforms:

> `Hash(Go_Impl(sbom-pretty.json))` == `Hash(Python_Impl(sbom-min.json))`

## Technical Specification

The exact rules for property exclusion, signature handling, and serialization
are detailed in the technical spec:

- [Technical Specification: JSON SBOM Canonical Hashing](/specs/README.md)

## Directory Structure

- `go/`: Go reference implementation.
- `java/`: Java reference implementation.
- `javascript/`: JavaScript reference implementation.
- `python/`: Python reference implementation.
- `rust/`: Rust reference implementation.
- `specs/`: Technical details on the canonicalization rules used.
- `test-vectors/`: Shared "golden" SBOMs and a test manifest file including the
  expected hashes used to verify cross-language parity.

## Getting Started

Refer to the `README.md` in each language directory for build instructions.
