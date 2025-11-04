KYVERNO_CRDS_VERSION="1.15.0"

## Created by Cursor, beware of AI sloppiness

# Template CRDs
helm template kyverno-crds \
    https://giantswarm.github.io/giantswarm-catalog/kyverno-crds-${KYVERNO_CRDS_VERSION}.tgz \
    --dry-run=client > crds/crds.yaml

# Split CRDs into individual files using '.metadata.name' as filename
# Convert camelCase to lowercase with underscores (e.g., AdmissionPolicy -> admission_policy.yaml)

# Function to convert camelCase to snake_case
camel_to_snake() {
    echo "$1" | sed 's/\([A-Z]\)/_\1/g' | sed 's/^_//' | tr '[:upper:]' '[:lower:]'
}

# Split the YAML file into individual documents
temp_dir=$(mktemp -d)

# Use awk to split documents (handles first document without leading ---)
awk -v temp_dir="$temp_dir" '
    BEGIN { doc_num=0 }
    /^---$/ {
        if (doc) {
            filename = temp_dir "/doc_" sprintf("%02d", ++doc_num) ".yaml"
            printf "%s", doc > filename
            close(filename)
            doc=""
        }
        next
    }
    {
        doc=doc $0 "\n"
    }
    END {
        if (doc) {
            filename = temp_dir "/doc_" sprintf("%02d", ++doc_num) ".yaml"
            printf "%s", doc > filename
            close(filename)
        }
    }
' crds/crds.yaml

# Process each split document
for doc_file in "$temp_dir"/doc_*.yaml; do
    # Skip if file doesn't exist or is empty
    [ ! -f "$doc_file" ] && continue
    [ ! -s "$doc_file" ] && continue

    # Extract the kind name from this document
    kind=$(yq eval '.spec.names.kind' "$doc_file" 2>/dev/null | grep -v "^---$" | grep -v "^$" | head -1)

    if [ -n "$kind" ] && [ "$kind" != "null" ] && [ "$kind" != "" ]; then
        # Convert camelCase to snake_case
        filename=$(camel_to_snake "$kind")

        # Copy the document to the final location
        cp "$doc_file" "crds/${filename}.yaml"
    fi
done

# Clean up temporary directory
rm -rf "$temp_dir"
rm -f crds/crds.yaml