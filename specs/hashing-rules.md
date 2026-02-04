# Technical Specification: SBOM Canonical Hashing

## 1. Input Processing

Implementations must accept a file path or a byte stream.

## 2. Format Detection

The implementation should detect the format based on the following:

- **JSON:** Must start with `{`.
- **XML:** Must start with `<?xml` or `<`.

## 3. Canonicalization Rules

### 3.1 JSON (CycloneDX & SPDX)

All JSON inputs must be processed according to
**[RFC 8785 (JSON Canonicalization Scheme)](https://www.rfc-editor.org/rfc/rfc8785)**.

- Properties must be sorted lexicographically.
- All insignificant whitespace must be removed.
- Use UTF-8 encoding.

### 3.2 XML (CycloneDX & SPDX)

All XML inputs must be processed according to
**[c14n11 (Canonical XML Version 1.1)](https://www.w3.org/TR/xml-c14n11/)**.

- Use Exclusive XML Canonicalization to ensure namespace portability.

## 4. Hashing

- **Algorithm:** SHA-256.
- **Output:** Hexadecimal string (lowercase).

## 5. Verification Flow

1. Load SBOM.
2. Identify Format (JSON/XML).
3. Apply property exclusion (for CycloneDX only)
4. Apply Canonicalization (JCS or C14N).
5. Compute SHA-256 of the canonicalized bytes.
6. Return Hex string.

## 6. Exclusion Rules (CycloneDX)

To support deterministic hashing for signing, the following JSF exclusion logic
applies:

1. **Default Exclusions:** The `signature` property (if present) is removed
before hashing.
2. **User-Defined Exclusions:** Implementations must support the `excludes`
property (if present) in JSF to remove other specified fields before hashing.
3. **Ordering:** Exclusions must be processed **before** the JCS or C14N
canonicalization.
