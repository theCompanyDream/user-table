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
        UNIQUE(EMAIL),
        UNIQUE(USER_NAME)
    );
    INSERT INTO USERS(ID, HASH, USER_NAME, FIRST_NAME, LAST_NAME, EMAIL, USER_STATUS, DEPARTMENT)
    VALUES ('558d700c-70e6-4dfe-a636-fa3dae53f8e8', '4656fbcc0cae25453c3ba728d5f1be3fbbffa9f7291de85f09070b4055bcac52', 'supermaueoeu', 'clark', 'kent', 'clark.keuoeuiuoet@gmail.com', 'I', 'oeuoeu'),
        ('df59addb-fd81-4a47-8207-2fb1854c6c0a', '7540856a1842e9fddf2f5d09b1e8386d6b100b97278bec316e5e22bc4541995e', 'uiaoei', 'clark', 'kent', 'clark.keuoeuiuouoet@gmail.comu', 'I', 'oeuoeu'),
        ('fb03747e-822b-4041-9269-002590c4274f', '1d74ba8bbcebf7b32d20c567f7c876c348f1f6572de24608ef79c489b892304e', 'uiaoeiuoeu', 'clark', 'kent', 'clark.keoeuuoet@gmail.comu ', 'I', 'oeuoeu'),
        ('6825f045-ad93-4244-997e-1be3199c3f12', '2e9a83447d9fe784f1ee3d59e7569e17bd569380e5dd2454a9f3173db4fc7829', 'uberman23423', 'clark', 'kent', 'clark.keoeuuoet@gmail.comu', 'I', 'oeuoeui')
EOSQL
