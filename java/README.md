# Java Reference Implementation

This directory contains the Java reference implementation for computing canonical
SHA-256 hashes of SBOMs used for signing and verification.

## Quick links

- Implementation: [src/main/java/io/shiftleftcyber/Hash.java](src/main/java/io/shiftleftcyber/Hash.java)
- Test Harness: [src/test/java/io/shiftleftcyber/HashTest.java](src/test/java/io/shiftleftcyber/HashTest.java)
- Maven Project Build Configuration: [pom.xml](pom.xml)
- Test Vectors (source of truth): [test-vectors/](../test-vectors/)
- Test Manifest: [test-vectors/manifest.json](../test-vectors/manifest.json)
- CI: [.github/workflows/java.yml](../.github/workflows/java.yml)

## Requirements

- Java JDK 25

## Quick Start

1. Build
   - `mvn compile`

2. Run Tests
   - `mvn test`

## Notes

- Tests read SBOM files from the project's [/test-vectors/](../test-vectors/)
  directory. Do not modify those files for test runs.
- The implementation should produce deterministic hashes across varied formats
  (original / minified / pretty / canonical).

## Contributing

- Follow repository guidelines in the root [README.md](../README.md) and
  [CONTRIBUTING.md](../CONTRIBUTING.md).
