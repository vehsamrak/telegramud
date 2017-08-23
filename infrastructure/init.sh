#!/bin/bash

PGPASSWORD=postgres

psql -h localhost -p 5433 -U postgres  -f init.sql
