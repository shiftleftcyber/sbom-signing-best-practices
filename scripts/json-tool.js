#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

// Grab arguments: [0] is node, [1] is script path, [2] is file, [3] is mode
const [,, filePath, mode = 'pretty'] = process.argv;

if (!filePath) {
    console.error("Usage: node json-tool.js <file-path> [pretty|min]");
    process.exit(1);
}

try {
    // Resolve and read the file
    const absolutePath = path.resolve(filePath);
    const rawData = fs.readFileSync(absolutePath, 'utf8');
    
    // Parse the JSON file
    const jsonObject = JSON.parse(rawData);
    
    // Process based on mode
    let result;
    if (mode.toLowerCase() === 'min') {
        result = JSON.stringify(jsonObject);
        process.stdout.write(result);
    } else {
        // Pretty print using 4-space indentation
        result = JSON.stringify(jsonObject, null, 4);
        console.log(result);
    }
} catch (err) {
    console.error(`Error: ${err.message}`);
    process.exit(1);
}
