#!/bin/bash

source local.env



if [ -z "$MIGRATION_DSN" ]; then
  echo "Переменная MIGRATION_DSN_L не определена. Проверьте файл local.env."
  exit 1
fi

GOOSE_DRIVER=postgres GOOSE_DBSTRING=${MIGRATION_DSN} ./bin/goose -dir migrations/ up -v
