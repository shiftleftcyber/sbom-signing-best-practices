import { describe, test, expect } from 'vitest';
import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import { computeJsonHash } from './hash.js';

const manifestPath = join('..', 'test-vectors', 'manifest.json');
const manifest = JSON.parse(readFileSync(manifestPath, 'utf8'));

describe('SBOM Hashing Interoperability', () => {
  manifest.forEach((testCase) => {
    if (testCase.format !== 'json') {
      return;
    }

    describe(testCase.name, () => {
      // Map the manifest keys to a clean array for testing
      const variants = [
        { name: 'original', path: testCase.file },
        { name: 'min', path: testCase['file-min'] },
        { name: 'pretty', path: testCase['file-pretty'] },
        { name: 'canonical', path: testCase['file-canonical'] },
      ];

      test.each(variants)('Variant: $name', (variant) => {
        const fullPath = join('..', 'test-vectors', variant.path);
        const content = readFileSync(fullPath);

        const hash = computeJsonHash(content);

        expect(hash).toBe(testCase.sha256);
      });
    });
  });
});
