package io.shiftleftcyber;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ArrayNode;
import com.fasterxml.jackson.databind.node.ObjectNode;
import java.security.MessageDigest;
import java.util.HexFormat;
import org.erdtman.jcs.JsonCanonicalizer;

/** Reference implementation for computing SBOM hashes. */
public class Hash {
  private static final ObjectMapper mapper = new ObjectMapper();

  /**
   * Computes the JSON hash of the provided SBOM input.
   *
   * @param input The SBOM as a byte array
   * @return The computed hash as a hexadecimal string
   * @throws Exception If an error occurs during processing
   */
  public static String computeJsonHash(byte[] input) throws Exception {
    // Parse JSON
    JsonNode root = mapper.readTree(input);
    if (!(root instanceof ObjectNode objectRoot)) {
      throw new IllegalArgumentException("Input is not a JSON object");
    }

    // Apply JSF Signature Exclusions
    JsonNode bomFormat = objectRoot.get("bomFormat");
    if (bomFormat != null && "CycloneDX".equals(bomFormat.asText())) {
      JsonNode signature = objectRoot.get("signature");
      if (signature != null && signature.isObject()) {
        ObjectNode signatureObject = (ObjectNode) signature;

        // Handle Dynamic Exclusions (from the 'excludes' property)
        JsonNode excludes = signature.get("excludes");
        if (excludes != null && excludes.isArray()) {
          ArrayNode excludesArray = (ArrayNode) excludes;
          // Remove each excluded property from the root object
          for (JsonNode entry : excludesArray) {
            String path = entry.asText();
            objectRoot.remove(path);
          }
        }

        // Handle JSF Core Requirement: Delete ONLY the "value" property
        // This leaves 'algorithm', 'publicKey', etc., in the hash.
        signatureObject.remove("value");
      }
    }

    // Apply RFC 8785 (JCS)
    String prunedJson = mapper.writeValueAsString(objectRoot);
    JsonCanonicalizer jcs = new JsonCanonicalizer(prunedJson.getBytes());
    byte[] canonicalBytes = jcs.getEncodedUTF8();

    // Hash the canonical bytes
    MessageDigest digest = MessageDigest.getInstance("SHA-256");
    byte[] hash = digest.digest(canonicalBytes);

    return HexFormat.of().formatHex(hash);
  }
}
