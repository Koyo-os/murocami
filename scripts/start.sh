#!/bin/bash

YAML_FILE="config.yaml"

if [[ ! -f "$YAML_FILE" ]]; then
  echo "Файл $YAML_FILE не найден."
  exit 1
fi

USE_UI=$(yq e '.use_ui' "$YAML_FILE")

if [[ "$USE_UI" == "true" ]]; then
    ./scripts/build_ui.sh
    make build
    ./bin/app
    exit 0
else
  make build
  ./bin/app 
  exit 0
fi