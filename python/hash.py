import hashlib
import json
import jcs


def compute_json_hash(input_bytes: bytes) -> str:
    # Parse JSON
    data = json.loads(input_bytes)

    # Apply JSF Signature Exclusions
    if data.get("bomFormat") == "CycloneDX":
        signature = data.get("signature")
        if isinstance(signature, dict):
            # Handle Dynamic Exclusions (from the 'excludes' property)
            exclusions = signature.get("excludes", [])

            # Apply all exclusions
            for property_name in exclusions:
                data.pop(property_name, None)

            # Handle JSF Core Requirement: Delete ONLY the "value" property
            # This leaves 'algorithm', 'publicKey', etc., in the hash.
            signature.pop("value", None)

    # Apply RFC 8785 (JCS)
    canonical_bytes = jcs.canonicalize(data)

    # Hash the canonical bytes
    return hashlib.sha256(canonical_bytes).hexdigest()
