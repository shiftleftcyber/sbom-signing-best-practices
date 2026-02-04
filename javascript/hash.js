import { createHash } from 'node:crypto';
import canonicalize from 'canonicalize';

export function computeJsonHash(input) {
  const data = JSON.parse(input);

  // Apply Exclusions (only for CycloneDX embedded signatures)
  if (data.bomFormat && data.bomFormat === 'CycloneDX') {
    // Check if there's a signature block
    if (data.signature) {
      // Extract any additional exclusions from the signature block
      if (data.signature.excludes && Array.isArray(data.signature.excludes)) {
        for (const property of data.signature.excludes) {
          delete data[property];
        }
      }
      // Always exclude the signature block itself
      delete data.signature;
    }
  }

  // RFC 8785 Canonicalization
  const canonicalJson = canonicalize(data);

  // SHA-256 Hashing
  return createHash('sha256').update(canonicalJson).digest('hex');
}
