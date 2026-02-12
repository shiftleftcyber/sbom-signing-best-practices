use serde_json::Value;
use sha2::{Digest, Sha256};

pub fn compute_json_hash(input: &[u8]) -> Result<String, Box<dyn std::error::Error>> {
    // Parse JSON
    let mut data: Value = serde_json::from_slice(input)?;

    // Apply JSF Signature Exclusions
    if data["bomFormat"] == "CycloneDX" {
        let mut exclusions = Vec::new();
        if let Some(signature) = data.get_mut("signature")
            && let Some(signature_map) = signature.as_object_mut()
        {
            // Handle Dynamic Exclusions (from the 'excludes' property)
            if let Some(excludes) = signature_map.get("excludes").and_then(|e| e.as_array()) {
                for property in excludes {
                    if let Some(prop_str) = property.as_str() {
                        exclusions.push(prop_str.to_string());
                    }
                }
            }

            // JSF Core Requirement: Remove ONLY the "value" property from signature
            // This keeps 'algorithm', 'excludes', etc. in the hash.
            signature_map.remove("value");
        }

        // Remove the excluded properties
        if let Some(root) = data.as_object_mut() {
            for property in exclusions {
                root.remove(&property);
            }
        }
    }

    // Apply RFC 8785 (JCS)
    let canonical_bytes = serde_jcs::to_vec(&data)?;

    // Hash the canonical bytes
    let mut hasher = Sha256::new();
    hasher.update(canonical_bytes);
    let result = hasher.finalize();

    Ok(hex::encode(result))
}
