import hashlib
import json
import jcs


def compute_json_hash(input_bytes: bytes) -> str:
    # 1. Parse JSON
    data = json.loads(input_bytes)

    # 2. Apply Exclusions (CycloneDX signature pruning)
    if data.get("bomFormat") == "CycloneDX":
        signature = data.get("signature")
        if isinstance(signature, dict):
            # Extract exclusions from the signature block
            exclusions = signature.get("excludes", [])
            # Always remove the signature itself
            if "signature" not in exclusions:
                exclusions.append("signature")

            # Apply all exclusions
            for property_name in exclusions:
                data.pop(property_name, None)

    # 3. Apply RFC 8785 (JCS).
    canonical_bytes = jcs.canonicalize(data)

    # 4. Hash the canonical bytes
    return hashlib.sha256(canonical_bytes).hexdigest()
