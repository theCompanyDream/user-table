#!/bin/bash
set -e

echo "Creating Users Table if doesn't Exist";

# Create the USERS table in the $POSTGRES_DB database
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE TABLE IF NOT EXISTS USERS (
        ID UUID NOT NULL,
        HASH varchar(64) NOT NULL,
        USER_NAME varchar(50) NOT NULL,
        FIRST_NAME varchar(255) NOT NULL,
        LAST_NAME varchar(255) NOT NULL,
        EMAIL varchar(255) NOT NULL,
        USER_STATUS varchar(1) NOT NULL,
        DEPARTMENT varchar(255),
        PRIMARY KEY(ID),
        UNIQUE(email),
        UNIQUE(USER_NAME)
    );
EOSQL
