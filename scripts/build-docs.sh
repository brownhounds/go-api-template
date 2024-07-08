#!/bin/bash

. "$(dirname -- "$0")/core.sh"

cat schema.yml | yq -R > .swagger/schema.json
json_schema=$(cat .swagger/schema.json)

escaped_json_schema=$(printf '%s\n' "$json_schema" | sed -e 's/[\/&]/\\&/g')
cp .swagger/index.tpl .swagger/index.html

sed -i "s/{{JSON_SCHEMA}}/$escaped_json_schema/g" .swagger/index.html

log_message_success "✨ Swagger Docs Generated"
