#!/bin/bash
export SERVER_ADDRESS=localhost:8080
export BASE_URL=http://localhost:8080
export FILE_STORAGE_PATH=urls.gob
export DATABASE_DSN=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
export PG_MIGRATIONS_PATH=file://../../repository/pg/migrations