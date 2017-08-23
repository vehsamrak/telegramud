#!/bin/bash

PGPASSWORD=postgres
PG_HOST=localhost
PG_USER=postgres
PG_DATABASE=postgres
TIMEOUT=20

echo "Importing data to database..."

until psql -h ${PG_HOST} -p 5433 -U ${PG_USER} -c "SELECT 1" > /dev/null 2>&1 || [ ${TIMEOUT} -eq 0 ]; do
  echo "Waiting for postgres server, $((TIMEOUT--)) remaining attempts..."
  sleep 1
done

psql -h ${PG_HOST} -p 5433 -U ${PG_USER} -f init.sql
