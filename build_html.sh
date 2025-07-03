#!/bin/bash
set -e

SRC_DIR="pages"
OUT_DIR="public"

rm -rf ./public
mkdir ./public

generate_complete_page() {
    local mdfile="$1"
    local outfile="$2"
    local title="$3"
    
    # Convert markdown to HTML and save to temp file
    local temp_content=$(mktemp)
    pandoc "$mdfile" > "$temp_content"
    
    # Use a more robust approach with awk to handle multiline content
    awk -v title="$title" -v content_file="$temp_content" '
    {
        if ($0 ~ /{{\.Title}}/) {
            gsub(/{{\.Title}}/, title)
        }
        if ($0 ~ /{{printf "%s" \.Content}}/) {
            while ((getline line < content_file) > 0) {
                print line
            }
            close(content_file)
            next
        }
        print
    }' wrapper.html > "$OUT_DIR/$outfile"
    
    rm "$temp_content"
}

find "$SRC_DIR" -name '*.md' | while read -r mdfile; do
    relpath="${mdfile#$SRC_DIR/}"
    
    outfile="${relpath%.md}.html"

    mkdir -p "$(dirname "$OUT_DIR/$outfile")"

    filename=$(basename "$mdfile" .md)
    case "$filename" in
        "about") title="Conor Murphy" ;;
        "resume") title="Resume" ;;
        "contact") title="Conor Murphy" ;;
        "articles_list") title="Articles" ;;
        *) title="$filename" ;;
    esac

    generate_complete_page "$mdfile" "$outfile" "$title"
    
    echo "Generated: $OUT_DIR/$outfile"
done
