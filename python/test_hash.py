import json
import os
import pytest
from hash import compute_json_hash

MANIFEST_PATH = os.path.join("..", "test-vectors", "manifest.json")


def load_manifest():
    if not os.path.exists(MANIFEST_PATH):
        pytest.fail(f"Manifest not found at {MANIFEST_PATH}")

    with open(MANIFEST_PATH, "r") as f:
        return json.load(f)


def get_test_cases():
    manifest = load_manifest()
    cases = []

    for entry in manifest:
        variants = [
            ("original", entry["file"]),
            ("min", entry["file-min"]),
            ("pretty", entry["file-pretty"]),
            ("canonical", entry["file-canonical"]),
        ]

        for v_name, v_path in variants:
            unique_id = f"{entry['name']}_{v_name}"
            cases.append(pytest.param(v_path, entry["sha256"], id=unique_id))

    return cases


@pytest.mark.parametrize("file_path, expected_hash", get_test_cases())
def test_hash_vectors(file_path, expected_hash):
    full_path = os.path.join("..", "test-vectors", file_path)

    with open(full_path, "rb") as f:
        content = f.read()

    computed_hash = None
    try:
        computed_hash = compute_json_hash(content)
    except Exception as e:
        pytest.fail(f"Hash error for {file_path}: {e}")

    assert computed_hash == expected_hash, (
        f"Hash Mismatch!\n"
        f"File: {file_path}\n"
        f"Expected: {expected_hash}\n"
        f"Got:      {computed_hash}"
    )
