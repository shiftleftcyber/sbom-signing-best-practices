# Contributing to SBOM Canonical Hash

Thank you for your interest in contributing! This project aims to demonstrate
interoperable SBOM hashing across multiple languages.

## Project Structure

This is a monorepo. Each language implementation lives in its own directory:

- `/go` - Go module
- `/python` - Python package
- `/rust` - Rust crate
- ...

## Development Guidelines

### 1. Test Vectors are the Source of Truth

- Do not modify test vectors inside language directories.
- The "Golden" test vectors are located in `/test-vectors`.
- All implementations must pass the suite defined in `/test-vectors/manifest.json`.

### 2. Adding a New Language

- Create a new directory for the language.
- Implement the canonicalization logic (JCS for JSON, XML-C14N for XML).
- Ensure your tests read from the root `/test-vectors` directory.
- Add a CI job to `.github/workflows/ci.yml`.

### 3. Reporting Issues

- Please specify which language implementation you are using (or if it is a
  flaw in the core spec).

## Documentation Style

We use `markdownlint` to keep our documentation consistent.

- You can run it locally using the
  [markdownlint-cli](https://github.com/igorshubovych/markdownlint-cli).
- CI will fail if your PR contains linting errors.
- Configuration is found in `.markdownlint.json`.
