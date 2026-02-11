import { createHash } from 'node:crypto';
import canonicalize from 'canonicalize';

export function computeJsonHash(input) {
  // Parse JSON
  const data = JSON.parse(input);

  // Apply JSF Signature Exclusions
  if (data.bomFormat && data.bomFormat === 'CycloneDX') {
    if (data.signature) {
      
      // Handle Dynamic Exclusions (from the 'excludes' property)
      if (data.signature.excludes && Array.isArray(data.signature.excludes)) {
        for (const property of data.signature.excludes) {
          delete data[property];
        }
      }
      
      // Handle JSF Core Requirement: Delete ONLY the "value" property
      // This leaves 'algorithm', 'publicKey', etc., in the hash.
      delete data.signature.value;
    }
  }

  // Apply RFC 8785 (JCS)
  const canonicalJson = canonicalize(data);

  // Hash the canonical bytes
  return createHash('sha256').update(canonicalJson).digest('hex');
}
