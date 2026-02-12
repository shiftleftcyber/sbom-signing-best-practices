# Test Vectors: SBOM Canonical Hashing

This directory contains the "golden" set of test vectors used to verify
interoperability between the different language implementations (Go, Java,
JavaScript, Python, Rust).

## Purpose

The goal of these test vectors is to ensure that regardless of the input
formatting — whether the SBOM is **pretty-printed**, **minified**, or already
**canonicalized** — the resulting SHA-256 hash is identical across all platforms
and implementations.

### Validation Workflow

Every reference implementation in this project is tested against the
`manifest.json` found in this directory.  A test is considered successful if:

1. The implementation parses the source file successfully.
2. The implementation applies the [Technical Specifications](../specs/README.md)
   successfully.
3. The resulting SHA-256 hex string matches the value recorded in the manifest.

## Test Vector Manifest

| Name | Standard | Format | Expected SHA-256 Hash |
| :--- | :--- | :--- | :--- |
| **Authenticity Verification** | CycloneDX 1.7 | JSON | `1e330add9d6b0808bf788a8a001c5dfc988f6148c8e0356c6b8200b882697ead` |
| **Identify Known Vulnerabilities** | CycloneDX 1.7 | JSON | `a206545ef31201222f9dcb9a264894655a707d693228d43e878f11670867095e` |
| **Integrity Verification** | CycloneDX 1.7 | JSON | `89971d8dac727f7b917936c3d757b0816bdfec1d183718ff95b1763c7e76dbea` |
| **Optimize Remediation Efforts** | CycloneDX 1.7 | JSON | `93d6a655ed6dde5d448eb965d208aea3174991d7193bdd23de02478c53997690` |
| **Pedigree** | CycloneDX 1.7 | JSON | `9462a4f0af951c8a002b813242bc5de8188ec430a928a4b912615f3d63deb638` |
| **Provenance** | CycloneDX 1.7 | JSON | `96889c27f7f1486cbd7d5b96b3af45e6e098cab6e996e5c31a36fe37aa26265c` |
| **Vulnerability Disclosure** | CycloneDX 1.7 | JSON | `d85921ac45fcc818592d205ef3e61cc558cbd6b71c79b26a31e4930f55c02740` |
| **Vulnerability Exploitability** | CycloneDX 1.7 | JSON | `3f7ea314603d9f10149637e237d80e653fd663a727a9e1d1e60c3f92c451d5f0` |
| **Package SBOM** | SPDX 3.0 | JSON | `a3cdd5ff28ae4c4d3bdbbaf2bb8ca669ded66833861764507257573ff2798faf` |
| **Full Example** | SPDX 3.0 | JSON | `8f04e9472e65f3a8561b143a7f70b300568829bf42cb0f64bb41132fe098a8e8` |
| **SPDX Example** | SPDX 2.2 | JSON | `8087b5f878972e7574a57bf6db47ef2211a870a88ff4087ff006a392c9e764d9` |

## Directory Structure

The files are organized hierarchically by **Standard**, **Version**, and **Format**:

```text
test-vectors/
├── cyclonedx-1.7/      # CycloneDX v1.7 test cases
│   └── json/           # JSON-formatted SBOMs
│   └── xml/            # XML-formatted SBOMs
├── spdx-2.2/           # SPDX v2.2 test cases
│   └── json/           # JSON-formatted SBOMs
│   └── xml/            # XML-formatted SBOMs
└── spdx-3.0/           # SPDX v3.0 test cases
    └── json/           # JSON-formatted SBOMs
│   └── xml/            # XML-formatted SBOMs
├── manifest.json       # Master registry of files and expected hashes
```

For each use case, multiple variations of the same SBOM are provided:

- `*.[cdx|spdx].json`: The original source version of the SBOM.
- `*-min.[cdx|spdx].json`: The minified version of the SBOM.
- `*-pretty.[cdx|spdx].json`: The pretty-printed version of the SBOM.
- `*-canonical.[cdx|spdx].json`: The file state after JSF pruning and JCS
  normalization. This is provided to help developers debug their canonicalization
  logic before the final hashing step.

## Sources

The test vectors in this directory are sourced from the official example repositories
of the respective standards:

- [CycloneDX Use Cases](https://cyclonedx.org/use-cases/)
- [SPDX Specification Examples](https://github.com/spdx/spdx-spec)
