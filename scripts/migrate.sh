#!/usr/bin/env bash
bin/migrate/migrate -database "${MY_ALIGNON_DATABASE_URL}" -path db/migrations "$@"