package io.shiftleftcyber;

import static org.junit.jupiter.api.Assertions.assertEquals;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Objects;
import java.util.stream.Stream;
import java.util.stream.StreamSupport;
import org.junit.jupiter.api.DynamicTest;
import org.junit.jupiter.api.TestFactory;

class HashTest {
  private static final ObjectMapper MAPPER = new ObjectMapper();
  private static final Path TEST_VECTORS_DIR = Path.of("..", "test-vectors");
  private static final Path MANIFEST_PATH = TEST_VECTORS_DIR.resolve("manifest.json");

  @TestFactory
  Stream<DynamicTest> testHash() throws IOException {
    if (!Files.exists(MANIFEST_PATH)) {
      throw new IllegalStateException("Test manifest missing: " + MANIFEST_PATH.toAbsolutePath());
    }

    JsonNode manifest = MAPPER.readTree(Files.readAllBytes(MANIFEST_PATH));

    return StreamSupport.stream(manifest.spliterator(), false).flatMap(this::createTests);
  }

  private Stream<DynamicTest> createTests(JsonNode testCase) {
    String testName =
        Objects.requireNonNull(testCase.get("name"), "Manifest entry missing 'name'").asText();
    String expectedHash =
        Objects.requireNonNull(testCase.get("sha256"), "Manifest entry missing 'sha256'").asText();

    return Stream.of(
            new FileVariant("original", "file"),
            new FileVariant("min", "file-min"),
            new FileVariant("pretty", "file-pretty"),
            new FileVariant("canonical", "file-canonical"))
        .map(
            variant -> {
              String filePath = testCase.get(variant.path()).asText();
              String displayName = "%s [%s]".formatted(testName, variant.name());

              return DynamicTest.dynamicTest(
                  displayName,
                  () -> {
                    Path path = TEST_VECTORS_DIR.resolve(filePath);

                    if (!Files.exists(path)) {
                      throw new AssertionError(
                          "Test vector file not found: " + path.toAbsolutePath());
                    }

                    byte[] fileBytes = Files.readAllBytes(path);
                    String computedHash = Hash.computeJsonHash(fileBytes);

                    assertEquals(
                        expectedHash,
                        computedHash,
                        () ->
                            "Hash Mismatch!\nTest: %s\nVariant: %s\nPath: %s"
                                .formatted(testName, variant.name(), path.toAbsolutePath()));
                  });
            });
  }

  private record FileVariant(String name, String path) {}
}
