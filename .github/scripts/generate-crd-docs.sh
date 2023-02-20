#!/bin/bash

API_DOMAIN="keptn.sh"
API_ROOT='operator/apis/'
TEMPLATE_DIR='./template'
RENDERER='markdown'
RENDERER_CONFIG_FILE='.github/crd-docs-generator-config.yaml'

echo "Running CRD docs auto-generator..."

for api_group in "$API_ROOT"*; do
  for api_version in "$api_group"/*; do
    sanitized_api_group="${api_group#$API_ROOT}"
    sanitized_api_version="${api_version#$API_ROOT$sanitized_api_group/}"
    OUTPUT_PATH="./docs/content/en/docs/crd-ref/$sanitized_api_group/$sanitized_api_version"

    echo "Arguments:"
    echo "TEMPLATE_DIR: $TEMPLATE_DIR"
    echo "API_ROOT: $API_ROOT"
    echo "API_GROUP: $sanitized_api_group"
    echo "API_VERSION: $sanitized_api_version"
    echo "RENDERER: $RENDERER"
    echo "RENDERER_CONFIG_FILE: $RENDERER_CONFIG_FILE"
    echo "OUTPUT_PATH: $OUTPUT_PATH/_index.md"

    echo "Creating docs folder $OUTPUT_PATH..."
    mkdir -p "$OUTPUT_PATH"

    echo "Generating CRD docs for $sanitized_api_group.$API_DOMAIN/$sanitized_api_version..."
    crd-ref-docs \
      --templates-dir "$TEMPLATE_DIR" \
      --source-path="./$api_version" \
      --renderer="$RENDERER" \
      --config "$RENDERER_CONFIG_FILE" \
      --output-path "$OUTPUT_PATH/_index.md"
    echo "---------------------"
  done
done
