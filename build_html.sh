#!/bin/bash
set -e

SRC_DIR="pages"
OUT_DIR="dist"

# Find all .md files in pages/, recursively
find "$SRC_DIR" -name '*.md' | while read -r mdfile; do
    # Strip the source directory prefix
    relpath="${mdfile#$SRC_DIR/}"
    
    # Change extension to .html
    outfile="${relpath%.md}.html"

    # Ensure destination directory exists
    mkdir -p "$(dirname "$OUT_DIR/$outfile")"

    # Convert Markdown to HTML
    pandoc "$mdfile" -o "$OUT_DIR/$outfile"
    
    echo "Generated: $OUT_DIR/$outfile"
done
