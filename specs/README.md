# Technical Specification: JSON SBOM Canonical Hashing

For consistency and simplicity, the same method is used for both CycloneDX
and SPDX JSON SBOMs to compute the hash for signing and verification. The main
difference is property exclusion and signature value removal is not required for
SPDX as SPDX signatures are stored as detached signatures.

## 1. Parse JSON

Implementations MUST be able to parse the input SBOM into a JSON data structure.

## 2. Apply JSF Signature Exclusions

To support self-signing integrity, the data model must be pruned before hashing.
For CycloneDX SBOMs, apply the following JSON Signature Format (JSF) rules to
the in-memory data structure:

1. **Top-Level Exclusions**: Any property name listed in the `signature.excludes`
array MUST be deleted from the document root.
2. **Signature Value Removal**: The `value` property inside the `signature`
object MUST be deleted. All other signature metadata (e.g., `algorithm`,
`publicKey`) remains to ensure the signature context is protected by the hash.

## 3. Apply JCS (RFC 8785)

The pruned data structure is transformed into a deterministic byte stream using
the JSON Canonicalization Scheme (JCS). This ensures that formatting is ignored
and only semantic content is hashed:

- **Whitespace Elimination**: All insignificant whitespace is stripped.
- **Data Normalization**: Uniform representation for serialization of primitive
  data types (e.g., `1.0`, `1e0`, and `1` are treated identically).
- **Strict Key Sorting**: Properties are sorted lexicographically (alphabetically).
- **Encoding**: The final output is UTF-8 encoded.

## 4. Hashing

For simplicity, all reference implementations use the SHA-256 hashing algorithm
and the hash output is returned as a hexadecimal string.

## Current Known Limitations / Future Improvements

- The reference implementations currently support the JSF simple signature format
 (signaturecore) only.  Multiple signatures (multisignature) and signature chains
 (signaturechain) are not yet supported.
- Hash algorithms are currently hard-coded to use SHA-256.  An improvement could
  be made to support multiple algorithms.
- We are using the same JCS process for both CycloneDX and SPDX, however JSF property
  exclusion is only applied for CycloneDX SBOMs.
- SBOM validation could be applied initially to better ensure the provided input
  SBOMs strictly conform to the CycloneDX or SPDX schemas.

## Notes

- JSON Signature Format (JSF) has been updated and formalized by the ITU-T as
  the X.590 Standard, and now referred to as the JSON Signature Scheme (JSS).

## References

- [https://cyclonedx.org/docs/1.7/json/#signature](https://cyclonedx.org/docs/1.7/json/#signature)
- [https://spdx.dev/use/specifications/](https://spdx.dev/use/specifications/)
- [https://cyberphone.github.io/doc/security/jsf.html](https://cyberphone.github.io/doc/security/jsf.html)
- [https://www.rfc-editor.org/rfc/rfc8785](https://www.rfc-editor.org/rfc/rfc8785)
- [https://www.itu.int/rec/T-REC-X.590-202310-I/en](https://www.itu.int/rec/T-REC-X.590-202310-I/en)
