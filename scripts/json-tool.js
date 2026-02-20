#!/usr/bin/env node

import canonicalize from 'canonicalize';
import fs from 'node:fs';
import path from 'node:path';

const [,, filePath, mode = 'pretty'] = process.argv;

if (!filePath) {
    console.error("Usage: node json-tool.js <file-path> [pretty|min|canon]");
    process.exit(1);
}

try {
    const absolutePath = path.resolve(filePath);
    const rawData = fs.readFileSync(absolutePath, 'utf8');
    const jsonObject = JSON.parse(rawData);
    
    let result;
    const lowerMode = mode.toLowerCase();

    if (lowerMode === 'min') {
        result = JSON.stringify(jsonObject);
        process.stdout.write(result); 
    } else if (lowerMode === 'canon') {
        result = canon(jsonObject); 
        process.stdout.write(result);
    } else {
        result = JSON.stringify(jsonObject, null, 4);
        console.log(result);
    }
} catch (err) {
    console.error(`Error processing ${filePath}: ${err.message}`);
    process.exit(1);
}

export function canon(data) {
  // Logic for CycloneDX / JSF Signature Exclusions
  if (data.bomFormat === 'CycloneDX' && data.signature) {
      if (Array.isArray(data.signature.excludes)) {
        for (const property of data.signature.excludes) {
          delete data[property];
        }
      }
      delete data.signature.value;
  }
  return canonicalize(data);
}
