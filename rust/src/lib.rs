use serde_json::Value;
use sha2::{Digest, Sha256};

pub fn compute_json_hash(input: &[u8]) -> Result<String, Box<dyn std::error::Error>> {
    // Parse JSON
    let mut data: Value = serde_json::from_slice(input)?;

    // Apply Exclusions (only for CycloneDX embedded signatures)
    if data["bomFormat"] == "CycloneDX" {
        let mut exclusions_to_remove = Vec::new();
        if let Some(signature) = data.get("signature") {
            // Check for "excludes" array in the signature block
            if let Some(excludes) = signature.get("excludes").and_then(|e| e.as_array()) {
                for property in excludes {
                    if let Some(prop_str) = property.as_str() {
                        exclusions_to_remove.push(prop_str.to_string());
                    }
                }
            }
        }

        // Remove the excluded properties
        if let Some(obj) = data.as_object_mut() {
            for property in exclusions_to_remove {
                obj.remove(&property);
            }
            // Always exclude the signature block itself
            obj.remove("signature");
        }
    }

    // RFC 8785 Canonicalization
    let canonical_bytes = serde_jcs::to_vec(&data)?;

    // SHA-256 Hashing
    let mut hasher = Sha256::new();
    hasher.update(canonical_bytes);
    let result = hasher.finalize();

    Ok(hex::encode(result))
}
