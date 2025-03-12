#!/usr/bin/env python3
import os
import psycopg2
import uuid
import random
from faker import Faker

# Initialize Faker
fake = Faker()

# Example departments list
departments = [
    "Accounting", "Sales", "Engineering", "HR", "Marketing", "IT", "Operations"
]

def get_db_connection():
    # Get database connection parameters from environment variables.
    # Adjust defaults as needed.
    postgres_user = os.environ.get("POSTGRES_USER", "postgres")
    postgres_db = os.environ.get("POSTGRES_DB", "postgres")
    postgres_password = os.environ.get("POSTGRES_PASSWORD", "")
    postgres_host = os.environ.get("POSTGRES_HOST", "localhost")
    postgres_port = os.environ.get("POSTGRES_PORT", "5432")

    conn = psycopg2.connect(
		"postgresql://postgres.wlvdpewxbmhdumlumfka:6&Up6YyLFjYuoKt9@aws-0-us-east-2.pooler.supabase.com:6543/postgres"
    )
    return conn

def create_users_table(conn):
    with conn.cursor() as cur:
        cur.execute("""
            CREATE TABLE IF NOT EXISTS USERS (
                ID UUID NOT NULL,
                HASH VARCHAR(64) NOT NULL,
                USER_NAME VARCHAR(50) NOT NULL,
                FIRST_NAME VARCHAR(255) NOT NULL,
                LAST_NAME VARCHAR(255) NOT NULL,
                EMAIL VARCHAR(255) NOT NULL,
                USER_STATUS VARCHAR(1) NOT NULL,
                DEPARTMENT VARCHAR(255),
                PRIMARY KEY(ID),
                UNIQUE(EMAIL),
                UNIQUE(USER_NAME)
            );
        """)
        conn.commit()
        print("USERS table created (or already exists).")

def generate_fake_user():
    # Generate fake user data using Faker and Python's uuid
    user_id = str(uuid.uuid4())
    # Generate a random 64-character hex string (32 bytes * 2 hex digits each)
    hash_value = os.urandom(32).hex()
    user_name = fake.user_name()[:50]
    first_name = fake.first_name()[:255]
    last_name = fake.last_name()[:255]
    email = fake.email()[:255]
    # Example: random user status letter; adjust choices as needed.
    user_status = random.choice(['A', 'T', 'P'])
    department = random.choice(departments)
    return (user_id, hash_value, user_name, first_name, last_name, email, user_status, department)

def insert_fake_users(conn, num_records):
    with conn.cursor() as cur:
        for _ in range(num_records):
            user_data = generate_fake_user()
            cur.execute("""
                INSERT INTO USERS (ID, HASH, USER_NAME, FIRST_NAME, LAST_NAME, EMAIL, USER_STATUS, DEPARTMENT)
                VALUES (%s, %s, %s, %s, %s, %s, %s, %s);
            """, user_data)
        conn.commit()
        print(f"Inserted {num_records} fake user records.")

def main():
    num_records = 500  # Adjust the number of records you want to generate
    conn = get_db_connection()
    try:
        create_users_table(conn)
        insert_fake_users(conn, num_records)
    finally:
        conn.close()

if __name__ == '__main__':
    main()
