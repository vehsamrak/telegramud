#!/bin/bash

PGPASSWORD=postgres

psql -h localhost -U postgres  -f init.sql
