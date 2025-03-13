#!/bin/bash

YAML_FILE="config.yaml"

OUTPUT="bin/app"

INPUT="cmd/main.go"

if [[ ! -f "$YAML_FILE" ]]; then
  echo "Файл $YAML_FILE не найден."
  exit 1
fi

USE_UI=$(yq e '.use_ui' "$YAML_FILE")

if [[ "$USE_UI" == "true" ]]; then
    ./scripts/build_ui.sh
    go build -o $OUTPUT $INPUT
    $OUTPUT
    exit 0
else
    go build -o $OUTPUT $INPUT
    $OUTPUT
    exit 0
fi