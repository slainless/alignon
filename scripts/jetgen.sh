#!/usr/bin/env bash
bin/jet/jet -schema=public -dsn="${MY_ALIGNON_DATABASE_URL}" -path=./pkg/internal/artifact/database "$@"