use sbom_hash_rust::compute_json_hash;
use serde::Deserialize;
use std::fs;
use std::path::PathBuf;

#[derive(Deserialize)]
struct TestCase {
    name: String,
    file: String,
    #[serde(rename = "file-min")]
    file_min: Option<String>,
    #[serde(rename = "file-pretty")]
    file_pretty: Option<String>,
    #[serde(rename = "file-canonical")]
    file_canonical: Option<String>,
    format: String,
    sha256: String,
}

#[test]
fn test_manifest_vectors() {
    let manifest_path = PathBuf::from("../test-vectors/manifest.json");
    let content = fs::read_to_string(manifest_path).expect("Failed to read manifest");
    let manifest: Vec<TestCase> = serde_json::from_str(&content).expect("Failed to parse manifest");

    for test_case in manifest {
        if test_case.format != "json" {
            continue;
        }

        let variants = vec![
            ("original", Some(test_case.file)),
            ("min", test_case.file_min),
            ("pretty", test_case.file_pretty),
            ("canonical", test_case.file_canonical),
        ];

        for (v_name, v_path) in variants {
            if let Some(path_str) = v_path {
                let full_path = PathBuf::from("../test-vectors").join(path_str);
                let bytes = fs::read(full_path).expect("Failed to read SBOM file");

                let computed = compute_json_hash(&bytes)
                    .expect(&format!("Hash failed for {} [{}]", test_case.name, v_name));

                assert_eq!(
                    computed, test_case.sha256,
                    "Mismatch in {} [{}]",
                    test_case.name, v_name
                );
            }
        }
    }
}
