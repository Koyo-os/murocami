#!/bin/bash

OUTPUT="bin/app"

INPUT="cmd/server/main.go"

YAML_FILE="config.yaml"

# Проверка, существует ли файл
if [[ ! -f "$YAML_FILE" ]]; then
  echo "Файл $YAML_FILE не найден."
  exit 1
fi

# Используем yq для парсинга YAML
USE_UI="false"

# Проверка значения use_ui
if [[ "$USE_UI" == "true" ]]; then
    ./scripts/build_ui.sh
    go build -o $OUTPUT ./$INPUT
    $OUTPUT
    exit 0
else
    go build -o $OUTPUT ./$INPUT
    $OUTPUT
    exit 0
fi